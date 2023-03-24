// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

// This file defines the 'references' query based on a serializable
// index constructed during type checking, thus avoiding the need to
// type-check packages at search time.
//
// See the ./xrefs/ subpackage for the index construction and lookup.
//
// This implementation does not intermingle objects from distinct
// calls to TypeCheck.

import (
	"context"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sort"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/tools/go/types/objectpath"
	"golang.org/x/tools/gopls/internal/lsp/protocol"
	"golang.org/x/tools/gopls/internal/lsp/safetoken"
	"golang.org/x/tools/gopls/internal/lsp/source/methodsets"
	"golang.org/x/tools/gopls/internal/span"
	"golang.org/x/tools/internal/bug"
	"golang.org/x/tools/internal/event"
)

// References returns a list of all references (sorted with
// definitions before uses) to the object denoted by the identifier at
// the given file/position, searching the entire workspace.
func References(ctx context.Context, snapshot Snapshot, fh FileHandle, pp protocol.Position, includeDeclaration bool) ([]protocol.Location, error) {
	references, err := references(ctx, snapshot, fh, pp, includeDeclaration)
	if err != nil {
		return nil, err
	}
	locations := make([]protocol.Location, len(references))
	for i, ref := range references {
		locations[i] = ref.location
	}
	return locations, nil
}

// A reference describes an identifier that refers to the same
// object as the subject of a References query.
type reference struct {
	isDeclaration bool
	location      protocol.Location
	pkgPath       PackagePath // of declaring package (same for all elements of the slice)
}

// references returns a list of all references (sorted with
// definitions before uses) to the object denoted by the identifier at
// the given file/position, searching the entire workspace.
func references(ctx context.Context, snapshot Snapshot, f FileHandle, pp protocol.Position, includeDeclaration bool) ([]reference, error) {
	ctx, done := event.Start(ctx, "source.references")
	defer done()

	// Is the cursor within the package name declaration?
	_, inPackageName, err := parsePackageNameDecl(ctx, snapshot, f, pp)
	if err != nil {
		return nil, err
	}

	var refs []reference
	if inPackageName {
		refs, err = packageReferences(ctx, snapshot, f.URI())
	} else {
		refs, err = ordinaryReferences(ctx, snapshot, f.URI(), pp)
	}
	if err != nil {
		return nil, err
	}

	sort.Slice(refs, func(i, j int) bool {
		x, y := refs[i], refs[j]
		if x.isDeclaration != y.isDeclaration {
			return x.isDeclaration // decls < refs
		}
		return protocol.CompareLocation(x.location, y.location) < 0
	})

	// De-duplicate by location, and optionally remove declarations.
	out := refs[:0]
	for _, ref := range refs {
		if !includeDeclaration && ref.isDeclaration {
			continue
		}
		if len(out) == 0 || out[len(out)-1].location != ref.location {
			out = append(out, ref)
		}
	}
	refs = out

	return refs, nil
}

