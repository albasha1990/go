// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"cmd/go/internal/base"
	"cmd/go/internal/cache"
	"cmd/go/internal/cfg"
	"cmd/go/internal/load"
	"cmd/go/internal/str"
	"cmd/internal/buildid"
)

// Build IDs
//
// Go packages and binaries are stamped with build IDs that record both
// the action ID, which is a hash of the inputs to the action that produced
// the packages or binary, and the content ID, which is a hash of the action
// output, namely the archive or binary itself. The hash is the same one
// used by the build artifact cache (see cmd/go/internal/cache), but
// truncated when stored in packages and binaries, as the full length is not
// needed and is a bit unwieldy. The precise form is
//
//	actionID/[.../]contentID
//
// where the actionID and contentID are prepared by hashToString below.
// and are found by looking for the first or last slash.
// Usually the buildID is simply actionID/contentID, but see below for an
// exception.
//
// The build ID serves two primary purposes.
//
// 1. The action ID half allows installed packages and binaries to serve as
// one-element cache entries. If we intend to build math.a with a given
// set of inputs summarized in the action ID, and the installed math.a already
// has that action ID, we can reuse the installed math.a instead of rebuilding it.
//
// 2. The content ID half allows the easy preparation of action IDs for steps
// that consume a particular package or binary. The content hash of every
// input file for a given action must be included in the action ID hash.
// Storing the content ID in the build ID lets us read it from the file with
// minimal I/O, instead of reading and hashing the entire file.
// This is especially effective since packages and binaries are typically
// the largest inputs to an action.
//
// Separating action ID from content ID is important for reproducible builds.
// The compiler is compiled with itself. If an output were represented by its
// own action ID (instead of content ID) when computing the action ID of
// the next step in the build process, then the compiler could never have its
// own input action ID as its output action ID (short of a miraculous hash collision).
// Instead we use the content IDs to compute the next action ID, and because
// the content IDs converge, so too do the action IDs and therefore the
// build IDs and the overall compiler binary. See cmd/dist's cmdbootstrap
// for the actual convergence sequence.
//
// The “one-element cache” purpose is a bit more complex for installed
// binaries. For a binary, like cmd/gofmt, there are two steps: compile
// cmd/gofmt/*.go into main.a, and then link main.a into the gofmt binary.
// We do not install gofmt's main.a, only the gofmt binary. Being able to
// decide that the gofmt binary is up-to-date means computing the action ID
// for the final link of the gofmt binary and comparing it against the
// already-installed gofmt binary. But computing the action ID for the link
// means knowing the content ID of main.a, which we did not keep.
// To sidestep this problem, each binary actually stores an expanded build ID:
//
//	actionID(binary)/actionID(main.a)/contentID(main.a)/contentID(binary)
//
// (Note that this can be viewed equivalently as:
//
//	actionID(binary)/buildID(main.a)/contentID(binary)
//
// Storing the buildID(main.a) in the middle lets the computations that care
// about the prefix or suffix halves ignore the middle and preserves the
// original build ID as a contiguous string.)
//
// During the build, when it's time to build main.a, the gofmt binary has the
// information needed to decide whether the eventual link would produce
// the same binary: if the action ID for main.a's inputs matches and then
// the action ID for the link step matches when assuming the given main.a
// content ID, then the binary as a whole is up-to-date and need not be rebuilt.
//
// This is all a bit complex and may be simplified once we can rely on the
// main cache, but at least at the start we will be using the content-based
// staleness determination without a cache beyond the usual installed
// package and binary locations.

const buildIDSeparator = "/"

// contentID returns the content ID half of a build ID.
func contentID(buildID string) string {
	return buildID[strings.LastIndex(buildID, buildIDSeparator)+1:]
}

// hashToString converts the hash h to a string to be recorded
// in package archives and binaries as part of the build ID.
// We use the first 96 bits of the hash and encode it in base64,
// resulting in a 16-byte string. Because this is only used for
// detecting the need to rebuild installed files (not for lookups
// in the object file cache), 96 bits are sufficient to drive the
// probability of a false "do not need to rebuild" decision to effectively zero.
// We embed two different hashes in archives and four in binaries,
// so cutting to 16 bytes is a significant savings when build IDs are displayed.
// (16*4+3 = 67 bytes compared to 64*4+3 = 259 bytes for the
// more straightforward option of printing the entire h in hex).
func hashToString(h [cache.HashSize]byte) string {
	const b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	const chunks = 5
	var dst [chunks * 4]byte
	for i := 0; i < chunks; i++ {
		v := uint32(h[3*i])<<16 | uint32(h[3*i+1])<<8 | uint32(h[3*i+2])
		dst[4*i+0] = b64[(v>>18)&0x3F]
		dst[4*i+1] = b64[(v>>12)&0x3F]
		dst[4*i+2] = b64[(v>>6)&0x3F]
		dst[4*i+3] = b64[v&0x3F]
	}
	return string(dst[:])
}

