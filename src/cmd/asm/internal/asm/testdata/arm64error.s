// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

TEXT errors(SB),$0
	AND	$1, RSP                                          // ERROR "illegal combination"
	ANDS	$1, R0, RSP                                      // ERROR "illegal combination"
	ADDSW	R7->32, R14, R13                                 // ERROR "shift amount out of range 0 to 31"
	ADD	R1.UXTB<<5, R2, R3                               // ERROR "shift amount out of range 0 to 4"
	ADDS	R1.UXTX<<7, R2, R3                               // ERROR "shift amount out of range 0 to 4"
	AND	$0x22220000, R2, RSP                             // ERROR "illegal combination"
	ANDS	$0x22220000, R2, RSP                             // ERROR "illegal combination"
	ADD	R1, R2, R3, R4                                   // ERROR "illegal combination"
	BICW	R7@>33, R5, R16                                  // ERROR "shift amount out of range 0 to 31"
	CINC	CS, R2, R3, R4                                   // ERROR "illegal combination"
	CSEL	LT, R1, R2                                       // ERROR "illegal combination"
	LDP.P	8(R2), (R2, R3)                                  // ERROR "constrained unpredictable behavior"
	LDP.W	8(R3), (R2, R3)                                  // ERROR "constrained unpredictable behavior"
	LDP	(R1), (R2, R2)                                   // ERROR "constrained unpredictable behavior"
	LDP	(R0), (F0, F1)                                   // ERROR "invalid register pair"
	LDP	(R0), (R3, ZR)                                   // ERROR "invalid register pair"
	LDXPW	(RSP), (R2, R2)                                  // ERROR "constrained unpredictable behavior"
	LDAXPW	(R5), (R2, R2)                                   // ERROR "constrained unpredictable behavior"
	MOVD.P	300(R2), R3                                      // ERROR "offset out of range [-255,254]"
	MOVD.P	R3, 344(R2)                                      // ERROR "offset out of range [-255,254]"
	MOVD	(R3)(R7.SXTX<<2), R8                             // ERROR "invalid index shift amount"
	MOVWU	(R5)(R4.UXTW<<3), R10                            // ERROR "invalid index shift amount"
	MOVWU	(R5)(R4<<1), R10                                 // ERROR "invalid index shift amount"
	MOVB	(R5)(R4.SXTW<<5), R10                            // ERROR "invalid index shift amount"
	MOVH	R5, (R6)(R2<<3)                                  // ERROR "invalid index shift amount"
	MADD	R1, R2, R3                                       // ERROR "illegal combination"
	MOVD.P	R1, 8(R1)                                        // ERROR "constrained unpredictable behavior"
	MOVD.W 	16(R2), R2                                       // ERROR "constrained unpredictable behavior"
	STP	(F2, F3), (R0)                                   // ERROR "invalid register pair"
	STP.W	(R1, R2), 8(R1)                                  // ERROR "constrained unpredictable behavior"
	STP.P	(R1, R2), 8(R2)                                  // ERROR "constrained unpredictable behavior"
	STLXP	(R6, R11), (RSP), R6                             // ERROR "constrained unpredictable behavior"
	STXP	(R6, R11), (R2), R2                              // ERROR "constrained unpredictable behavior"
	STLXR	R3, (RSP), R3                                    // ERROR "constrained unpredictable behavior"
	STXR	R3, (R4), R4                                     // ERROR "constrained unpredictable behavior"
	STLXRB	R2, (R5), R5                                     // ERROR "constrained unpredictable behavior"
	VLD1	(R8)(R13), [V2.B16]                              // ERROR "illegal combination"
	VLD1	8(R9), [V2.B16]                                  // ERROR "illegal combination"
	VST1	[V1.B16], (R8)(R13)                              // ERROR "illegal combination"
	VST1	[V1.B16], 9(R2)                                  // ERROR "illegal combination"
	VLD1	8(R8)(R13), [V2.B16]                             // ERROR "illegal combination"
	VMOV	V8.D[2], V12.D[1]                                // ERROR "register element index out of range 0 to 1"
	VMOV	V8.S[4], V12.S[1]                                // ERROR "register element index out of range 0 to 3"
	VMOV	V8.H[8], V12.H[1]                                // ERROR "register element index out of range 0 to 7"
	VMOV	V8.B[16], V12.B[1]                               // ERROR "register element index out of range 0 to 15"
	VMOV	V8.D[0], V12.S[1]                                // ERROR "operand mismatch"
	VMOV	V8.D[0], V12.H[1]                                // ERROR "operand mismatch"
	VMOV	V8.D[0], V12.B[1]                                // ERROR "operand mismatch"
	VMOV	V8.S[0], V12.H[1]                                // ERROR "operand mismatch"
	VMOV	V8.S[0], V12.B[1]                                // ERROR "operand mismatch"
	VMOV	V8.H[0], V12.B[1]                                // ERROR "operand mismatch"
	VMOV	V8.B[16], R3                                     // ERROR "register element index out of range 0 to 15"
	VMOV	V8.H[9], R3                                      // ERROR "register element index out of range 0 to 7"
	VMOV	V8.S[4], R3                                      // ERROR "register element index out of range 0 to 3"
	VMOV	V8.D[2], R3                                      // ERROR "register element index out of range 0 to 1"
	VDUP	V8.B[16], R3.B16                                 // ERROR "register element index out of range 0 to 15"
	VDUP	V8.B[17], R3.B8                                  // ERROR "register element index out of range 0 to 15"
	VDUP	V8.H[9], R3.H4                                   // ERROR "register element index out of range 0 to 7"
	VDUP	V8.H[9], R3.H8                                   // ERROR "register element index out of range 0 to 7"
	VDUP	V8.S[4], R3.S2                                   // ERROR "register element index out of range 0 to 3"
	VDUP	V8.S[4], R3.S4                                   // ERROR "register element index out of range 0 to 3"
	VDUP	V8.D[2], R3.D2                                   // ERROR "register element index out of range 0 to 1"
	VFMLA	V1.D2, V12.D2, V3.S2                             // ERROR "operand mismatch"
	VFMLA	V1.S2, V12.S2, V3.D2                             // ERROR "operand mismatch"
	VFMLA	V1.S4, V12.S2, V3.D2                             // ERROR "operand mismatch"
	VFMLA	V1.H4, V12.H4, V3.D2                             // ERROR "operand mismatch"
	VFMLS	V1.S2, V12.S2, V3.S4                             // ERROR "operand mismatch"
	VFMLS	V1.S2, V12.D2, V3.S4                             // ERROR "operand mismatch"
	VFMLS	V1.S2, V12.S4, V3.D2                             // ERROR "operand mismatch"
	VFMLA	V1.B8, V12.B8, V3.B8                             // ERROR "invalid arrangement"
	VFMLA	V1.B16, V12.B16, V3.B16                          // ERROR "invalid arrangement"
	VFMLA	V1.H4, V12.H4, V3.H4                             // ERROR "invalid arrangement"
	VFMLA	V1.H8, V12.H8, V3.H8                             // ERROR "invalid arrangement"
	VFMLA	V1.H4, V12.H4, V3.H4                             // ERROR "invalid arrangement"
	VFMLS	V1.B8, V12.B8, V3.B8                             // ERROR "invalid arrangement"
	VFMLS	V1.B16, V12.B16, V3.B16                          // ERROR "invalid arrangement"
	VFMLS	V1.H4, V12.H4, V3.H4                             // ERROR "invalid arrangement"
	VFMLS	V1.H8, V12.H8, V3.H8                             // ERROR "invalid arrangement"
	VFMLS	V1.H4, V12.H4, V3.H4                             // ERROR "invalid arrangement"
	VST1.P	[V4.S4,V5.S4], 48(R1)                            // ERROR "invalid post-increment offset"
	VST1.P	[V4.S4], 8(R1)                                   // ERROR "invalid post-increment offset"
	VLD1.P	32(R1), [V8.S4, V9.S4, V10.S4]                   // ERROR "invalid post-increment offset"
	VLD1.P	48(R1), [V7.S4, V8.S4, V9.S4, V10.S4]            // ERROR "invalid post-increment offset"
	VPMULL	V1.D1, V2.H4, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL	V1.H4, V2.H4, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL	V1.D2, V2.D2, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL	V1.B16, V2.B16, V3.H8                            // ERROR "invalid arrangement"
	VPMULL2	V1.D2, V2.H4, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL2	V1.H4, V2.H4, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL2	V1.D1, V2.D1, V3.Q1                              // ERROR "invalid arrangement"
	VPMULL2	V1.B8, V2.B8, V3.H8                              // ERROR "invalid arrangement"
	VEXT	$8, V1.B16, V2.B8, V2.B16                        // ERROR "invalid arrangement"
	VEXT	$8, V1.H8, V2.H8, V2.H8                          // ERROR "invalid arrangement"
	VRBIT	V1.B16, V2.B8                                    // ERROR "invalid arrangement"
	VRBIT	V1.H4, V2.H4                                     // ERROR "invalid arrangement"
	VUSHR	$56, V1.D2, V2.H4                                // ERROR "invalid arrangement"
	VUSHR	$127, V1.D2, V2.D2                               // ERROR "shift out of range"
	VLD1.P	(R8)(R9.SXTX<<2), [V2.B16]                       // ERROR "invalid extended register"
	VLD1.P	(R8)(R9<<2), [V2.B16]                            // ERROR "invalid extended register"
	VST1.P	[V1.B16], (R8)(R9.UXTW)                          // ERROR "invalid extended register"
	VST1.P	[V1.B16], (R8)(R9<<1)                            // ERROR "invalid extended register"
	VREV64	V1.H4, V2.H8                                     // ERROR "invalid arrangement"
	VREV64	V1.D1, V2.D1                                     // ERROR "invalid arrangement"
	VREV16	V1.D1, V2.D1                                     // ERROR "invalid arrangement"
	VREV16	V1.B8, V2.B16                                    // ERROR "invalid arrangement"
	VREV16	V1.H4, V2.H4                                     // ERROR "invalid arrangement"
	FLDPD	(R0), (R1, R2)                                   // ERROR "invalid register pair"
	FLDPD	(R1), (F2, F2)                                   // ERROR "constrained unpredictable behavior"
	FLDPS	(R2), (F3, F3)                                   // ERROR "constrained unpredictable behavior"
	FSTPD	(R1, R2), (R0)                                   // ERROR "invalid register pair"
	FMOVS	(F2), F0                                         // ERROR "illegal combination"
	FMOVD	F0, (F1)                                         // ERROR "illegal combination"
	LDADDD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDLD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDLW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDLH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDLB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDLD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDLW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDLH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDANDLB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORLD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORLW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORLH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDEORLB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORLD	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORLW	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORLH	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDORLB	R5, (R6), ZR                                     // ERROR "illegal destination register"
	LDADDAD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDAW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDAH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDAB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDALD	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDADDALW	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDADDALH	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDADDALB	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDADDD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDLD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDLW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDLH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDADDLB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDAD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDAW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDAH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDAB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDALD	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDANDALW	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDANDALH	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDANDALB	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDANDD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDLD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDLW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDLH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDANDLB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORAD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORAW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORAH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORAB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORALD	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDEORALW	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDEORALH	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDEORALB	R5, (R6), RSP                            // ERROR "illegal destination register"
	LDEORD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORLD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORLW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORLH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDEORLB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORAD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORAW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORAH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORAB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORALD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORALW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORALH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORALB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORLD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORLW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORLH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	LDORLB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPAD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPAW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPAH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPAB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPALD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPALW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPALH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPALB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPLD	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPLW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPLH	R5, (R6), RSP                                    // ERROR "illegal destination register"
	SWPLB	R5, (R6), RSP                                    // ERROR "illegal destination register"
	STXR	R5, (R6), RSP                                    // ERROR "illegal destination register"
	STXRW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	STLXR	R5, (R6), RSP                                    // ERROR "illegal destination register"
	STLXRW	R5, (R6), RSP                                    // ERROR "illegal destination register"
	STXP	(R5, R7), (R6), RSP                              // ERROR "illegal destination register"
	STXPW	(R5, R7), (R6), RSP                              // ERROR "illegal destination register"
	STLXP	(R5, R7), (R6), RSP                              // ERROR "illegal destination register"
	STLXP	(R5, R7), (R6), RSP                              // ERROR "illegal destination register"
	MSR	OSLAR_EL1, R5                                    // ERROR "illegal combination"
	MRS	R11, AIDR_EL1                                    // ERROR "illegal combination"
	MSR	R6, AIDR_EL1                                     // ERROR "system register is not writable"
	MSR	R6, AMCFGR_EL0                                   // ERROR "system register is not writable"
	MSR	R6, AMCGCR_EL0                                   // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER00_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER01_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER02_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER03_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER04_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER05_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER06_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER07_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER08_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER09_EL0                              // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER010_EL0                             // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER011_EL0                             // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER012_EL0                             // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER013_EL0                             // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER014_EL0                             // ERROR "system register is not writable"
	MSR	R6, AMEVTYPER015_EL0                             // ERROR "system register is not writable"
	MSR	R6, CCSIDR2_EL1                                  // ERROR "system register is not writable"
	MSR	R6, CCSIDR_EL1                                   // ERROR "system register is not writable"
	MSR	R6, CLIDR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, CNTPCT_EL0                                   // ERROR "system register is not writable"
	MSR	R6, CNTVCT_EL0                                   // ERROR "system register is not writable"
	MSR	R6, CTR_EL0                                      // ERROR "system register is not writable"
	MSR	R6, CurrentEL                                    // ERROR "system register is not writable"
	MSR	R6, DBGAUTHSTATUS_EL1                            // ERROR "system register is not writable"
	MSR	R6, DBGDTRRX_EL0                                 // ERROR "system register is not writable"
	MSR	R6, DCZID_EL0                                    // ERROR "system register is not writable"
	MSR	R6, ERRIDR_EL1                                   // ERROR "system register is not writable"
	MSR	R6, ERXFR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, ERXPFGF_EL1                                  // ERROR "system register is not writable"
	MSR	R6, GMID_EL1                                     // ERROR "system register is not writable"
	MSR	R6, ICC_HPPIR0_EL1                               // ERROR "system register is not writable"
	MSR	R6, ICC_HPPIR1_EL1                               // ERROR "system register is not writable"
	MSR	R6, ICC_IAR0_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ICC_IAR1_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ICC_RPR_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ICV_HPPIR0_EL1                               // ERROR "system register is not writable"
	MSR	R6, ICV_HPPIR1_EL1                               // ERROR "system register is not writable"
	MSR	R6, ICV_IAR0_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ICV_IAR1_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ICV_RPR_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ID_AA64AFR0_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64AFR1_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64DFR0_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64DFR1_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64ISAR0_EL1                             // ERROR "system register is not writable"
	MSR	R6, ID_AA64ISAR1_EL1                             // ERROR "system register is not writable"
	MSR	R6, ID_AA64MMFR0_EL1                             // ERROR "system register is not writable"
	MSR	R6, ID_AA64MMFR1_EL1                             // ERROR "system register is not writable"
	MSR	R6, ID_AA64MMFR2_EL1                             // ERROR "system register is not writable"
	MSR	R6, ID_AA64PFR0_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64PFR1_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AA64ZFR0_EL1                              // ERROR "system register is not writable"
	MSR	R6, ID_AFR0_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ID_DFR0_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ID_ISAR0_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR1_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR2_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR3_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR4_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR5_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_ISAR6_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_MMFR0_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_MMFR1_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_MMFR2_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_MMFR3_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_MMFR4_EL1                                 // ERROR "system register is not writable"
	MSR	R6, ID_PFR0_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ID_PFR1_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ID_PFR2_EL1                                  // ERROR "system register is not writable"
	MSR	R6, ISR_EL1                                      // ERROR "system register is not writable"
	MSR	R6, LORID_EL1                                    // ERROR "system register is not writable"
	MSR	R6, MDCCSR_EL0                                   // ERROR "system register is not writable"
	MSR	R6, MDRAR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, MIDR_EL1                                     // ERROR "system register is not writable"
	MSR	R6, MPAMIDR_EL1                                  // ERROR "system register is not writable"
	MSR	R6, MPIDR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, MVFR0_EL1                                    // ERROR "system register is not writable"
	MSR	R6, MVFR1_EL1                                    // ERROR "system register is not writable"
	MSR	R6, MVFR2_EL1                                    // ERROR "system register is not writable"
	MSR	R6, OSLSR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, PMBIDR_EL1                                   // ERROR "system register is not writable"
	MSR	R6, PMCEID0_EL0                                  // ERROR "system register is not writable"
	MSR	R6, PMCEID1_EL0                                  // ERROR "system register is not writable"
	MSR	R6, PMMIR_EL1                                    // ERROR "system register is not writable"
	MSR	R6, PMSIDR_EL1                                   // ERROR "system register is not writable"
	MSR	R6, REVIDR_EL1                                   // ERROR "system register is not writable"
	MSR	R6, RNDR                                         // ERROR "system register is not writable"
	MRS	DBGDTRTX_EL0, R5                                 // ERROR "system register is not readable"
	MRS	ICV_DIR_EL1, R5                                  // ERROR "system register is not readable"
	MRS	ICC_SGI1R_EL1, R5                                // ERROR "system register is not readable"
	MRS	ICC_SGI0R_EL1, R5                                // ERROR "system register is not readable"
	MRS	ICC_EOIR1_EL1, R5                                // ERROR "system register is not readable"
	MRS	ICC_EOIR0_EL1, R5                                // ERROR "system register is not readable"
	MRS	ICC_DIR_EL1, R5                                  // ERROR "system register is not readable"
	MRS	ICC_ASGI1R_EL1, R5                               // ERROR "system register is not readable"
	MRS	ICV_EOIR0_EL1, R3                                // ERROR "system register is not readable"
	MRS	ICV_EOIR1_EL1, R3                                // ERROR "system register is not readable"
	MRS	PMSWINC_EL0, R3                                  // ERROR "system register is not readable"
	MRS	OSLAR_EL1, R3                                    // ERROR "system register is not readable"
	VLD3R.P	24(R15), [V15.H4,V16.H4,V17.H4]                  // ERROR "invalid post-increment offset"
	VBIT	V1.H4, V12.H4, V3.H4                             // ERROR "invalid arrangement"
	VBSL	V1.D2, V12.D2, V3.D2                             // ERROR "invalid arrangement"
	VUXTL	V30.D2, V30.H8                                   // ERROR "operand mismatch"
	VUXTL2	V20.B8, V21.H8                                   // ERROR "operand mismatch"
	VUXTL	V3.D2, V4.B8                                     // ERROR "operand mismatch"
	RET
