// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by go run make_examples.go. DO NOT EDIT.

package bits_test

import (
	"fmt"
	"math/bits"
)

func ExampleLeadingZeros8() {
	fmt.Printf("LeadingZeros8(%08b) = %d\n", 1, bits.LeadingZeros8(1))
	// Output:
	// LeadingZeros8(00000001) = 7
}

func ExampleLeadingZeros16() {
	fmt.Printf("LeadingZeros16(%016b) = %d\n", 1, bits.LeadingZeros16(1))
	// Output:
	// LeadingZeros16(0000000000000001) = 15
}

func ExampleLeadingZeros32() {
	fmt.Printf("LeadingZeros32(%032b) = %d\n", 1, bits.LeadingZeros32(1))
	// Output:
	// LeadingZeros32(00000000000000000000000000000001) = 31
}

func ExampleLeadingZeros64() {
	fmt.Printf("LeadingZeros64(%064b) = %d\n", 1, bits.LeadingZeros64(1))
	// Output:
	// LeadingZeros64(0000000000000000000000000000000000000000000000000000000000000001) = 63
}

func ExampleTrailingZeros8() {
	fmt.Printf("TrailingZeros8(%08b) = %d\n", 14, bits.TrailingZeros8(14))
	// Output:
	// TrailingZeros8(00001110) = 1
}

func ExampleTrailingZeros16() {
	fmt.Printf("TrailingZeros16(%016b) = %d\n", 14, bits.TrailingZeros16(14))
	// Output:
	// TrailingZeros16(0000000000001110) = 1
}

func ExampleTrailingZeros32() {
	fmt.Printf("TrailingZeros32(%032b) = %d\n", 14, bits.TrailingZeros32(14))
	// Output:
	// TrailingZeros32(00000000000000000000000000001110) = 1
}

func ExampleTrailingZeros64() {
	fmt.Printf("TrailingZeros64(%064b) = %d\n", 14, bits.TrailingZeros64(14))
	// Output:
	// TrailingZeros64(0000000000000000000000000000000000000000000000000000000000001110) = 1
}

func ExampleOnesCount8() {
	fmt.Printf("OnesCount8(%08b) = %d\n", 14, bits.OnesCount8(14))
	// Output:
	// OnesCount8(00001110) = 3
}

func ExampleOnesCount16() {
	fmt.Printf("OnesCount16(%016b) = %d\n", 14, bits.OnesCount16(14))
	// Output:
	// OnesCount16(0000000000001110) = 3
}

func ExampleOnesCount32() {
	fmt.Printf("OnesCount32(%032b) = %d\n", 14, bits.OnesCount32(14))
	// Output:
	// OnesCount32(00000000000000000000000000001110) = 3
}

func ExampleOnesCount64() {
	fmt.Printf("OnesCount64(%064b) = %d\n", 14, bits.OnesCount64(14))
	// Output:
	// OnesCount64(0000000000000000000000000000000000000000000000000000000000001110) = 3
}

func ExampleRotateLeft8() {
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", bits.RotateLeft8(15, 2))
	// Output:
	// 00001111
	// 00111100
}

func ExampleRotateLeft16() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.RotateLeft16(15, 2))
	// Output:
	// 0000000000001111
	// 0000000000111100
}

func ExampleRotateLeft32() {
	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.RotateLeft32(15, 2))
	// Output:
	// 00000000000000000000000000001111
	// 00000000000000000000000000111100
}

func ExampleRotateLeft64() {
	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.RotateLeft64(15, 2))
	// Output:
	// 0000000000000000000000000000000000000000000000000000000000001111
	// 0000000000000000000000000000000000000000000000000000000000111100
}

func ExampleReverse8() {
	fmt.Printf("%08b\n", 19)
	fmt.Printf("%08b\n", bits.Reverse8(19))
	// Output:
	// 00010011
	// 11001000
}

func ExampleReverse16() {
	fmt.Printf("%016b\n", 19)
	fmt.Printf("%016b\n", bits.Reverse16(19))
	// Output:
	// 0000000000010011
	// 1100100000000000
}

func ExampleReverse32() {
	fmt.Printf("%032b\n", 19)
	fmt.Printf("%032b\n", bits.Reverse32(19))
	// Output:
	// 00000000000000000000000000010011
	// 11001000000000000000000000000000
}

func ExampleReverse64() {
	fmt.Printf("%064b\n", 19)
	fmt.Printf("%064b\n", bits.Reverse64(19))
	// Output:
	// 0000000000000000000000000000000000000000000000000000000000010011
	// 1100100000000000000000000000000000000000000000000000000000000000
}

func ExampleReverseBytes16() {
	fmt.Printf("%016b\n", 15)
	fmt.Printf("%016b\n", bits.ReverseBytes16(15))
	// Output:
	// 0000000000001111
	// 0000111100000000
}

func ExampleReverseBytes32() {
	fmt.Printf("%032b\n", 15)
	fmt.Printf("%032b\n", bits.ReverseBytes32(15))
	// Output:
	// 00000000000000000000000000001111
	// 00001111000000000000000000000000
}

func ExampleReverseBytes64() {
	fmt.Printf("%064b\n", 15)
	fmt.Printf("%064b\n", bits.ReverseBytes64(15))
	// Output:
	// 0000000000000000000000000000000000000000000000000000000000001111
	// 0000111100000000000000000000000000000000000000000000000000000000
}

func ExampleLen8() {
	fmt.Printf("Len8(%08b) = %d\n", 8, bits.Len8(8))
	// Output:
	// Len8(00001000) = 4
}

func ExampleLen16() {
	fmt.Printf("Len16(%016b) = %d\n", 8, bits.Len16(8))
	// Output:
	// Len16(0000000000001000) = 4
}

func ExampleLen32() {
	fmt.Printf("Len32(%032b) = %d\n", 8, bits.Len32(8))
	// Output:
	// Len32(00000000000000000000000000001000) = 4
}

func ExampleLen64() {
	fmt.Printf("Len64(%064b) = %d\n", 8, bits.Len64(8))
	// Output:
	// Len64(0000000000000000000000000000000000000000000000000000000000001000) = 4
}