// toolID returns the unique ID to use for the current copy of the
// named tool (asm, compile, cover, link).
//
// It is important that if the tool changes (for example a compiler bug is fixed
// and the compiler reinstalled), toolID returns a different string, so that old
// package archives look stale and are rebuilt (with the fixed compiler).
// This suggests using a content hash of the tool binary, as stored in the build ID.
//
// Unfortunately, we can't just open the tool binary, because the tool might be
// invoked via a wrapper program specified by -toolexec and we don't know
// what the wrapper program does. In particular, we want "-toolexec toolstash"
// to continue working: it does no good if "-toolexec toolstash" is executing a
// stashed copy of the compiler but the go command is acting as if it will run
// the standard copy of the compiler. The solution is to ask the tool binary to tell
// us its own build ID using the "-V=full" flag now supported by all tools.
// Then we know we're getting the build ID of the compiler that will actually run
// during the build. (How does the compiler binary know its own content hash?
// We store it there using updateBuildID after the standard link step.)
//
// A final twist is that we'd prefer to have reproducible builds for release toolchains.
// It should be possible to cross-compile for Windows from either Linux or Mac
// or Windows itself and produce the same binaries, bit for bit. If the tool ID,
// which influences the action ID half of the build ID, is based on the content ID,
// then the Linux compiler binary and Mac compiler binary will have different tool IDs
// and therefore produce executables with different action IDs.
// To avoids this problem, for releases we use the release version string instead
// of the compiler binary's content hash. This assumes that all compilers built
// on all different systems are semantically equivalent, which is of course only true
// modulo bugs. (Producing the exact same executables also requires that the different
// build setups agree on details like $GOROOT and file name paths, but at least the
// tool IDs do not make it impossible.)
func (b *Builder) toolID(name string) string {
	b.id.Lock()
	id := b.toolIDCache[name]
	b.id.Unlock()

	if id != "" {
		return id
	}

	cmdline := str.StringList(cfg.BuildToolexec, base.Tool(name), "-V=full")
	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Env = base.EnvForDir(cmd.Dir, os.Environ())
	out, err := cmd.CombinedOutput()
	if err != nil {
		base.Fatalf("go tool %s: %v\n%s", name, err, out)
	}

	line := string(out)
	f := strings.Fields(line)
	if len(f) < 3 || f[0] != name || f[1] != "version" || f[2] == "devel" && !strings.HasPrefix(f[len(f)-1], "buildID=") {
		base.Fatalf("go tool %s -V=full: unexpected output:\n\t%s", name, line)
	}
	if f[2] == "devel" {
		// On the development branch, use the content ID part of the build ID.
		id = contentID(f[len(f)-1])
	} else {
		// For a release, the output is like: "compile version go1.9.1". Use the whole line.
		id = f[2]
	}

	b.id.Lock()
	b.toolIDCache[name] = id
	b.id.Unlock()

	return id
}

// buildID returns the build ID found in the given file.
// If no build ID is found, buildID returns the content hash of the file.
func (b *Builder) buildID(file string) string {
	b.id.Lock()
	id := b.buildIDCache[file]
	b.id.Unlock()

	if id != "" {
		return id
	}

	id, err := buildid.ReadFile(file)
	if err != nil {
		id = b.fileHash(file)
	}

	b.id.Lock()
	b.buildIDCache[file] = id
	b.id.Unlock()

	return id
}

// fileHash returns the content hash of the named file.
func (b *Builder) fileHash(file string) string {
	sum, err := cache.FileHash(file)
	if err != nil {
		return ""
	}
	return hashToString(sum)
}

