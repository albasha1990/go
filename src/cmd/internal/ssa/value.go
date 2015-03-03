// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"fmt"
	"strings"
)

// A Value represents a value in the SSA representation of the program.
// The ID and Type fields must not be modified.  The remainder may be modified
// if they preserve the value of the Value (e.g. changing a (mul 2 x) to an (add x x)).
type Value struct {
	// A unique identifier for the value.  For performance we allocate these IDs
	// densely starting at 0.  There is no guarantee that there won't be occasional holes, though.
	ID ID

	// The operation that computes this value.  See op.go.
	Op Op

	// The type of this value.  Normally this will be a Go type, but there
	// are a few other pseudo-types, see type.go.
	Type Type

	// Auxiliary info for this value.  The type of this information depends on the opcode (& type).
	Aux interface{}

	// Arguments of this value
	Args []*Value

	// Containing basic block
	Block *Block

	// Storage for the first two args
	argstorage [2]*Value
}

// Examples:
// Opcode          aux   args
//  OpAdd          nil      2
//  OpConstStr  string      0
//  OpConstInt   int64      0
//  OpAddcq      int64      1    amd64 op: v = arg[0] + constant

// short form print.  Just v#.
func (v *Value) String() string {
	return fmt.Sprintf("v%d", v.ID)
}

// long form print.  v# = opcode <type> [aux] args [: reg]
func (v *Value) LongString() string {
	s := fmt.Sprintf("v%d = %s", v.ID, strings.TrimPrefix(v.Op.String(), "Op"))
	s += " <" + v.Type.String() + ">"
	if v.Aux != nil {
		s += fmt.Sprintf(" [%v]", v.Aux)
	}
	for _, a := range v.Args {
		s += fmt.Sprintf(" %v", a)
	}
	r := v.Block.Func.RegAlloc
	if r != nil && r[v.ID] != nil {
		s += " : " + r[v.ID].Name()
	}
	return s
}

func (v *Value) AddArg(w *Value) {
	v.Args = append(v.Args, w)
}
func (v *Value) AddArgs(a ...*Value) {
	v.Args = append(v.Args, a...)
}
func (v *Value) SetArg(i int, w *Value) {
	v.Args[i] = w
}
func (v *Value) RemoveArg(i int) {
	copy(v.Args[i:], v.Args[i+1:])
	v.Args = v.Args[:len(v.Args)-1]
}
func (v *Value) SetArgs1(a *Value) {
	v.resetArgs()
	v.AddArg(a)
}
func (v *Value) SetArgs2(a *Value, b *Value) {
	v.resetArgs()
	v.AddArg(a)
	v.AddArg(b)
}

func (v *Value) resetArgs() {
	v.argstorage[0] = nil
	v.argstorage[1] = nil
	v.Args = v.argstorage[:0]
}

// CopyFrom converts v to be the same value as w.  v and w must
// have the same type.
func (v *Value) CopyFrom(w *Value) {
	if !typeIdentical(v.Type, w.Type) {
		panic("copyFrom with unequal types")
	}
	v.Op = w.Op
	v.Aux = w.Aux
	v.resetArgs()
	v.AddArgs(w.Args...)
}

// SetType sets the type of v.  v must not have had its type
// set yet (it must be TypeInvalid).
func (v *Value) SetType() {
	if v.Type != TypeInvalid {
		panic("setting type when it is already set")
	}
	opcodeTable[v.Op].typer(v)
}
