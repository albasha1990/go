// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import "strings"

// A FuncFlag records bits about a function, passed to the runtime.
type FuncFlag uint8

// Note: This list must match the list in runtime/symtab.go.
const (
	FuncFlag_TOPFRAME = 1 << iota
	FuncFlag_SPWRITE
)

// A FuncID identifies particular functions that need to be treated
// specially by the runtime.
// Note that in some situations involving plugins, there may be multiple
// copies of a particular special runtime function.
type FuncID uint8

// Note: this list must match the list in runtime/symtab.go.
const (
	FuncID_normal FuncID = iota // not a special function
	FuncID_asmcgocall
	FuncID_asyncPreempt
	FuncID_cgocallback
	FuncID_debugCallV1
	FuncID_externalthreadhandler
	FuncID_gcBgMarkWorker
	FuncID_goexit
	FuncID_gogo
	FuncID_gopanic
	FuncID_handleAsyncEvent
	FuncID_jmpdefer
	FuncID_mcall
	FuncID_morestack
	FuncID_mstart
	FuncID_panicwrap
	FuncID_rt0_go
	FuncID_runfinq
	FuncID_runtime_main
	FuncID_sigpanic
	FuncID_systemstack
	FuncID_systemstack_switch
	FuncID_wrapper // any autogenerated code (hash/eq algorithms, method wrappers, etc.)
)

var funcIDs = map[string]FuncID{
	"asmcgocall":            FuncID_asmcgocall,
	"asyncPreempt":          FuncID_asyncPreempt,
	"cgocallback":           FuncID_cgocallback,
	"debugCallV1":           FuncID_debugCallV1,
	"externalthreadhandler": FuncID_externalthreadhandler,
	"gcBgMarkWorker":        FuncID_gcBgMarkWorker,
	"go":                    FuncID_rt0_go,
	"goexit":                FuncID_goexit,
	"gogo":                  FuncID_gogo,
	"gopanic":               FuncID_gopanic,
	"handleAsyncEvent":      FuncID_handleAsyncEvent,
	"jmpdefer":              FuncID_jmpdefer,
	"main":                  FuncID_runtime_main,
	"mcall":                 FuncID_mcall,
	"morestack":             FuncID_morestack,
	"mstart":                FuncID_mstart,
	"panicwrap":             FuncID_panicwrap,
	"runfinq":               FuncID_runfinq,
	"sigpanic":              FuncID_sigpanic,
	"switch":                FuncID_systemstack_switch,
	"systemstack":           FuncID_systemstack,

	// Don't show in call stack but otherwise not special.
	"deferreturn":       FuncID_wrapper,
	"runOpenDeferFrame": FuncID_wrapper,
	"reflectcallSave":   FuncID_wrapper,
	"deferCallSave":     FuncID_wrapper,
}

// Get the function ID for the named function in the named file.
// The function should be package-qualified.
func GetFuncID(name string, isWrapper bool) FuncID {
	if isWrapper {
		return FuncID_wrapper
	}
	if strings.HasPrefix(name, "runtime.") {
		if id, ok := funcIDs[name[len("runtime."):]]; ok {
			return id
		}
	}
	return FuncID_normal
}