// packageReferences returns a list of references to the package
// declaration of the specified name and uri by searching among the
// import declarations of all packages that directly import the target
// package.
func packageReferences(ctx context.Context, snapshot Snapshot, uri span.URI) ([]reference, error) {
	metas, err := snapshot.MetadataForFile(ctx, uri)
	if err != nil {
		return nil, err
	}
	if len(metas) == 0 {
		return nil, fmt.Errorf("found no package containing %s", uri)
	}

	var refs []reference

	// Find external references to the package declaration
	// from each direct import of the package.
	//
	// The narrowest package is the most broadly imported,
	// so we choose it for the external references.
	//
	// But if the file ends with _test.go then we need to
	// find the package it is testing; there's no direct way
	// to do that, so pick a file from the same package that
	// doesn't end in _test.go and start over.
	narrowest := metas[0]
	if narrowest.ForTest != "" && strings.HasSuffix(string(uri), "_test.go") {
		for _, f := range narrowest.CompiledGoFiles {
			if !strings.HasSuffix(string(f), "_test.go") {
				return packageReferences(ctx, snapshot, f)
			}
		}
		// This package has no non-test files.
		// Skip the search for external references.
		// (Conceivably one could blank-import an empty package, but why?)
	} else {
		rdeps, err := snapshot.ReverseDependencies(ctx, narrowest.ID, false) // direct
		if err != nil {
			return nil, err
		}
		for _, rdep := range rdeps {
			for _, uri := range rdep.CompiledGoFiles {
				fh, err := snapshot.ReadFile(ctx, uri)
				if err != nil {
					return nil, err
				}
				f, err := snapshot.ParseGo(ctx, fh, ParseHeader)
				if err != nil {
					return nil, err
				}
				for _, imp := range f.File.Imports {
					if rdep.DepsByImpPath[UnquoteImportPath(imp)] == narrowest.ID {
						refs = append(refs, reference{
							isDeclaration: false,
							location:      mustLocation(f, imp),
							pkgPath:       narrowest.PkgPath,
						})
					}
				}
			}
		}
	}

	// Find internal "references" to the package from
	// of each package declaration in the target package itself.
	//
	// The widest package (possibly a test variant) has the
	// greatest number of files and thus we choose it for the
	// "internal" references.
	widest := metas[len(metas)-1]
	for _, uri := range widest.CompiledGoFiles {
		fh, err := snapshot.ReadFile(ctx, uri)
		if err != nil {
			return nil, err
		}
		f, err := snapshot.ParseGo(ctx, fh, ParseHeader)
		if err != nil {
			return nil, err
		}
		refs = append(refs, reference{
			isDeclaration: true, // (one of many)
			location:      mustLocation(f, f.File.Name),
			pkgPath:       widest.PkgPath,
		})
	}

	return refs, nil
}

