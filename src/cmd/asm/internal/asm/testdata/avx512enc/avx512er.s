// Code generated by avx512test. DO NOT EDIT.

#include "../../../../../../runtime/textflag.h"

TEXT asmtest_avx512er(SB), NOSPLIT, $0
	//TODO: VEXP2PD Z17, K7, Z20                               // 62a2fd4fc8e1
	//TODO: VEXP2PD Z0, K7, Z20                                // 62e2fd4fc8e0
	//TODO: VEXP2PD Z17, K7, Z0                                // 62b2fd4fc8c1
	//TODO: VEXP2PD Z0, K7, Z0                                 // 62f2fd4fc8c0
	//TODO: VEXP2PD Z31, K2, Z17                               // 6282fd4ac8cf
	//TODO: VEXP2PD Z0, K2, Z17                                // 62e2fd4ac8c8
	//TODO: VEXP2PD (R14), K2, Z17                             // 62c2fd4ac80e
	//TODO: VEXP2PD -7(DI)(R8*8), K2, Z17                      // 62a2fd4ac88cc7f9ffffff
	//TODO: VEXP2PD Z31, K2, Z23                               // 6282fd4ac8ff
	//TODO: VEXP2PD Z0, K2, Z23                                // 62e2fd4ac8f8
	//TODO: VEXP2PD (R14), K2, Z23                             // 62c2fd4ac83e
	//TODO: VEXP2PD -7(DI)(R8*8), K2, Z23                      // 62a2fd4ac8bcc7f9ffffff
	//TODO: VEXP2PS Z6, K4, Z21                                // 62e27d4cc8ee
	//TODO: VEXP2PS Z9, K4, Z21                                // 62c27d4cc8e9
	//TODO: VEXP2PS Z6, K4, Z9                                 // 62727d4cc8ce
	//TODO: VEXP2PS Z9, K4, Z9                                 // 62527d4cc8c9
	//TODO: VEXP2PS Z20, K1, Z1                                // 62b27d49c8cc
	//TODO: VEXP2PS Z9, K1, Z1                                 // 62d27d49c8c9
	//TODO: VEXP2PS 99(R15)(R15*4), K1, Z1                     // 62927d49c88cbf63000000
	//TODO: VEXP2PS 15(DX), K1, Z1                             // 62f27d49c88a0f000000
	//TODO: VEXP2PS Z20, K1, Z9                                // 62327d49c8cc
	//TODO: VEXP2PS Z9, K1, Z9                                 // 62527d49c8c9
	//TODO: VEXP2PS 99(R15)(R15*4), K1, Z9                     // 62127d49c88cbf63000000
	//TODO: VEXP2PS 15(DX), K1, Z9                             // 62727d49c88a0f000000
	//TODO: VRCP28PD Z13, K7, Z11                              // 6252fd4fcadd
	//TODO: VRCP28PD Z14, K7, Z11                              // 6252fd4fcade
	//TODO: VRCP28PD Z13, K7, Z5                               // 62d2fd4fcaed
	//TODO: VRCP28PD Z14, K7, Z5                               // 62d2fd4fcaee
	//TODO: VRCP28PD Z2, K2, Z5                                // 62f2fd4acaea
	//TODO: VRCP28PD -7(CX)(DX*1), K2, Z5                      // 62f2fd4acaac11f9ffffff
	//TODO: VRCP28PD -15(R14)(R15*4), K2, Z5                   // 6292fd4acaacbef1ffffff
	//TODO: VRCP28PD Z2, K2, Z23                               // 62e2fd4acafa
	//TODO: VRCP28PD -7(CX)(DX*1), K2, Z23                     // 62e2fd4acabc11f9ffffff
	//TODO: VRCP28PD -15(R14)(R15*4), K2, Z23                  // 6282fd4acabcbef1ffffff
	//TODO: VRCP28PS Z26, K5, Z6                               // 62927d4dcaf2
	//TODO: VRCP28PS Z14, K5, Z6                               // 62d27d4dcaf6
	//TODO: VRCP28PS Z26, K5, Z14                              // 62127d4dcaf2
	//TODO: VRCP28PS Z14, K5, Z14                              // 62527d4dcaf6
	//TODO: VRCP28PS Z13, K3, Z28                              // 62427d4bcae5
	//TODO: VRCP28PS Z21, K3, Z28                              // 62227d4bcae5
	//TODO: VRCP28PS 15(DX)(BX*1), K3, Z28                     // 62627d4bcaa41a0f000000
	//TODO: VRCP28PS -7(CX)(DX*2), K3, Z28                     // 62627d4bcaa451f9ffffff
	//TODO: VRCP28PS Z13, K3, Z6                               // 62d27d4bcaf5
	//TODO: VRCP28PS Z21, K3, Z6                               // 62b27d4bcaf5
	//TODO: VRCP28PS 15(DX)(BX*1), K3, Z6                      // 62f27d4bcab41a0f000000
	//TODO: VRCP28PS -7(CX)(DX*2), K3, Z6                      // 62f27d4bcab451f9ffffff
	//TODO: VRCP28SD X25, X14, K4, X19                         // 62828d0ccbd9
	//TODO: VRCP28SD X11, X14, K4, X19                         // 62c28d0ccbdb
	//TODO: VRCP28SD X17, X14, K4, X19                         // 62a28d0ccbd9
	//TODO: VRCP28SD X25, X0, K4, X19                          // 6282fd0ccbd9
	//TODO: VRCP28SD X11, X0, K4, X19                          // 62c2fd0ccbdb
	//TODO: VRCP28SD X17, X0, K4, X19                          // 62a2fd0ccbd9
	//TODO: VRCP28SD X25, X14, K4, X13                         // 62128d0ccbe9
	//TODO: VRCP28SD X11, X14, K4, X13                         // 62528d0ccbeb
	//TODO: VRCP28SD X17, X14, K4, X13                         // 62328d0ccbe9
	//TODO: VRCP28SD X25, X0, K4, X13                          // 6212fd0ccbe9
	//TODO: VRCP28SD X11, X0, K4, X13                          // 6252fd0ccbeb
	//TODO: VRCP28SD X17, X0, K4, X13                          // 6232fd0ccbe9
	//TODO: VRCP28SD X25, X14, K4, X2                          // 62928d0ccbd1
	//TODO: VRCP28SD X11, X14, K4, X2                          // 62d28d0ccbd3
	//TODO: VRCP28SD X17, X14, K4, X2                          // 62b28d0ccbd1
	//TODO: VRCP28SD X25, X0, K4, X2                           // 6292fd0ccbd1
	//TODO: VRCP28SD X11, X0, K4, X2                           // 62d2fd0ccbd3
	//TODO: VRCP28SD X17, X0, K4, X2                           // 62b2fd0ccbd1
	//TODO: VRCP28SD X2, X2, K2, X18                           // 62e2ed0acbd2 or 62e2ed2acbd2 or 62e2ed4acbd2
	//TODO: VRCP28SD X27, X2, K2, X18                          // 6282ed0acbd3 or 6282ed2acbd3 or 6282ed4acbd3
	//TODO: VRCP28SD X26, X2, K2, X18                          // 6282ed0acbd2 or 6282ed2acbd2 or 6282ed4acbd2
	//TODO: VRCP28SD 17(SP)(BP*8), X2, K2, X18                 // 62e2ed0acb94ec11000000 or 62e2ed2acb94ec11000000 or 62e2ed4acb94ec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X2, K2, X18                 // 62e2ed0acb94ac11000000 or 62e2ed2acb94ac11000000 or 62e2ed4acb94ac11000000
	//TODO: VRCP28SD X2, X24, K2, X18                          // 62e2bd02cbd2 or 62e2bd22cbd2 or 62e2bd42cbd2
	//TODO: VRCP28SD X27, X24, K2, X18                         // 6282bd02cbd3 or 6282bd22cbd3 or 6282bd42cbd3
	//TODO: VRCP28SD X26, X24, K2, X18                         // 6282bd02cbd2 or 6282bd22cbd2 or 6282bd42cbd2
	//TODO: VRCP28SD 17(SP)(BP*8), X24, K2, X18                // 62e2bd02cb94ec11000000 or 62e2bd22cb94ec11000000 or 62e2bd42cb94ec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X24, K2, X18                // 62e2bd02cb94ac11000000 or 62e2bd22cb94ac11000000 or 62e2bd42cb94ac11000000
	//TODO: VRCP28SD X2, X2, K2, X11                           // 6272ed0acbda or 6272ed2acbda or 6272ed4acbda
	//TODO: VRCP28SD X27, X2, K2, X11                          // 6212ed0acbdb or 6212ed2acbdb or 6212ed4acbdb
	//TODO: VRCP28SD X26, X2, K2, X11                          // 6212ed0acbda or 6212ed2acbda or 6212ed4acbda
	//TODO: VRCP28SD 17(SP)(BP*8), X2, K2, X11                 // 6272ed0acb9cec11000000 or 6272ed2acb9cec11000000 or 6272ed4acb9cec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X2, K2, X11                 // 6272ed0acb9cac11000000 or 6272ed2acb9cac11000000 or 6272ed4acb9cac11000000
	//TODO: VRCP28SD X2, X24, K2, X11                          // 6272bd02cbda or 6272bd22cbda or 6272bd42cbda
	//TODO: VRCP28SD X27, X24, K2, X11                         // 6212bd02cbdb or 6212bd22cbdb or 6212bd42cbdb
	//TODO: VRCP28SD X26, X24, K2, X11                         // 6212bd02cbda or 6212bd22cbda or 6212bd42cbda
	//TODO: VRCP28SD 17(SP)(BP*8), X24, K2, X11                // 6272bd02cb9cec11000000 or 6272bd22cb9cec11000000 or 6272bd42cb9cec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X24, K2, X11                // 6272bd02cb9cac11000000 or 6272bd22cb9cac11000000 or 6272bd42cb9cac11000000
	//TODO: VRCP28SD X2, X2, K2, X9                            // 6272ed0acbca or 6272ed2acbca or 6272ed4acbca
	//TODO: VRCP28SD X27, X2, K2, X9                           // 6212ed0acbcb or 6212ed2acbcb or 6212ed4acbcb
	//TODO: VRCP28SD X26, X2, K2, X9                           // 6212ed0acbca or 6212ed2acbca or 6212ed4acbca
	//TODO: VRCP28SD 17(SP)(BP*8), X2, K2, X9                  // 6272ed0acb8cec11000000 or 6272ed2acb8cec11000000 or 6272ed4acb8cec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X2, K2, X9                  // 6272ed0acb8cac11000000 or 6272ed2acb8cac11000000 or 6272ed4acb8cac11000000
	//TODO: VRCP28SD X2, X24, K2, X9                           // 6272bd02cbca or 6272bd22cbca or 6272bd42cbca
	//TODO: VRCP28SD X27, X24, K2, X9                          // 6212bd02cbcb or 6212bd22cbcb or 6212bd42cbcb
	//TODO: VRCP28SD X26, X24, K2, X9                          // 6212bd02cbca or 6212bd22cbca or 6212bd42cbca
	//TODO: VRCP28SD 17(SP)(BP*8), X24, K2, X9                 // 6272bd02cb8cec11000000 or 6272bd22cb8cec11000000 or 6272bd42cb8cec11000000
	//TODO: VRCP28SD 17(SP)(BP*4), X24, K2, X9                 // 6272bd02cb8cac11000000 or 6272bd22cb8cac11000000 or 6272bd42cb8cac11000000
	//TODO: VRCP28SS X13, X11, K2, X22                         // 62c2250acbf5
	//TODO: VRCP28SS X6, X11, K2, X22                          // 62e2250acbf6
	//TODO: VRCP28SS X12, X11, K2, X22                         // 62c2250acbf4
	//TODO: VRCP28SS X13, X15, K2, X22                         // 62c2050acbf5
	//TODO: VRCP28SS X6, X15, K2, X22                          // 62e2050acbf6
	//TODO: VRCP28SS X12, X15, K2, X22                         // 62c2050acbf4
	//TODO: VRCP28SS X13, X30, K2, X22                         // 62c20d02cbf5
	//TODO: VRCP28SS X6, X30, K2, X22                          // 62e20d02cbf6
	//TODO: VRCP28SS X12, X30, K2, X22                         // 62c20d02cbf4
	//TODO: VRCP28SS X13, X11, K2, X30                         // 6242250acbf5
	//TODO: VRCP28SS X6, X11, K2, X30                          // 6262250acbf6
	//TODO: VRCP28SS X12, X11, K2, X30                         // 6242250acbf4
	//TODO: VRCP28SS X13, X15, K2, X30                         // 6242050acbf5
	//TODO: VRCP28SS X6, X15, K2, X30                          // 6262050acbf6
	//TODO: VRCP28SS X12, X15, K2, X30                         // 6242050acbf4
	//TODO: VRCP28SS X13, X30, K2, X30                         // 62420d02cbf5
	//TODO: VRCP28SS X6, X30, K2, X30                          // 62620d02cbf6
	//TODO: VRCP28SS X12, X30, K2, X30                         // 62420d02cbf4
	//TODO: VRCP28SS X13, X11, K2, X3                          // 62d2250acbdd
	//TODO: VRCP28SS X6, X11, K2, X3                           // 62f2250acbde
	//TODO: VRCP28SS X12, X11, K2, X3                          // 62d2250acbdc
	//TODO: VRCP28SS X13, X15, K2, X3                          // 62d2050acbdd
	//TODO: VRCP28SS X6, X15, K2, X3                           // 62f2050acbde
	//TODO: VRCP28SS X12, X15, K2, X3                          // 62d2050acbdc
	//TODO: VRCP28SS X13, X30, K2, X3                          // 62d20d02cbdd
	//TODO: VRCP28SS X6, X30, K2, X3                           // 62f20d02cbde
	//TODO: VRCP28SS X12, X30, K2, X3                          // 62d20d02cbdc
	//TODO: VRCP28SS X26, X20, K3, X23                         // 62825d03cbfa or 62825d23cbfa or 62825d43cbfa
	//TODO: VRCP28SS X19, X20, K3, X23                         // 62a25d03cbfb or 62a25d23cbfb or 62a25d43cbfb
	//TODO: VRCP28SS X0, X20, K3, X23                          // 62e25d03cbf8 or 62e25d23cbf8 or 62e25d43cbf8
	//TODO: VRCP28SS -7(CX), X20, K3, X23                      // 62e25d03cbb9f9ffffff or 62e25d23cbb9f9ffffff or 62e25d43cbb9f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X20, K3, X23                // 62e25d03cbbc9a0f000000 or 62e25d23cbbc9a0f000000 or 62e25d43cbbc9a0f000000
	//TODO: VRCP28SS X26, X2, K3, X23                          // 62826d0bcbfa or 62826d2bcbfa or 62826d4bcbfa
	//TODO: VRCP28SS X19, X2, K3, X23                          // 62a26d0bcbfb or 62a26d2bcbfb or 62a26d4bcbfb
	//TODO: VRCP28SS X0, X2, K3, X23                           // 62e26d0bcbf8 or 62e26d2bcbf8 or 62e26d4bcbf8
	//TODO: VRCP28SS -7(CX), X2, K3, X23                       // 62e26d0bcbb9f9ffffff or 62e26d2bcbb9f9ffffff or 62e26d4bcbb9f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X2, K3, X23                 // 62e26d0bcbbc9a0f000000 or 62e26d2bcbbc9a0f000000 or 62e26d4bcbbc9a0f000000
	//TODO: VRCP28SS X26, X9, K3, X23                          // 6282350bcbfa or 6282352bcbfa or 6282354bcbfa
	//TODO: VRCP28SS X19, X9, K3, X23                          // 62a2350bcbfb or 62a2352bcbfb or 62a2354bcbfb
	//TODO: VRCP28SS X0, X9, K3, X23                           // 62e2350bcbf8 or 62e2352bcbf8 or 62e2354bcbf8
	//TODO: VRCP28SS -7(CX), X9, K3, X23                       // 62e2350bcbb9f9ffffff or 62e2352bcbb9f9ffffff or 62e2354bcbb9f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X9, K3, X23                 // 62e2350bcbbc9a0f000000 or 62e2352bcbbc9a0f000000 or 62e2354bcbbc9a0f000000
	//TODO: VRCP28SS X26, X20, K3, X30                         // 62025d03cbf2 or 62025d23cbf2 or 62025d43cbf2
	//TODO: VRCP28SS X19, X20, K3, X30                         // 62225d03cbf3 or 62225d23cbf3 or 62225d43cbf3
	//TODO: VRCP28SS X0, X20, K3, X30                          // 62625d03cbf0 or 62625d23cbf0 or 62625d43cbf0
	//TODO: VRCP28SS -7(CX), X20, K3, X30                      // 62625d03cbb1f9ffffff or 62625d23cbb1f9ffffff or 62625d43cbb1f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X20, K3, X30                // 62625d03cbb49a0f000000 or 62625d23cbb49a0f000000 or 62625d43cbb49a0f000000
	//TODO: VRCP28SS X26, X2, K3, X30                          // 62026d0bcbf2 or 62026d2bcbf2 or 62026d4bcbf2
	//TODO: VRCP28SS X19, X2, K3, X30                          // 62226d0bcbf3 or 62226d2bcbf3 or 62226d4bcbf3
	//TODO: VRCP28SS X0, X2, K3, X30                           // 62626d0bcbf0 or 62626d2bcbf0 or 62626d4bcbf0
	//TODO: VRCP28SS -7(CX), X2, K3, X30                       // 62626d0bcbb1f9ffffff or 62626d2bcbb1f9ffffff or 62626d4bcbb1f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X2, K3, X30                 // 62626d0bcbb49a0f000000 or 62626d2bcbb49a0f000000 or 62626d4bcbb49a0f000000
	//TODO: VRCP28SS X26, X9, K3, X30                          // 6202350bcbf2 or 6202352bcbf2 or 6202354bcbf2
	//TODO: VRCP28SS X19, X9, K3, X30                          // 6222350bcbf3 or 6222352bcbf3 or 6222354bcbf3
	//TODO: VRCP28SS X0, X9, K3, X30                           // 6262350bcbf0 or 6262352bcbf0 or 6262354bcbf0
	//TODO: VRCP28SS -7(CX), X9, K3, X30                       // 6262350bcbb1f9ffffff or 6262352bcbb1f9ffffff or 6262354bcbb1f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X9, K3, X30                 // 6262350bcbb49a0f000000 or 6262352bcbb49a0f000000 or 6262354bcbb49a0f000000
	//TODO: VRCP28SS X26, X20, K3, X8                          // 62125d03cbc2 or 62125d23cbc2 or 62125d43cbc2
	//TODO: VRCP28SS X19, X20, K3, X8                          // 62325d03cbc3 or 62325d23cbc3 or 62325d43cbc3
	//TODO: VRCP28SS X0, X20, K3, X8                           // 62725d03cbc0 or 62725d23cbc0 or 62725d43cbc0
	//TODO: VRCP28SS -7(CX), X20, K3, X8                       // 62725d03cb81f9ffffff or 62725d23cb81f9ffffff or 62725d43cb81f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X20, K3, X8                 // 62725d03cb849a0f000000 or 62725d23cb849a0f000000 or 62725d43cb849a0f000000
	//TODO: VRCP28SS X26, X2, K3, X8                           // 62126d0bcbc2 or 62126d2bcbc2 or 62126d4bcbc2
	//TODO: VRCP28SS X19, X2, K3, X8                           // 62326d0bcbc3 or 62326d2bcbc3 or 62326d4bcbc3
	//TODO: VRCP28SS X0, X2, K3, X8                            // 62726d0bcbc0 or 62726d2bcbc0 or 62726d4bcbc0
	//TODO: VRCP28SS -7(CX), X2, K3, X8                        // 62726d0bcb81f9ffffff or 62726d2bcb81f9ffffff or 62726d4bcb81f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X2, K3, X8                  // 62726d0bcb849a0f000000 or 62726d2bcb849a0f000000 or 62726d4bcb849a0f000000
	//TODO: VRCP28SS X26, X9, K3, X8                           // 6212350bcbc2 or 6212352bcbc2 or 6212354bcbc2
	//TODO: VRCP28SS X19, X9, K3, X8                           // 6232350bcbc3 or 6232352bcbc3 or 6232354bcbc3
	//TODO: VRCP28SS X0, X9, K3, X8                            // 6272350bcbc0 or 6272352bcbc0 or 6272354bcbc0
	//TODO: VRCP28SS -7(CX), X9, K3, X8                        // 6272350bcb81f9ffffff or 6272352bcb81f9ffffff or 6272354bcb81f9ffffff
	//TODO: VRCP28SS 15(DX)(BX*4), X9, K3, X8                  // 6272350bcb849a0f000000 or 6272352bcb849a0f000000 or 6272354bcb849a0f000000
	//TODO: VRSQRT28PD Z7, K3, Z3                              // 62f2fd4bccdf
	//TODO: VRSQRT28PD Z9, K3, Z3                              // 62d2fd4bccd9
	//TODO: VRSQRT28PD Z7, K3, Z27                             // 6262fd4bccdf
	//TODO: VRSQRT28PD Z9, K3, Z27                             // 6242fd4bccd9
	//TODO: VRSQRT28PD Z20, K3, Z0                             // 62b2fd4bccc4
	//TODO: VRSQRT28PD Z28, K3, Z0                             // 6292fd4bccc4
	//TODO: VRSQRT28PD (SI), K3, Z0                            // 62f2fd4bcc06
	//TODO: VRSQRT28PD 7(SI)(DI*2), K3, Z0                     // 62f2fd4bcc847e07000000
	//TODO: VRSQRT28PD Z20, K3, Z6                             // 62b2fd4bccf4
	//TODO: VRSQRT28PD Z28, K3, Z6                             // 6292fd4bccf4
	//TODO: VRSQRT28PD (SI), K3, Z6                            // 62f2fd4bcc36
	//TODO: VRSQRT28PD 7(SI)(DI*2), K3, Z6                     // 62f2fd4bccb47e07000000
	//TODO: VRSQRT28PS Z9, K2, Z3                              // 62d27d4accd9
	//TODO: VRSQRT28PS Z19, K2, Z3                             // 62b27d4accdb
	//TODO: VRSQRT28PS Z9, K2, Z30                             // 62427d4accf1
	//TODO: VRSQRT28PS Z19, K2, Z30                            // 62227d4accf3
	//TODO: VRSQRT28PS Z11, K1, Z12                            // 62527d49cce3
	//TODO: VRSQRT28PS Z5, K1, Z12                             // 62727d49cce5
	//TODO: VRSQRT28PS 17(SP)(BP*8), K1, Z12                   // 62727d49cca4ec11000000
	//TODO: VRSQRT28PS 17(SP)(BP*4), K1, Z12                   // 62727d49cca4ac11000000
	//TODO: VRSQRT28PS Z11, K1, Z22                            // 62c27d49ccf3
	//TODO: VRSQRT28PS Z5, K1, Z22                             // 62e27d49ccf5
	//TODO: VRSQRT28PS 17(SP)(BP*8), K1, Z22                   // 62e27d49ccb4ec11000000
	//TODO: VRSQRT28PS 17(SP)(BP*4), K1, Z22                   // 62e27d49ccb4ac11000000
	//TODO: VRSQRT28SD X20, X20, K2, X31                       // 6222dd02cdfc
	//TODO: VRSQRT28SD X16, X20, K2, X31                       // 6222dd02cdf8
	//TODO: VRSQRT28SD X12, X20, K2, X31                       // 6242dd02cdfc
	//TODO: VRSQRT28SD X20, X24, K2, X31                       // 6222bd02cdfc
	//TODO: VRSQRT28SD X16, X24, K2, X31                       // 6222bd02cdf8
	//TODO: VRSQRT28SD X12, X24, K2, X31                       // 6242bd02cdfc
	//TODO: VRSQRT28SD X20, X7, K2, X31                        // 6222c50acdfc
	//TODO: VRSQRT28SD X16, X7, K2, X31                        // 6222c50acdf8
	//TODO: VRSQRT28SD X12, X7, K2, X31                        // 6242c50acdfc
	//TODO: VRSQRT28SD X20, X20, K2, X3                        // 62b2dd02cddc
	//TODO: VRSQRT28SD X16, X20, K2, X3                        // 62b2dd02cdd8
	//TODO: VRSQRT28SD X12, X20, K2, X3                        // 62d2dd02cddc
	//TODO: VRSQRT28SD X20, X24, K2, X3                        // 62b2bd02cddc
	//TODO: VRSQRT28SD X16, X24, K2, X3                        // 62b2bd02cdd8
	//TODO: VRSQRT28SD X12, X24, K2, X3                        // 62d2bd02cddc
	//TODO: VRSQRT28SD X20, X7, K2, X3                         // 62b2c50acddc
	//TODO: VRSQRT28SD X16, X7, K2, X3                         // 62b2c50acdd8
	//TODO: VRSQRT28SD X12, X7, K2, X3                         // 62d2c50acddc
	//TODO: VRSQRT28SD X20, X20, K2, X28                       // 6222dd02cde4
	//TODO: VRSQRT28SD X16, X20, K2, X28                       // 6222dd02cde0
	//TODO: VRSQRT28SD X12, X20, K2, X28                       // 6242dd02cde4
	//TODO: VRSQRT28SD X20, X24, K2, X28                       // 6222bd02cde4
	//TODO: VRSQRT28SD X16, X24, K2, X28                       // 6222bd02cde0
	//TODO: VRSQRT28SD X12, X24, K2, X28                       // 6242bd02cde4
	//TODO: VRSQRT28SD X20, X7, K2, X28                        // 6222c50acde4
	//TODO: VRSQRT28SD X16, X7, K2, X28                        // 6222c50acde0
	//TODO: VRSQRT28SD X12, X7, K2, X28                        // 6242c50acde4
	//TODO: VRSQRT28SD X8, X6, K1, X6                          // 62d2cd09cdf0 or 62d2cd29cdf0 or 62d2cd49cdf0
	//TODO: VRSQRT28SD X6, X6, K1, X6                          // 62f2cd09cdf6 or 62f2cd29cdf6 or 62f2cd49cdf6
	//TODO: VRSQRT28SD X0, X6, K1, X6                          // 62f2cd09cdf0 or 62f2cd29cdf0 or 62f2cd49cdf0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X6, K1, X6              // 6292cd09cdb43f63000000 or 6292cd29cdb43f63000000 or 6292cd49cdb43f63000000
	//TODO: VRSQRT28SD (DX), X6, K1, X6                        // 62f2cd09cd32 or 62f2cd29cd32 or 62f2cd49cd32
	//TODO: VRSQRT28SD X8, X1, K1, X6                          // 62d2f509cdf0 or 62d2f529cdf0 or 62d2f549cdf0
	//TODO: VRSQRT28SD X6, X1, K1, X6                          // 62f2f509cdf6 or 62f2f529cdf6 or 62f2f549cdf6
	//TODO: VRSQRT28SD X0, X1, K1, X6                          // 62f2f509cdf0 or 62f2f529cdf0 or 62f2f549cdf0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X1, K1, X6              // 6292f509cdb43f63000000 or 6292f529cdb43f63000000 or 6292f549cdb43f63000000
	//TODO: VRSQRT28SD (DX), X1, K1, X6                        // 62f2f509cd32 or 62f2f529cd32 or 62f2f549cd32
	//TODO: VRSQRT28SD X8, X8, K1, X6                          // 62d2bd09cdf0 or 62d2bd29cdf0 or 62d2bd49cdf0
	//TODO: VRSQRT28SD X6, X8, K1, X6                          // 62f2bd09cdf6 or 62f2bd29cdf6 or 62f2bd49cdf6
	//TODO: VRSQRT28SD X0, X8, K1, X6                          // 62f2bd09cdf0 or 62f2bd29cdf0 or 62f2bd49cdf0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X8, K1, X6              // 6292bd09cdb43f63000000 or 6292bd29cdb43f63000000 or 6292bd49cdb43f63000000
	//TODO: VRSQRT28SD (DX), X8, K1, X6                        // 62f2bd09cd32 or 62f2bd29cd32 or 62f2bd49cd32
	//TODO: VRSQRT28SD X8, X6, K1, X17                         // 62c2cd09cdc8 or 62c2cd29cdc8 or 62c2cd49cdc8
	//TODO: VRSQRT28SD X6, X6, K1, X17                         // 62e2cd09cdce or 62e2cd29cdce or 62e2cd49cdce
	//TODO: VRSQRT28SD X0, X6, K1, X17                         // 62e2cd09cdc8 or 62e2cd29cdc8 or 62e2cd49cdc8
	//TODO: VRSQRT28SD 99(R15)(R15*1), X6, K1, X17             // 6282cd09cd8c3f63000000 or 6282cd29cd8c3f63000000 or 6282cd49cd8c3f63000000
	//TODO: VRSQRT28SD (DX), X6, K1, X17                       // 62e2cd09cd0a or 62e2cd29cd0a or 62e2cd49cd0a
	//TODO: VRSQRT28SD X8, X1, K1, X17                         // 62c2f509cdc8 or 62c2f529cdc8 or 62c2f549cdc8
	//TODO: VRSQRT28SD X6, X1, K1, X17                         // 62e2f509cdce or 62e2f529cdce or 62e2f549cdce
	//TODO: VRSQRT28SD X0, X1, K1, X17                         // 62e2f509cdc8 or 62e2f529cdc8 or 62e2f549cdc8
	//TODO: VRSQRT28SD 99(R15)(R15*1), X1, K1, X17             // 6282f509cd8c3f63000000 or 6282f529cd8c3f63000000 or 6282f549cd8c3f63000000
	//TODO: VRSQRT28SD (DX), X1, K1, X17                       // 62e2f509cd0a or 62e2f529cd0a or 62e2f549cd0a
	//TODO: VRSQRT28SD X8, X8, K1, X17                         // 62c2bd09cdc8 or 62c2bd29cdc8 or 62c2bd49cdc8
	//TODO: VRSQRT28SD X6, X8, K1, X17                         // 62e2bd09cdce or 62e2bd29cdce or 62e2bd49cdce
	//TODO: VRSQRT28SD X0, X8, K1, X17                         // 62e2bd09cdc8 or 62e2bd29cdc8 or 62e2bd49cdc8
	//TODO: VRSQRT28SD 99(R15)(R15*1), X8, K1, X17             // 6282bd09cd8c3f63000000 or 6282bd29cd8c3f63000000 or 6282bd49cd8c3f63000000
	//TODO: VRSQRT28SD (DX), X8, K1, X17                       // 62e2bd09cd0a or 62e2bd29cd0a or 62e2bd49cd0a
	//TODO: VRSQRT28SD X8, X6, K1, X28                         // 6242cd09cde0 or 6242cd29cde0 or 6242cd49cde0
	//TODO: VRSQRT28SD X6, X6, K1, X28                         // 6262cd09cde6 or 6262cd29cde6 or 6262cd49cde6
	//TODO: VRSQRT28SD X0, X6, K1, X28                         // 6262cd09cde0 or 6262cd29cde0 or 6262cd49cde0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X6, K1, X28             // 6202cd09cda43f63000000 or 6202cd29cda43f63000000 or 6202cd49cda43f63000000
	//TODO: VRSQRT28SD (DX), X6, K1, X28                       // 6262cd09cd22 or 6262cd29cd22 or 6262cd49cd22
	//TODO: VRSQRT28SD X8, X1, K1, X28                         // 6242f509cde0 or 6242f529cde0 or 6242f549cde0
	//TODO: VRSQRT28SD X6, X1, K1, X28                         // 6262f509cde6 or 6262f529cde6 or 6262f549cde6
	//TODO: VRSQRT28SD X0, X1, K1, X28                         // 6262f509cde0 or 6262f529cde0 or 6262f549cde0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X1, K1, X28             // 6202f509cda43f63000000 or 6202f529cda43f63000000 or 6202f549cda43f63000000
	//TODO: VRSQRT28SD (DX), X1, K1, X28                       // 6262f509cd22 or 6262f529cd22 or 6262f549cd22
	//TODO: VRSQRT28SD X8, X8, K1, X28                         // 6242bd09cde0 or 6242bd29cde0 or 6242bd49cde0
	//TODO: VRSQRT28SD X6, X8, K1, X28                         // 6262bd09cde6 or 6262bd29cde6 or 6262bd49cde6
	//TODO: VRSQRT28SD X0, X8, K1, X28                         // 6262bd09cde0 or 6262bd29cde0 or 6262bd49cde0
	//TODO: VRSQRT28SD 99(R15)(R15*1), X8, K1, X28             // 6202bd09cda43f63000000 or 6202bd29cda43f63000000 or 6202bd49cda43f63000000
	//TODO: VRSQRT28SD (DX), X8, K1, X28                       // 6262bd09cd22 or 6262bd29cd22 or 6262bd49cd22
	//TODO: VRSQRT28SS X16, X6, K7, X11                        // 62324d0fcdd8
	//TODO: VRSQRT28SS X28, X6, K7, X11                        // 62124d0fcddc
	//TODO: VRSQRT28SS X8, X6, K7, X11                         // 62524d0fcdd8
	//TODO: VRSQRT28SS X16, X22, K7, X11                       // 62324d07cdd8
	//TODO: VRSQRT28SS X28, X22, K7, X11                       // 62124d07cddc
	//TODO: VRSQRT28SS X8, X22, K7, X11                        // 62524d07cdd8
	//TODO: VRSQRT28SS X16, X12, K7, X11                       // 62321d0fcdd8
	//TODO: VRSQRT28SS X28, X12, K7, X11                       // 62121d0fcddc
	//TODO: VRSQRT28SS X8, X12, K7, X11                        // 62521d0fcdd8
	//TODO: VRSQRT28SS X16, X6, K7, X16                        // 62a24d0fcdc0
	//TODO: VRSQRT28SS X28, X6, K7, X16                        // 62824d0fcdc4
	//TODO: VRSQRT28SS X8, X6, K7, X16                         // 62c24d0fcdc0
	//TODO: VRSQRT28SS X16, X22, K7, X16                       // 62a24d07cdc0
	//TODO: VRSQRT28SS X28, X22, K7, X16                       // 62824d07cdc4
	//TODO: VRSQRT28SS X8, X22, K7, X16                        // 62c24d07cdc0
	//TODO: VRSQRT28SS X16, X12, K7, X16                       // 62a21d0fcdc0
	//TODO: VRSQRT28SS X28, X12, K7, X16                       // 62821d0fcdc4
	//TODO: VRSQRT28SS X8, X12, K7, X16                        // 62c21d0fcdc0
	//TODO: VRSQRT28SS X16, X6, K7, X6                         // 62b24d0fcdf0
	//TODO: VRSQRT28SS X28, X6, K7, X6                         // 62924d0fcdf4
	//TODO: VRSQRT28SS X8, X6, K7, X6                          // 62d24d0fcdf0
	//TODO: VRSQRT28SS X16, X22, K7, X6                        // 62b24d07cdf0
	//TODO: VRSQRT28SS X28, X22, K7, X6                        // 62924d07cdf4
	//TODO: VRSQRT28SS X8, X22, K7, X6                         // 62d24d07cdf0
	//TODO: VRSQRT28SS X16, X12, K7, X6                        // 62b21d0fcdf0
	//TODO: VRSQRT28SS X28, X12, K7, X6                        // 62921d0fcdf4
	//TODO: VRSQRT28SS X8, X12, K7, X6                         // 62d21d0fcdf0
	//TODO: VRSQRT28SS X14, X19, K1, X15                       // 62526501cdfe or 62526521cdfe or 62526541cdfe
	//TODO: VRSQRT28SS X0, X19, K1, X15                        // 62726501cdf8 or 62726521cdf8 or 62726541cdf8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X19, K1, X15             // 62126501cdbcb00f000000 or 62126521cdbcb00f000000 or 62126541cdbcb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X19, K1, X15              // 62726501cdbc91f9ffffff or 62726521cdbc91f9ffffff or 62726541cdbc91f9ffffff
	//TODO: VRSQRT28SS X14, X13, K1, X15                       // 62521509cdfe or 62521529cdfe or 62521549cdfe
	//TODO: VRSQRT28SS X0, X13, K1, X15                        // 62721509cdf8 or 62721529cdf8 or 62721549cdf8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X13, K1, X15             // 62121509cdbcb00f000000 or 62121529cdbcb00f000000 or 62121549cdbcb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X13, K1, X15              // 62721509cdbc91f9ffffff or 62721529cdbc91f9ffffff or 62721549cdbc91f9ffffff
	//TODO: VRSQRT28SS X14, X2, K1, X15                        // 62526d09cdfe or 62526d29cdfe or 62526d49cdfe
	//TODO: VRSQRT28SS X0, X2, K1, X15                         // 62726d09cdf8 or 62726d29cdf8 or 62726d49cdf8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X2, K1, X15              // 62126d09cdbcb00f000000 or 62126d29cdbcb00f000000 or 62126d49cdbcb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X2, K1, X15               // 62726d09cdbc91f9ffffff or 62726d29cdbc91f9ffffff or 62726d49cdbc91f9ffffff
	//TODO: VRSQRT28SS X14, X19, K1, X11                       // 62526501cdde or 62526521cdde or 62526541cdde
	//TODO: VRSQRT28SS X0, X19, K1, X11                        // 62726501cdd8 or 62726521cdd8 or 62726541cdd8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X19, K1, X11             // 62126501cd9cb00f000000 or 62126521cd9cb00f000000 or 62126541cd9cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X19, K1, X11              // 62726501cd9c91f9ffffff or 62726521cd9c91f9ffffff or 62726541cd9c91f9ffffff
	//TODO: VRSQRT28SS X14, X13, K1, X11                       // 62521509cdde or 62521529cdde or 62521549cdde
	//TODO: VRSQRT28SS X0, X13, K1, X11                        // 62721509cdd8 or 62721529cdd8 or 62721549cdd8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X13, K1, X11             // 62121509cd9cb00f000000 or 62121529cd9cb00f000000 or 62121549cd9cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X13, K1, X11              // 62721509cd9c91f9ffffff or 62721529cd9c91f9ffffff or 62721549cd9c91f9ffffff
	//TODO: VRSQRT28SS X14, X2, K1, X11                        // 62526d09cdde or 62526d29cdde or 62526d49cdde
	//TODO: VRSQRT28SS X0, X2, K1, X11                         // 62726d09cdd8 or 62726d29cdd8 or 62726d49cdd8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X2, K1, X11              // 62126d09cd9cb00f000000 or 62126d29cd9cb00f000000 or 62126d49cd9cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X2, K1, X11               // 62726d09cd9c91f9ffffff or 62726d29cd9c91f9ffffff or 62726d49cd9c91f9ffffff
	//TODO: VRSQRT28SS X14, X19, K1, X1                        // 62d26501cdce or 62d26521cdce or 62d26541cdce
	//TODO: VRSQRT28SS X0, X19, K1, X1                         // 62f26501cdc8 or 62f26521cdc8 or 62f26541cdc8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X19, K1, X1              // 62926501cd8cb00f000000 or 62926521cd8cb00f000000 or 62926541cd8cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X19, K1, X1               // 62f26501cd8c91f9ffffff or 62f26521cd8c91f9ffffff or 62f26541cd8c91f9ffffff
	//TODO: VRSQRT28SS X14, X13, K1, X1                        // 62d21509cdce or 62d21529cdce or 62d21549cdce
	//TODO: VRSQRT28SS X0, X13, K1, X1                         // 62f21509cdc8 or 62f21529cdc8 or 62f21549cdc8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X13, K1, X1              // 62921509cd8cb00f000000 or 62921529cd8cb00f000000 or 62921549cd8cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X13, K1, X1               // 62f21509cd8c91f9ffffff or 62f21529cd8c91f9ffffff or 62f21549cd8c91f9ffffff
	//TODO: VRSQRT28SS X14, X2, K1, X1                         // 62d26d09cdce or 62d26d29cdce or 62d26d49cdce
	//TODO: VRSQRT28SS X0, X2, K1, X1                          // 62f26d09cdc8 or 62f26d29cdc8 or 62f26d49cdc8
	//TODO: VRSQRT28SS 15(R8)(R14*4), X2, K1, X1               // 62926d09cd8cb00f000000 or 62926d29cd8cb00f000000 or 62926d49cd8cb00f000000
	//TODO: VRSQRT28SS -7(CX)(DX*4), X2, K1, X1                // 62f26d09cd8c91f9ffffff or 62f26d29cd8c91f9ffffff or 62f26d49cd8c91f9ffffff
	RET