// useCache tries to satisfy the action a, which has action ID actionHash,
// by using a cached result from an earlier build. At the moment, the only
// cached result is the installed package or binary at target.
// If useCache decides that the cache can be used, it sets a.buildID
// and a.built for use by parent actions and then returns true.
// Otherwise it sets a.buildID to a temporary build ID for use in the build
// and returns false. When useCache returns false the expectation is that
// the caller will build the target and then call updateBuildID to finish the
// build ID computation.
func (b *Builder) useCache(a *Action, p *load.Package, actionHash cache.ActionID, target string) bool {
	// The second half of the build ID here is a placeholder for the content hash.
	// It's important that the overall buildID be unlikely verging on impossible
	// to appear in the output by chance, but that should be taken care of by
	// the actionID half; if it also appeared in the input that would be like an
	// engineered 96-bit partial SHA256 collision.
	actionID := hashToString(actionHash)
	contentID := "(MISSING CONTENT ID)" // same length has hashToString result
	a.buildID = actionID + buildIDSeparator + contentID

	// Executable binaries also record the main build ID in the middle.
	// See "Build IDs" comment above.
	if a.Mode == "link" {
		mainpkg := a.Deps[0]
		a.buildID = actionID + buildIDSeparator + mainpkg.buildID + buildIDSeparator + contentID
	}

	// Check to see if target exists and matches the expected action ID.
	// If so, it's up to date and we can reuse it instead of rebuilding it.
	var buildID string
	if target != "" && !cfg.BuildA {
		var err error
		buildID, err = buildid.ReadFile(target)
		if err != nil && b.ComputeStaleOnly {
			if p != nil && !p.Stale {
				p.Stale = true
				p.StaleReason = "target missing"
			}
			return true
		}
		if strings.HasPrefix(buildID, actionID+buildIDSeparator) {
			a.buildID = buildID
			a.built = target
			// Poison a.Target to catch uses later in the build.
			a.Target = "DO NOT USE - " + a.Mode
			return true
		}
	}

	// Special case for building a main package: if the only thing we
	// want the package for is to link a binary, and the binary is
	// already up-to-date, then to avoid a rebuild, report the package
	// as up-to-date as well. See "Build IDs" comment above.
	if target != "" && !cfg.BuildA && a.Mode == "build" && len(a.triggers) == 1 && a.triggers[0].Mode == "link" {
		buildID, err := buildid.ReadFile(target)
		if err == nil {
			id := strings.Split(buildID, buildIDSeparator)
			if len(id) == 4 && id[1] == actionID {
				// Temporarily assume a.buildID is the package build ID
				// stored in the installed binary, and see if that makes
				// the upcoming link action ID a match. If so, report that
				// we built the package, safe in the knowledge that the
				// link step will not ask us for the actual package file.
				// Note that (*Builder).LinkAction arranged that all of
				// a.triggers[0]'s dependencies other than a are also
				// dependencies of a, so that we can be sure that,
				// other than a.buildID, b.linkActionID is only accessing
				// build IDs of completed actions.
				oldBuildID := a.buildID
				a.buildID = id[1] + buildIDSeparator + id[2]
				linkID := hashToString(b.linkActionID(a.triggers[0]))
				if id[0] == linkID {
					// Poison a.Target to catch uses later in the build.
					a.Target = "DO NOT USE - main build pseudo-cache Target"
					a.built = "DO NOT USE - main build pseudo-cache built"
					return true
				}
				// Otherwise restore old build ID for main build.
				a.buildID = oldBuildID
			}
		}
	}

	if b.ComputeStaleOnly {
		// Invoked during go list only to compute and record staleness.
		if p := a.Package; p != nil && !p.Stale {
			p.Stale = true
			if cfg.BuildA {
				p.StaleReason = "build -a flag in use"
			} else {
				p.StaleReason = "build ID mismatch"
				for _, p1 := range p.Internal.Imports {
					if p1.Stale && p1.StaleReason != "" {
						if strings.HasPrefix(p1.StaleReason, "stale dependency: ") {
							p.StaleReason = p1.StaleReason
							break
						}
						if strings.HasPrefix(p.StaleReason, "build ID mismatch") {
							p.StaleReason = "stale dependency: " + p1.ImportPath
						}
					}
				}
			}
		}
		return true
	}

	return false
}

// updateBuildID updates the build ID in the target written by action a.
// It requires that useCache was called for action a and returned false,
// and that the build was then carried out and given the temporary
// a.buildID to record as the build ID in the resulting package or binary.
// updateBuildID computes the final content ID and updates the build IDs
// in the binary.
func (b *Builder) updateBuildID(a *Action, target string) error {
	if cfg.BuildX || cfg.BuildN {
		b.Showcmd("", "%s # internal", joinUnambiguously(str.StringList(base.Tool("buildid"), "-w", target)))
		if cfg.BuildN {
			return nil
		}
	}

	// Find occurrences of old ID and compute new content-based ID.
	r, err := os.Open(target)
	if err != nil {
		return err
	}
	matches, hash, err := buildid.FindAndHash(r, a.buildID, 0)
	r.Close()
	if err != nil {
		return err
	}
	newID := a.buildID[:strings.LastIndex(a.buildID, buildIDSeparator)] + buildIDSeparator + hashToString(hash)
	if len(newID) != len(a.buildID) {
		return fmt.Errorf("internal error: build ID length mismatch %q vs %q", a.buildID, newID)
	}

	// Replace with new content-based ID.
	a.buildID = newID
	if len(matches) == 0 {
		// Assume the user specified -buildid= to override what we were going to choose.
		return nil
	}
	w, err := os.OpenFile(target, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	err = buildid.Rewrite(w, matches, newID)
	if err != nil {
		w.Close()
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
