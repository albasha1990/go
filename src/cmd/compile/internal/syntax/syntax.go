// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import (
	"cmd/internal/src"
	"fmt"
	"io"
	"os"
)

// Mode describes the parser mode.
type Mode uint

// Modes supported by the parser.
const (
	CheckBranches Mode = 1 << iota // check correct use of labels, break, continue, and goto statements
)

// Error describes a syntax error. Error implements the error interface.
type Error struct {
	Pos src.Pos
	Msg string
}

func (err Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Pos, err.Msg)
}

var _ error = Error{} // verify that Error implements error

// An ErrorHandler is called for each error encountered reading a .go file.
type ErrorHandler func(err error)

// A Pragma value is a set of flags that augment a function or
// type declaration. Callers may assign meaning to the flags as
// appropriate.
type Pragma uint16

// A PragmaHandler is used to process //go: directives as
// they're scanned. The returned Pragma value will be unioned into the
// next FuncDecl node.
type PragmaHandler func(pos src.Pos, text string) Pragma

// A FilenameHandler is used to process each filename encountered
// in //line directives. The returned value is used as the absolute filename.
type FilenameHandler func(name string) string

// Parse parses a single Go source file from src and returns the corresponding
// syntax tree. If there are errors, Parse will return the first error found,
// and a possibly partially constructed syntax tree, or nil.
//
// If errh != nil, it is called with each error encountered, and Parse will
// process as much source as possible. In this case, the returned syntax tree
// is only nil if no correct package clause was found.
// If errh is nil, Parse will terminate immediately upon encountering the first
// error, and the returned syntax tree is nil.
//
// If pragh != nil, it is called with each pragma encountered.
//
// If fileh != nil, it is called to process each filename
// encountered in //line directives.
//
func Parse(base *src.PosBase, src io.Reader, errh ErrorHandler, pragh PragmaHandler, fileh FilenameHandler, mode Mode) (_ *File, first error) {
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(Error); ok {
				first = err
				return
			}
			panic(p)
		}
	}()

	var p parser
	p.init(base, src, errh, pragh, fileh, mode)
	p.next()
	return p.fileOrNil(), p.first
}

// ParseFile behaves like Parse but it reads the source from the named file.
func ParseFile(filename string, errh ErrorHandler, pragh PragmaHandler, mode Mode) (*File, error) {
	f, err := os.Open(filename)
	if err != nil {
		if errh != nil {
			errh(err)
		}
		return nil, err
	}
	defer f.Close()
	return Parse(src.NewFileBase(filename, filename), f, errh, pragh, nil, mode)
}