// ordinaryReferences computes references for all ordinary objects (not package declarations).
func ordinaryReferences(ctx context.Context, snapshot Snapshot, uri span.URI, pp protocol.Position) ([]reference, error) {
	// Strategy: use the reference information computed by the
	// type checker to find the declaration. First type-check this
	// package to find the declaration, then type check the
	// declaring package (which may be different), plus variants,
	// to find local (in-package) references.
	// Global references are satisfied by the index.

	// Strictly speaking, a wider package could provide a different
	// declaration (e.g. because the _test.go files can change the
	// meaning of a field or method selection), but the narrower
	// package reports the more broadly referenced object.
	pkg, pgf, err := PackageForFile(ctx, snapshot, uri, NarrowestPackage)
	if err != nil {
		return nil, err
	}

	// Find the selected object (declaration or reference).
	pos, err := pgf.PositionPos(pp)
	if err != nil {
		return nil, err
	}
	candidates, _, err := objectsAt(pkg.GetTypesInfo(), pgf.File, pos)
	if err != nil {
		return nil, err
	}

	// Pick first object arbitrarily.
	// The case variables of a type switch have different
	// types but that difference is immaterial here.
	var obj types.Object
	for obj = range candidates {
		break
	}
	if obj == nil {
		return nil, ErrNoIdentFound // can't happen
	}

	// nil, error, error.Error, iota, or other built-in?
	if obj.Pkg() == nil {
		// For some reason, existing tests require that iota has no references,
		// nor an error. TODO(adonovan): do something more principled.
		if obj.Name() == "iota" {
			return nil, nil
		}

		return nil, fmt.Errorf("references to builtin %q are not supported", obj.Name())
	}

	// Find metadata of all packages containing the object's defining file.
	// This may include the query pkg, and possibly other variants.
	declPosn := safetoken.StartPosition(pkg.FileSet(), obj.Pos())
	declURI := span.URIFromPath(declPosn.Filename)
	variants, err := snapshot.MetadataForFile(ctx, declURI)
	if err != nil {
		return nil, err
	}
	if len(variants) == 0 {
		return nil, fmt.Errorf("no packages for file %q", declURI) // can't happen
	}

	// Is object exported?
	// If so, compute scope and targets of the global search.
	var (
		globalScope   = make(map[PackageID]*Metadata)
		globalTargets map[PackagePath]map[objectpath.Path]unit
		expansions    = make(map[*Metadata]unit) // packages that caused search expansion
	)
	// TODO(adonovan): what about generic functions? Need to consider both
	// uninstantiated and instantiated. The latter have no objectpath. Use Origin?
	if path, err := objectpath.For(obj); err == nil && obj.Exported() {
		pkgPath := variants[0].PkgPath // (all variants have same package path)
		globalTargets = map[PackagePath]map[objectpath.Path]unit{
			pkgPath: {path: {}}, // primary target
		}

		// How far need we search?
		// For package-level objects, we need only search the direct importers.
		// For fields and methods, we must search transitively.
		transitive := obj.Pkg().Scope().Lookup(obj.Name()) != obj

		// The scope is the union of rdeps of each variant.
		// (Each set is disjoint so there's no benefit to
		// to combining the metadata graph traversals.)
		for _, m := range variants {
			rdeps, err := snapshot.ReverseDependencies(ctx, m.ID, transitive)
			if err != nil {
				return nil, err
			}
			for id, rdep := range rdeps {
				globalScope[id] = rdep
			}
		}

		// Is object a method?
		//
		// If so, expand the search so that the targets include
		// all methods that correspond to it through interface
		// satisfaction, and the scope includes the rdeps of
		// the package that declares each corresponding type.
		//
		// 'expansions' records the packages that declared
		// such types.
		if recv := effectiveReceiver(obj); recv != nil {
			if err := expandMethodSearch(ctx, snapshot, obj.(*types.Func), recv, globalScope, globalTargets, expansions); err != nil {
				return nil, err
			}
		}
	}

	// The search functions will call report(loc) for each hit.
	var (
		refsMu sync.Mutex
		refs   []reference
	)
	report := func(loc protocol.Location, isDecl bool) {
		ref := reference{
			isDeclaration: isDecl,
			location:      loc,
			pkgPath:       pkg.Metadata().PkgPath,
		}
		refsMu.Lock()
		refs = append(refs, ref)
		refsMu.Unlock()
	}

	// Loop over the variants of the declaring package,
	// and perform both the local (in-package) and global
	// (cross-package) searches, in parallel.
	//
	// TODO(adonovan): opt: support LSP reference streaming. See:
	// - https://github.com/microsoft/vscode-languageserver-node/pull/164
	// - https://github.com/microsoft/language-server-protocol/pull/182
	//
	// Careful: this goroutine must not return before group.Wait.
	var group errgroup.Group

	// Compute local references for each variant.
	// The target objects are identified by (URI, offset).
	for _, m := range variants {
		// We want the ordinary importable package,
		// plus any test-augmented variants, since
		// declarations in _test.go files may change
		// the reference of a selection, or even a
		// field into a method or vice versa.
		//
		// But we don't need intermediate test variants,
		// as their local references will be covered
		// already by other variants.
		if m.IsIntermediateTestVariant() {
			continue
		}
		m := m
		group.Go(func() error {
			// TODO(adonovan): opt: batch these TypeChecks.
			pkgs, err := snapshot.TypeCheck(ctx, m.ID)
			if err != nil {
				return err
			}
			pkg := pkgs[0]

			// Find the declaration of the corresponding
			// object in this package based on (URI, offset).
			pgf, err := pkg.File(declURI)
			if err != nil {
				return err
			}
			pos, err := safetoken.Pos(pgf.Tok, declPosn.Offset)
			if err != nil {
				return err
			}
			objects, _, err := objectsAt(pkg.GetTypesInfo(), pgf.File, pos)
			if err != nil {
				return err // unreachable? (probably caught earlier)
			}

			// Report the locations of the declaration(s).
			// TODO(adonovan): what about for corresponding methods? Add tests.
			for _, node := range objects {
				report(mustLocation(pgf, node), true)
			}

			// Convert targets map to set.
			targets := make(map[types.Object]bool)
			for obj := range objects {
				targets[obj] = true
			}

			return localReferences(pkg, targets, report)
		})
	}

	// Also compute local references within packages that declare
	// corresponding methods (see above), which expand the global search.
	// The target objects are identified by (PkgPath, objectpath).
	for m := range expansions {
		m := m
		group.Go(func() error {
			// TODO(adonovan): opt: batch these TypeChecks.
			pkgs, err := snapshot.TypeCheck(ctx, m.ID)
			if err != nil {
				return err
			}
			pkg := pkgs[0]

			targets := make(map[types.Object]bool)
			for objpath := range globalTargets[m.PkgPath] {
				obj, err := objectpath.Object(pkg.GetTypes(), objpath)
				if err != nil {
					return err // can't happen?
				}
				targets[obj] = true
			}
			return localReferences(pkg, targets, report)
		})
	}

	// Compute global references for selected reverse dependencies.
	group.Go(func() error {
		var globalIDs []PackageID
		for id := range globalScope {
			globalIDs = append(globalIDs, id)
		}
		indexes, err := snapshot.References(ctx, globalIDs...)
		if err != nil {
			return err
		}
		for _, index := range indexes {
			for _, loc := range index.Lookup(globalTargets) {
				report(loc, false)
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		return nil, err
	}
	return refs, nil
}

// expandMethodSearch expands the scope and targets of a global search
// for an exported method to include all methods that correspond to
// it through interface satisfaction.
//
// Each package that declares a corresponding type is added to
// expansions so that we can also find local references to the type
// within the package, which of course requires type checking.
//
// recv is the method's effective receiver type, for method-set computations.
func expandMethodSearch(ctx context.Context, snapshot Snapshot, method *types.Func, recv types.Type, scope map[PackageID]*Metadata, targets map[PackagePath]map[objectpath.Path]unit, expansions map[*Metadata]unit) error {
	// Compute the method-set fingerprint used as a key to the global search.
	key, hasMethods := methodsets.KeyOf(recv)
	if !hasMethods {
		return bug.Errorf("KeyOf(%s)={} yet %s is a method", recv, method)
	}
	metas, err := snapshot.AllMetadata(ctx)
	if err != nil {
		return err
	}
	allIDs := make([]PackageID, 0, len(metas))
	for _, m := range metas {
		allIDs = append(allIDs, m.ID)
	}
	// Search the methodset index of each package in the workspace.
	indexes, err := snapshot.MethodSets(ctx, allIDs...)
	if err != nil {
		return err
	}
	var mu sync.Mutex // guards scope, targets, expansions
	var group errgroup.Group
	for i, index := range indexes {
		i := i
		index := index
		group.Go(func() error {
			// Consult index for matching methods.
			results := index.Search(key, method.Name())
			if len(results) == 0 {
				return nil
			}

			// Expand global search scope to include rdeps of this pkg.
			rdeps, err := snapshot.ReverseDependencies(ctx, allIDs[i], true)
			if err != nil {
				return err
			}

			mu.Lock()
			defer mu.Unlock()

			expansions[metas[i]] = unit{}

			for _, rdep := range rdeps {
				scope[rdep.ID] = rdep
			}

			// Add each corresponding method the to set of global search targets.
			for _, res := range results {
				methodPkg := PackagePath(res.PkgPath)
				opaths, ok := targets[methodPkg]
				if !ok {
					opaths = make(map[objectpath.Path]unit)
					targets[methodPkg] = opaths
				}
				opaths[res.ObjectPath] = unit{}
			}
			return nil
		})
	}
	return group.Wait()
}

// localReferences traverses syntax and reports each reference to one of the target objects.
func localReferences(pkg Package, targets map[types.Object]bool, report func(loc protocol.Location, isDecl bool)) error {
	// If we're searching for references to a method, broaden the
	// search to include references to corresponding methods of
	// mutually assignable receiver types.
	// (We use a slice, but objectsAt never returns >1 methods.)
	var methodRecvs []types.Type
	var methodName string // name of an arbitrary target, iff a method
	for obj := range targets {
		if t := effectiveReceiver(obj); t != nil {
			methodRecvs = append(methodRecvs, t)
			methodName = obj.Name()
		}
	}

	// matches reports whether obj either is or corresponds to a target.
	// (Correspondence is defined as usual for interface methods.)
	matches := func(obj types.Object) bool {
		if targets[obj] {
			return true
		} else if methodRecvs != nil && obj.Name() == methodName {
			if orecv := effectiveReceiver(obj); orecv != nil {
				for _, mrecv := range methodRecvs {
					if concreteImplementsIntf(orecv, mrecv) {
						return true
					}
				}
			}
		}
		return false
	}

	// Scan through syntax looking for uses of one of the target objects.
	for _, pgf := range pkg.CompiledGoFiles() {
		ast.Inspect(pgf.File, func(n ast.Node) bool {
			if id, ok := n.(*ast.Ident); ok {
				if obj, ok := pkg.GetTypesInfo().Uses[id]; ok && matches(obj) {
					report(mustLocation(pgf, id), false)
				}
			}
			return true
		})
	}
	return nil
}

// effectiveReceiver returns the effective receiver type for method-set
// comparisons for obj, if it is a method, or nil otherwise.
func effectiveReceiver(obj types.Object) types.Type {
	if fn, ok := obj.(*types.Func); ok {
		if recv := fn.Type().(*types.Signature).Recv(); recv != nil {
			return methodsets.EnsurePointer(recv.Type())
		}
	}
	return nil
}

// objectsAt returns the non-empty set of objects denoted (def or use)
// by the specified position within a file syntax tree, or an error if
// none were found.
//
// The result may contain more than one element because all case
// variables of a type switch appear to be declared at the same
// position.
//
// Each object is mapped to the syntax node that was treated as an
// identifier, which is not always an ast.Ident. The second component
// of the result is the innermost node enclosing pos.
//
// TODO(adonovan): factor in common with referencedObject.
func objectsAt(info *types.Info, file *ast.File, pos token.Pos) (map[types.Object]ast.Node, ast.Node, error) {
	path := pathEnclosingObjNode(file, pos)
	if path == nil {
		return nil, nil, ErrNoIdentFound
	}

	targets := make(map[types.Object]ast.Node)

	switch leaf := path[0].(type) {
	case *ast.Ident:
		// If leaf represents an implicit type switch object or the type
		// switch "assign" variable, expand to all of the type switch's
		// implicit objects.
		if implicits, _ := typeSwitchImplicits(info, path); len(implicits) > 0 {
			for _, obj := range implicits {
				targets[obj] = leaf
			}
		} else {
			obj := info.ObjectOf(leaf)
			if obj == nil {
				return nil, nil, fmt.Errorf("%w for %q", errNoObjectFound, leaf.Name)
			}
			targets[obj] = leaf
		}
	case *ast.ImportSpec:
		// Look up the implicit *types.PkgName.
		obj := info.Implicits[leaf]
		if obj == nil {
			return nil, nil, fmt.Errorf("%w for import %s", errNoObjectFound, UnquoteImportPath(leaf))
		}
		targets[obj] = leaf
	}

	if len(targets) == 0 {
		return nil, nil, fmt.Errorf("objectAt: internal error: no targets") // can't happen
	}
	return targets, path[0], nil
}

// mustLocation reports the location interval a syntax node,
// which must belong to m.File.
//
// Safe for use only by references and implementations.
func mustLocation(pgf *ParsedGoFile, n ast.Node) protocol.Location {
	loc, err := pgf.NodeLocation(n)
	if err != nil {
		panic(err) // can't happen in references or implementations
	}
	return loc
}