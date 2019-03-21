// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "go_tls.h"
#include "textflag.h"

// On darwin/arm, the runtime always use runtime/cgo
// for resolution. This will just exit with nominal
// exit code

TEXT runtime·res_nsearch_trampoline(SB),NOSPLIT,$0
    MOVW    $80, R0
    BL    libc_exit(SB)
    RET

TEXT runtime·res_ninit_trampoline(SB),NOSPLIT,$0
    MOVW    $81, R0
    BL    libc_exit(SB)
    RET
