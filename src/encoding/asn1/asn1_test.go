// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"
	"time"
)

type boolTest struct {
	in  []byte
	ok  bool
	out bool
}

var boolTestData = []boolTest{
	{[]byte{0x00}, true, false},
	{[]byte{0xff}, true, true},
	{[]byte{0x00, 0x00}, false, false},
	{[]byte{0xff, 0xff}, false, false},
	{[]byte{0x01}, false, false},
}

func TestParseBool(t *testing.T) {
	for i, test := range boolTestData {
		ret, err := parseBool(test.in)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if test.ok && ret != test.out {
			t.Errorf("#%d: Bad result: %v (expected %v)", i, ret, test.out)
		}
	}
}

type int64Test struct {
	in  []byte
	ok  bool
	out int64
}

var int64TestData = []int64Test{
	{[]byte{0x00}, true, 0},
	{[]byte{0x7f}, true, 127},
	{[]byte{0x00, 0x80}, true, 128},
	{[]byte{0x01, 0x00}, true, 256},
	{[]byte{0x80}, true, -128},
	{[]byte{0xff, 0x7f}, true, -129},
	{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, true, -1},
	{[]byte{0xff}, true, -1},
	{[]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, true, -9223372036854775808},
	{[]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false, 0},
}

func TestParseInt64(t *testing.T) {
	for i, test := range int64TestData {
		ret, err := parseInt64(test.in)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if test.ok && ret != test.out {
			t.Errorf("#%d: Bad result: %v (expected %v)", i, ret, test.out)
		}
	}
}

type int32Test struct {
	in  []byte
	ok  bool
	out int32
}

var int32TestData = []int32Test{
	{[]byte{0x00}, true, 0},
	{[]byte{0x7f}, true, 127},
	{[]byte{0x00, 0x80}, true, 128},
	{[]byte{0x01, 0x00}, true, 256},
	{[]byte{0x80}, true, -128},
	{[]byte{0xff, 0x7f}, true, -129},
	{[]byte{0xff, 0xff, 0xff, 0xff}, true, -1},
	{[]byte{0xff}, true, -1},
	{[]byte{0x80, 0x00, 0x00, 0x00}, true, -2147483648},
	{[]byte{0x80, 0x00, 0x00, 0x00, 0x00}, false, 0},
}

func TestParseInt32(t *testing.T) {
	for i, test := range int32TestData {
		ret, err := parseInt32(test.in)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if test.ok && int32(ret) != test.out {
			t.Errorf("#%d: Bad result: %v (expected %v)", i, ret, test.out)
		}
	}
}

var bigIntTests = []struct {
	in     []byte
	base10 string
}{
	{[]byte{0xff}, "-1"},
	{[]byte{0x00}, "0"},
	{[]byte{0x01}, "1"},
	{[]byte{0x00, 0xff}, "255"},
	{[]byte{0xff, 0x00}, "-256"},
	{[]byte{0x01, 0x00}, "256"},
}

func TestParseBigInt(t *testing.T) {
	for i, test := range bigIntTests {
		ret := parseBigInt(test.in)
		if ret.String() != test.base10 {
			t.Errorf("#%d: bad result from %x, got %s want %s", i, test.in, ret.String(), test.base10)
		}
		fw := newForkableWriter()
		marshalBigInt(fw, ret)
		result := fw.Bytes()
		if !bytes.Equal(result, test.in) {
			t.Errorf("#%d: got %x from marshaling %s, want %x", i, result, ret, test.in)
		}
	}
}

type bitStringTest struct {
	in        []byte
	ok        bool
	out       []byte
	bitLength int
}

var bitStringTestData = []bitStringTest{
	{[]byte{}, false, []byte{}, 0},
	{[]byte{0x00}, true, []byte{}, 0},
	{[]byte{0x07, 0x00}, true, []byte{0x00}, 1},
	{[]byte{0x07, 0x01}, false, []byte{}, 0},
	{[]byte{0x07, 0x40}, false, []byte{}, 0},
	{[]byte{0x08, 0x00}, false, []byte{}, 0},
}

func TestBitString(t *testing.T) {
	for i, test := range bitStringTestData {
		ret, err := parseBitString(test.in)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if err == nil {
			if test.bitLength != ret.BitLength || !bytes.Equal(ret.Bytes, test.out) {
				t.Errorf("#%d: Bad result: %v (expected %v %v)", i, ret, test.out, test.bitLength)
			}
		}
	}
}

func TestBitStringAt(t *testing.T) {
	bs := BitString{[]byte{0x82, 0x40}, 16}
	if bs.At(0) != 1 {
		t.Error("#1: Failed")
	}
	if bs.At(1) != 0 {
		t.Error("#2: Failed")
	}
	if bs.At(6) != 1 {
		t.Error("#3: Failed")
	}
	if bs.At(9) != 1 {
		t.Error("#4: Failed")
	}
	if bs.At(-1) != 0 {
		t.Error("#5: Failed")
	}
	if bs.At(17) != 0 {
		t.Error("#6: Failed")
	}
}

type bitStringRightAlignTest struct {
	in    []byte
	inlen int
	out   []byte
}

var bitStringRightAlignTests = []bitStringRightAlignTest{
	{[]byte{0x80}, 1, []byte{0x01}},
	{[]byte{0x80, 0x80}, 9, []byte{0x01, 0x01}},
	{[]byte{}, 0, []byte{}},
	{[]byte{0xce}, 8, []byte{0xce}},
	{[]byte{0xce, 0x47}, 16, []byte{0xce, 0x47}},
	{[]byte{0x34, 0x50}, 12, []byte{0x03, 0x45}},
}

func TestBitStringRightAlign(t *testing.T) {
	for i, test := range bitStringRightAlignTests {
		bs := BitString{test.in, test.inlen}
		out := bs.RightAlign()
		if !bytes.Equal(out, test.out) {
			t.Errorf("#%d got: %x want: %x", i, out, test.out)
		}
	}
}

type objectIdentifierTest struct {
	in  []byte
	ok  bool
	out []int
}

var objectIdentifierTestData = []objectIdentifierTest{
	{[]byte{}, false, []int{}},
	{[]byte{85}, true, []int{2, 5}},
	{[]byte{85, 0x02}, true, []int{2, 5, 2}},
	{[]byte{85, 0x02, 0xc0, 0x00}, true, []int{2, 5, 2, 0x2000}},
	{[]byte{0x81, 0x34, 0x03}, true, []int{2, 100, 3}},
	{[]byte{85, 0x02, 0xc0, 0x80, 0x80, 0x80, 0x80}, false, []int{}},
}

func TestObjectIdentifier(t *testing.T) {
	for i, test := range objectIdentifierTestData {
		ret, err := parseObjectIdentifier(test.in)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if err == nil {
			if !reflect.DeepEqual(test.out, ret) {
				t.Errorf("#%d: Bad result: %v (expected %v)", i, ret, test.out)
			}
		}
	}

	if s := ObjectIdentifier([]int{1, 2, 3, 4}).String(); s != "1.2.3.4" {
		t.Errorf("bad ObjectIdentifier.String(). Got %s, want 1.2.3.4", s)
	}
}

type timeTest struct {
	in  string
	ok  bool
	out time.Time
}

var utcTestData = []timeTest{
	{"910506164540-0700", true, time.Date(1991, 05, 06, 16, 45, 40, 0, time.FixedZone("", -7*60*60))},
	{"910506164540+0730", true, time.Date(1991, 05, 06, 16, 45, 40, 0, time.FixedZone("", 7*60*60+30*60))},
	{"910506234540Z", true, time.Date(1991, 05, 06, 23, 45, 40, 0, time.UTC)},
	{"9105062345Z", true, time.Date(1991, 05, 06, 23, 45, 0, 0, time.UTC)},
	{"5105062345Z", true, time.Date(1951, 05, 06, 23, 45, 0, 0, time.UTC)},
	{"a10506234540Z", false, time.Time{}},
	{"91a506234540Z", false, time.Time{}},
	{"9105a6234540Z", false, time.Time{}},
	{"910506a34540Z", false, time.Time{}},
	{"910506334a40Z", false, time.Time{}},
	{"91050633444aZ", false, time.Time{}},
	{"910506334461Z", false, time.Time{}},
	{"910506334400Za", false, time.Time{}},
	/* These are invalid times. However, the time package normalises times
	 * and they were accepted in some versions. See #11134. */
	{"000100000000Z", false, time.Time{}},
	{"101302030405Z", false, time.Time{}},
	{"100002030405Z", false, time.Time{}},
	{"100100030405Z", false, time.Time{}},
	{"100132030405Z", false, time.Time{}},
	{"100231030405Z", false, time.Time{}},
	{"100102240405Z", false, time.Time{}},
	{"100102036005Z", false, time.Time{}},
	{"100102030460Z", false, time.Time{}},
	{"-100102030410Z", false, time.Time{}},
	{"10-0102030410Z", false, time.Time{}},
	{"10-0002030410Z", false, time.Time{}},
	{"1001-02030410Z", false, time.Time{}},
	{"100102-030410Z", false, time.Time{}},
	{"10010203-0410Z", false, time.Time{}},
	{"1001020304-10Z", false, time.Time{}},
}

func TestUTCTime(t *testing.T) {
	for i, test := range utcTestData {
		ret, err := parseUTCTime([]byte(test.in))
		if err != nil {
			if test.ok {
				t.Errorf("#%d: parseUTCTime(%q) = error %v", i, test.in, err)
			}
			continue
		}
		if !test.ok {
			t.Errorf("#%d: parseUTCTime(%q) succeeded, should have failed", i, test.in)
			continue
		}
		const format = "Jan _2 15:04:05 -0700 2006" // ignore zone name, just offset
		have := ret.Format(format)
		want := test.out.Format(format)
		if have != want {
			t.Errorf("#%d: parseUTCTime(%q) = %s, want %s", i, test.in, have, want)
		}
	}
}

var generalizedTimeTestData = []timeTest{
	{"20100102030405Z", true, time.Date(2010, 01, 02, 03, 04, 05, 0, time.UTC)},
	{"20100102030405", false, time.Time{}},
	{"20100102030405+0607", true, time.Date(2010, 01, 02, 03, 04, 05, 0, time.FixedZone("", 6*60*60+7*60))},
	{"20100102030405-0607", true, time.Date(2010, 01, 02, 03, 04, 05, 0, time.FixedZone("", -6*60*60-7*60))},
	/* These are invalid times. However, the time package normalises times
	 * and they were accepted in some versions. See #11134. */
	{"00000100000000Z", false, time.Time{}},
	{"20101302030405Z", false, time.Time{}},
	{"20100002030405Z", false, time.Time{}},
	{"20100100030405Z", false, time.Time{}},
	{"20100132030405Z", false, time.Time{}},
	{"20100231030405Z", false, time.Time{}},
	{"20100102240405Z", false, time.Time{}},
	{"20100102036005Z", false, time.Time{}},
	{"20100102030460Z", false, time.Time{}},
	{"-20100102030410Z", false, time.Time{}},
	{"2010-0102030410Z", false, time.Time{}},
	{"2010-0002030410Z", false, time.Time{}},
	{"201001-02030410Z", false, time.Time{}},
	{"20100102-030410Z", false, time.Time{}},
	{"2010010203-0410Z", false, time.Time{}},
	{"201001020304-10Z", false, time.Time{}},
}

func TestGeneralizedTime(t *testing.T) {
	for i, test := range generalizedTimeTestData {
		ret, err := parseGeneralizedTime([]byte(test.in))
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did fail? %v, expected: %v)", i, err == nil, test.ok)
		}
		if err == nil {
			if !reflect.DeepEqual(test.out, ret) {
				t.Errorf("#%d: Bad result: %q → %v (expected %v)", i, test.in, ret, test.out)
			}
		}
	}
}

type tagAndLengthTest struct {
	in  []byte
	ok  bool
	out tagAndLength
}

var tagAndLengthData = []tagAndLengthTest{
	{[]byte{0x80, 0x01}, true, tagAndLength{2, 0, 1, false}},
	{[]byte{0xa0, 0x01}, true, tagAndLength{2, 0, 1, true}},
	{[]byte{0x02, 0x00}, true, tagAndLength{0, 2, 0, false}},
	{[]byte{0xfe, 0x00}, true, tagAndLength{3, 30, 0, true}},
	{[]byte{0x1f, 0x01, 0x00}, true, tagAndLength{0, 1, 0, false}},
	{[]byte{0x1f, 0x81, 0x00, 0x00}, true, tagAndLength{0, 128, 0, false}},
	{[]byte{0x1f, 0x81, 0x80, 0x01, 0x00}, true, tagAndLength{0, 0x4001, 0, false}},
	{[]byte{0x00, 0x81, 0x01}, true, tagAndLength{0, 0, 1, false}},
	{[]byte{0x00, 0x82, 0x01, 0x00}, true, tagAndLength{0, 0, 256, false}},
	{[]byte{0x00, 0x83, 0x01, 0x00}, false, tagAndLength{}},
	{[]byte{0x1f, 0x85}, false, tagAndLength{}},
	{[]byte{0x30, 0x80}, false, tagAndLength{}},
	// Superfluous zeros in the length should be an error.
	{[]byte{0xa0, 0x82, 0x00, 0x01}, false, tagAndLength{}},
	// Lengths up to the maximum size of an int should work.
	{[]byte{0xa0, 0x84, 0x7f, 0xff, 0xff, 0xff}, true, tagAndLength{2, 0, 0x7fffffff, true}},
	// Lengths that would overflow an int should be rejected.
	{[]byte{0xa0, 0x84, 0x80, 0x00, 0x00, 0x00}, false, tagAndLength{}},
}

func TestParseTagAndLength(t *testing.T) {
	for i, test := range tagAndLengthData {
		tagAndLength, _, err := parseTagAndLength(test.in, 0)
		if (err == nil) != test.ok {
			t.Errorf("#%d: Incorrect error result (did pass? %v, expected: %v)", i, err == nil, test.ok)
		}
		if err == nil && !reflect.DeepEqual(test.out, tagAndLength) {
			t.Errorf("#%d: Bad result: %v (expected %v)", i, tagAndLength, test.out)
		}
	}
}

type parseFieldParametersTest struct {
	in  string
	out fieldParameters
}

func newInt(n int) *int { return &n }

func newInt64(n int64) *int64 { return &n }

func newString(s string) *string { return &s }

func newBool(b bool) *bool { return &b }

var parseFieldParametersTestData []parseFieldParametersTest = []parseFieldParametersTest{
	{"", fieldParameters{}},
	{"ia5", fieldParameters{stringType: tagIA5String}},
	{"generalized", fieldParameters{timeType: tagGeneralizedTime}},
	{"utc", fieldParameters{timeType: tagUTCTime}},
	{"printable", fieldParameters{stringType: tagPrintableString}},
	{"optional", fieldParameters{optional: true}},
	{"explicit", fieldParameters{explicit: true, tag: new(int)}},
	{"application", fieldParameters{application: true, tag: new(int)}},
	{"optional,explicit", fieldParameters{optional: true, explicit: true, tag: new(int)}},
	{"default:42", fieldParameters{defaultValue: newInt64(42)}},
	{"tag:17", fieldParameters{tag: newInt(17)}},
	{"optional,explicit,default:42,tag:17", fieldParameters{optional: true, explicit: true, defaultValue: newInt64(42), tag: newInt(17)}},
	{"optional,explicit,default:42,tag:17,rubbish1", fieldParameters{true, true, false, newInt64(42), newInt(17), 0, 0, false, false}},
	{"set", fieldParameters{set: true}},
}

func TestParseFieldParameters(t *testing.T) {
	for i, test := range parseFieldParametersTestData {
		f := parseFieldParameters(test.in)
		if !reflect.DeepEqual(f, test.out) {
			t.Errorf("#%d: Bad result: %v (expected %v)", i, f, test.out)
		}
	}
}

type TestObjectIdentifierStruct struct {
	OID ObjectIdentifier
}

type TestContextSpecificTags struct {
	A int `asn1:"tag:1"`
}

type TestContextSpecificTags2 struct {
	A int `asn1:"explicit,tag:1"`
	B int
}

type TestContextSpecificTags3 struct {
	S string `asn1:"tag:1,utf8"`
}

type TestElementsAfterString struct {
	S    string
	A, B int
}

type TestBigInt struct {
	X *big.Int
}

type TestSet struct {
	Ints []int `asn1:"set"`
}

var unmarshalTestData = []struct {
	in  []byte
	out interface{}
}{
	{[]byte{0x02, 0x01, 0x42}, newInt(0x42)},
	{[]byte{0x30, 0x08, 0x06, 0x06, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d}, &TestObjectIdentifierStruct{[]int{1, 2, 840, 113549}}},
	{[]byte{0x03, 0x04, 0x06, 0x6e, 0x5d, 0xc0}, &BitString{[]byte{110, 93, 192}, 18}},
	{[]byte{0x30, 0x09, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}, &[]int{1, 2, 3}},
	{[]byte{0x02, 0x01, 0x10}, newInt(16)},
	{[]byte{0x13, 0x04, 't', 'e', 's', 't'}, newString("test")},
	{[]byte{0x16, 0x04, 't', 'e', 's', 't'}, newString("test")},
	{[]byte{0x16, 0x04, 't', 'e', 's', 't'}, &RawValue{0, 22, false, []byte("test"), []byte("\x16\x04test")}},
	{[]byte{0x04, 0x04, 1, 2, 3, 4}, &RawValue{0, 4, false, []byte{1, 2, 3, 4}, []byte{4, 4, 1, 2, 3, 4}}},
	{[]byte{0x30, 0x03, 0x81, 0x01, 0x01}, &TestContextSpecificTags{1}},
	{[]byte{0x30, 0x08, 0xa1, 0x03, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02}, &TestContextSpecificTags2{1, 2}},
	{[]byte{0x30, 0x03, 0x81, 0x01, '@'}, &TestContextSpecificTags3{"@"}},
	{[]byte{0x01, 0x01, 0x00}, newBool(false)},
	{[]byte{0x01, 0x01, 0xff}, newBool(true)},
	{[]byte{0x30, 0x0b, 0x13, 0x03, 0x66, 0x6f, 0x6f, 0x02, 0x01, 0x22, 0x02, 0x01, 0x33}, &TestElementsAfterString{"foo", 0x22, 0x33}},
	{[]byte{0x30, 0x05, 0x02, 0x03, 0x12, 0x34, 0x56}, &TestBigInt{big.NewInt(0x123456)}},
	{[]byte{0x30, 0x0b, 0x31, 0x09, 0x02, 0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x01, 0x03}, &TestSet{Ints: []int{1, 2, 3}}},
}

func TestUnmarshal(t *testing.T) {
	for i, test := range unmarshalTestData {
		pv := reflect.New(reflect.TypeOf(test.out).Elem())
		val := pv.Interface()
		_, err := Unmarshal(test.in, val)
		if err != nil {
			t.Errorf("Unmarshal failed at index %d %v", i, err)
		}
		if !reflect.DeepEqual(val, test.out) {
			t.Errorf("#%d:\nhave %#v\nwant %#v", i, val, test.out)
		}
	}
}

type Certificate struct {
	TBSCertificate     TBSCertificate
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     BitString
}

type TBSCertificate struct {
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       RawValue
	SignatureAlgorithm AlgorithmIdentifier
	Issuer             RDNSequence
	Validity           Validity
	Subject            RDNSequence
	PublicKey          PublicKeyInfo
}

type AlgorithmIdentifier struct {
	Algorithm ObjectIdentifier
}

type RDNSequence []RelativeDistinguishedNameSET

type RelativeDistinguishedNameSET []AttributeTypeAndValue

type AttributeTypeAndValue struct {
	Type  ObjectIdentifier
	Value interface{}
}

type Validity struct {
	NotBefore, NotAfter time.Time
}

type PublicKeyInfo struct {
	Algorithm AlgorithmIdentifier
	PublicKey BitString
}

func TestCertificate(t *testing.T) {
	// This is a minimal, self-signed certificate that should parse correctly.
	var cert Certificate
	if _, err := Unmarshal(derEncodedSelfSignedCertBytes, &cert); err != nil {
		t.Errorf("Unmarshal failed: %v", err)
	}
	if !reflect.DeepEqual(cert, derEncodedSelfSignedCert) {
		t.Errorf("Bad result:\ngot: %+v\nwant: %+v", cert, derEncodedSelfSignedCert)
	}
}

func TestCertificateWithNUL(t *testing.T) {
	// This is the paypal NUL-hack certificate. It should fail to parse because
	// NUL isn't a permitted character in a PrintableString.

	var cert Certificate
	if _, err := Unmarshal(derEncodedPaypalNULCertBytes, &cert); err == nil {
		t.Error("Unmarshal succeeded, should not have")
	}
}

type rawStructTest struct {
	Raw RawContent
	A   int
}

func TestRawStructs(t *testing.T) {
	var s rawStructTest
	input := []byte{0x30, 0x03, 0x02, 0x01, 0x50}

	rest, err := Unmarshal(input, &s)
	if len(rest) != 0 {
		t.Errorf("incomplete parse: %x", rest)
		return
	}
	if err != nil {
		t.Error(err)
		return
	}
	if s.A != 0x50 {
		t.Errorf("bad value for A: got %d want %d", s.A, 0x50)
	}
	if !bytes.Equal([]byte(s.Raw), input) {
		t.Errorf("bad value for Raw: got %x want %x", s.Raw, input)
	}
}

type oiEqualTest struct {
	first  ObjectIdentifier
	second ObjectIdentifier
	same   bool
}

var oiEqualTests = []oiEqualTest{
	{
		ObjectIdentifier{1, 2, 3},
		ObjectIdentifier{1, 2, 3},
		true,
	},
	{
		ObjectIdentifier{1},
		ObjectIdentifier{1, 2, 3},
		false,
	},
	{
		ObjectIdentifier{1, 2, 3},
		ObjectIdentifier{10, 11, 12},
		false,
	},
}

func TestObjectIdentifierEqual(t *testing.T) {
	for _, o := range oiEqualTests {
		if s := o.first.Equal(o.second); s != o.same {
			t.Errorf("ObjectIdentifier.Equal: got: %t want: %t", s, o.same)
		}
	}
}

var derEncodedSelfSignedCert = Certificate{
	TBSCertificate: TBSCertificate{
		Version:            0,
		SerialNumber:       RawValue{Class: 0, Tag: 2, IsCompound: false, Bytes: []uint8{0x0, 0x8c, 0xc3, 0x37, 0x92, 0x10, 0xec, 0x2c, 0x98}, FullBytes: []byte{2, 9, 0x0, 0x8c, 0xc3, 0x37, 0x92, 0x10, 0xec, 0x2c, 0x98}},
		SignatureAlgorithm: AlgorithmIdentifier{Algorithm: ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}},
		Issuer: RDNSequence{
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 6}, Value: "XX"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 8}, Value: "Some-State"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 7}, Value: "City"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 10}, Value: "Internet Widgits Pty Ltd"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 3}, Value: "false.example.com"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}, Value: "false@example.com"}},
		},
		Validity: Validity{
			NotBefore: time.Date(2009, 10, 8, 00, 25, 53, 0, time.UTC),
			NotAfter:  time.Date(2010, 10, 8, 00, 25, 53, 0, time.UTC),
		},
		Subject: RDNSequence{
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 6}, Value: "XX"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 8}, Value: "Some-State"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 7}, Value: "City"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 10}, Value: "Internet Widgits Pty Ltd"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{2, 5, 4, 3}, Value: "false.example.com"}},
			RelativeDistinguishedNameSET{AttributeTypeAndValue{Type: ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}, Value: "false@example.com"}},
		},
		PublicKey: PublicKeyInfo{
			Algorithm: AlgorithmIdentifier{Algorithm: ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}},
			PublicKey: BitString{
				Bytes: []uint8{
					0x30, 0x48, 0x2, 0x41, 0x0, 0xcd, 0xb7,
					0x63, 0x9c, 0x32, 0x78, 0xf0, 0x6, 0xaa, 0x27, 0x7f, 0x6e, 0xaf, 0x42,
					0x90, 0x2b, 0x59, 0x2d, 0x8c, 0xbc, 0xbe, 0x38, 0xa1, 0xc9, 0x2b, 0xa4,
					0x69, 0x5a, 0x33, 0x1b, 0x1d, 0xea, 0xde, 0xad, 0xd8, 0xe9, 0xa5, 0xc2,
					0x7e, 0x8c, 0x4c, 0x2f, 0xd0, 0xa8, 0x88, 0x96, 0x57, 0x72, 0x2a, 0x4f,
					0x2a, 0xf7, 0x58, 0x9c, 0xf2, 0xc7, 0x70, 0x45, 0xdc, 0x8f, 0xde, 0xec,
					0x35, 0x7d, 0x2, 0x3, 0x1, 0x0, 0x1,
				},
				BitLength: 592,
			},
		},
	},
	SignatureAlgorithm: AlgorithmIdentifier{Algorithm: ObjectIdentifier{1, 2, 840, 113549, 1, 1, 5}},
	SignatureValue: BitString{
		Bytes: []uint8{
			0xa6, 0x7b, 0x6, 0xec, 0x5e, 0xce,
			0x92, 0x77, 0x2c, 0xa4, 0x13, 0xcb, 0xa3, 0xca, 0x12, 0x56, 0x8f, 0xdc, 0x6c,
			0x7b, 0x45, 0x11, 0xcd, 0x40, 0xa7, 0xf6, 0x59, 0x98, 0x4, 0x2, 0xdf, 0x2b,
			0x99, 0x8b, 0xb9, 0xa4, 0xa8, 0xcb, 0xeb, 0x34, 0xc0, 0xf0, 0xa7, 0x8c, 0xf8,
			0xd9, 0x1e, 0xde, 0x14, 0xa5, 0xed, 0x76, 0xbf, 0x11, 0x6f, 0xe3, 0x60, 0xaa,
			0xfa, 0x88, 0x21, 0x49, 0x4, 0x35,
		},
		BitLength: 512,
	},
}

var derEncodedSelfSignedCertBytes = []byte{
	0x30, 0x82, 0x02, 0x18, 0x30,
	0x82, 0x01, 0xc2, 0x02, 0x09, 0x00, 0x8c, 0xc3, 0x37, 0x92, 0x10, 0xec, 0x2c,
	0x98, 0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01,
	0x05, 0x05, 0x00, 0x30, 0x81, 0x92, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
	0x04, 0x06, 0x13, 0x02, 0x58, 0x58, 0x31, 0x13, 0x30, 0x11, 0x06, 0x03, 0x55,
	0x04, 0x08, 0x13, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x31, 0x0d, 0x30, 0x0b, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04, 0x0a, 0x13,
	0x18, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64,
	0x67, 0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x31,
	0x1a, 0x30, 0x18, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x11, 0x66, 0x61, 0x6c,
	0x73, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x31, 0x20, 0x30, 0x1e, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d,
	0x01, 0x09, 0x01, 0x16, 0x11, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x40, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x30, 0x1e, 0x17, 0x0d,
	0x30, 0x39, 0x31, 0x30, 0x30, 0x38, 0x30, 0x30, 0x32, 0x35, 0x35, 0x33, 0x5a,
	0x17, 0x0d, 0x31, 0x30, 0x31, 0x30, 0x30, 0x38, 0x30, 0x30, 0x32, 0x35, 0x35,
	0x33, 0x5a, 0x30, 0x81, 0x92, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04,
	0x06, 0x13, 0x02, 0x58, 0x58, 0x31, 0x13, 0x30, 0x11, 0x06, 0x03, 0x55, 0x04,
	0x08, 0x13, 0x0a, 0x53, 0x6f, 0x6d, 0x65, 0x2d, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x31, 0x0d, 0x30, 0x0b, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x04, 0x43, 0x69,
	0x74, 0x79, 0x31, 0x21, 0x30, 0x1f, 0x06, 0x03, 0x55, 0x04, 0x0a, 0x13, 0x18,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x20, 0x57, 0x69, 0x64, 0x67,
	0x69, 0x74, 0x73, 0x20, 0x50, 0x74, 0x79, 0x20, 0x4c, 0x74, 0x64, 0x31, 0x1a,
	0x30, 0x18, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x11, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x31, 0x20, 0x30, 0x1e, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01,
	0x09, 0x01, 0x16, 0x11, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x40, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x30, 0x5c, 0x30, 0x0d, 0x06,
	0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x01, 0x05, 0x00, 0x03,
	0x4b, 0x00, 0x30, 0x48, 0x02, 0x41, 0x00, 0xcd, 0xb7, 0x63, 0x9c, 0x32, 0x78,
	0xf0, 0x06, 0xaa, 0x27, 0x7f, 0x6e, 0xaf, 0x42, 0x90, 0x2b, 0x59, 0x2d, 0x8c,
	0xbc, 0xbe, 0x38, 0xa1, 0xc9, 0x2b, 0xa4, 0x69, 0x5a, 0x33, 0x1b, 0x1d, 0xea,
	0xde, 0xad, 0xd8, 0xe9, 0xa5, 0xc2, 0x7e, 0x8c, 0x4c, 0x2f, 0xd0, 0xa8, 0x88,
	0x96, 0x57, 0x72, 0x2a, 0x4f, 0x2a, 0xf7, 0x58, 0x9c, 0xf2, 0xc7, 0x70, 0x45,
	0xdc, 0x8f, 0xde, 0xec, 0x35, 0x7d, 0x02, 0x03, 0x01, 0x00, 0x01, 0x30, 0x0d,
	0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x05, 0x05, 0x00,
	0x03, 0x41, 0x00, 0xa6, 0x7b, 0x06, 0xec, 0x5e, 0xce, 0x92, 0x77, 0x2c, 0xa4,
	0x13, 0xcb, 0xa3, 0xca, 0x12, 0x56, 0x8f, 0xdc, 0x6c, 0x7b, 0x45, 0x11, 0xcd,
	0x40, 0xa7, 0xf6, 0x59, 0x98, 0x04, 0x02, 0xdf, 0x2b, 0x99, 0x8b, 0xb9, 0xa4,
	0xa8, 0xcb, 0xeb, 0x34, 0xc0, 0xf0, 0xa7, 0x8c, 0xf8, 0xd9, 0x1e, 0xde, 0x14,
	0xa5, 0xed, 0x76, 0xbf, 0x11, 0x6f, 0xe3, 0x60, 0xaa, 0xfa, 0x88, 0x21, 0x49,
	0x04, 0x35,
}

var derEncodedPaypalNULCertBytes = []byte{
	0x30, 0x82, 0x06, 0x44, 0x30,
	0x82, 0x05, 0xad, 0xa0, 0x03, 0x02, 0x01, 0x02, 0x02, 0x03, 0x00, 0xf0, 0x9b,
	0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x05,
	0x05, 0x00, 0x30, 0x82, 0x01, 0x12, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55,
	0x04, 0x06, 0x13, 0x02, 0x45, 0x53, 0x31, 0x12, 0x30, 0x10, 0x06, 0x03, 0x55,
	0x04, 0x08, 0x13, 0x09, 0x42, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x6f, 0x6e, 0x61,
	0x31, 0x12, 0x30, 0x10, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x09, 0x42, 0x61,
	0x72, 0x63, 0x65, 0x6c, 0x6f, 0x6e, 0x61, 0x31, 0x29, 0x30, 0x27, 0x06, 0x03,
	0x55, 0x04, 0x0a, 0x13, 0x20, 0x49, 0x50, 0x53, 0x20, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x20, 0x73, 0x2e, 0x6c, 0x2e, 0x31, 0x2e,
	0x30, 0x2c, 0x06, 0x03, 0x55, 0x04, 0x0a, 0x14, 0x25, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x40, 0x69, 0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d,
	0x20, 0x43, 0x2e, 0x49, 0x2e, 0x46, 0x2e, 0x20, 0x20, 0x42, 0x2d, 0x42, 0x36,
	0x32, 0x32, 0x31, 0x30, 0x36, 0x39, 0x35, 0x31, 0x2e, 0x30, 0x2c, 0x06, 0x03,
	0x55, 0x04, 0x0b, 0x13, 0x25, 0x69, 0x70, 0x73, 0x43, 0x41, 0x20, 0x43, 0x4c,
	0x41, 0x53, 0x45, 0x41, 0x31, 0x20, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x31, 0x2e, 0x30, 0x2c, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13,
	0x25, 0x69, 0x70, 0x73, 0x43, 0x41, 0x20, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41,
	0x31, 0x20, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x31,
	0x20, 0x30, 0x1e, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x09,
	0x01, 0x16, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x40, 0x69, 0x70,
	0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x30, 0x1e, 0x17, 0x0d, 0x30, 0x39,
	0x30, 0x32, 0x32, 0x34, 0x32, 0x33, 0x30, 0x34, 0x31, 0x37, 0x5a, 0x17, 0x0d,
	0x31, 0x31, 0x30, 0x32, 0x32, 0x34, 0x32, 0x33, 0x30, 0x34, 0x31, 0x37, 0x5a,
	0x30, 0x81, 0x94, 0x31, 0x0b, 0x30, 0x09, 0x06, 0x03, 0x55, 0x04, 0x06, 0x13,
	0x02, 0x55, 0x53, 0x31, 0x13, 0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x08, 0x13,
	0x0a, 0x43, 0x61, 0x6c, 0x69, 0x66, 0x6f, 0x72, 0x6e, 0x69, 0x61, 0x31, 0x16,
	0x30, 0x14, 0x06, 0x03, 0x55, 0x04, 0x07, 0x13, 0x0d, 0x53, 0x61, 0x6e, 0x20,
	0x46, 0x72, 0x61, 0x6e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x31, 0x11, 0x30, 0x0f,
	0x06, 0x03, 0x55, 0x04, 0x0a, 0x13, 0x08, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x31, 0x14, 0x30, 0x12, 0x06, 0x03, 0x55, 0x04, 0x0b, 0x13, 0x0b,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x65, 0x20, 0x55, 0x6e, 0x69, 0x74, 0x31, 0x2f,
	0x30, 0x2d, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x26, 0x77, 0x77, 0x77, 0x2e,
	0x70, 0x61, 0x79, 0x70, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x73, 0x73,
	0x6c, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x63, 0x30, 0x81, 0x9f, 0x30, 0x0d,
	0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x01, 0x05, 0x00,
	0x03, 0x81, 0x8d, 0x00, 0x30, 0x81, 0x89, 0x02, 0x81, 0x81, 0x00, 0xd2, 0x69,
	0xfa, 0x6f, 0x3a, 0x00, 0xb4, 0x21, 0x1b, 0xc8, 0xb1, 0x02, 0xd7, 0x3f, 0x19,
	0xb2, 0xc4, 0x6d, 0xb4, 0x54, 0xf8, 0x8b, 0x8a, 0xcc, 0xdb, 0x72, 0xc2, 0x9e,
	0x3c, 0x60, 0xb9, 0xc6, 0x91, 0x3d, 0x82, 0xb7, 0x7d, 0x99, 0xff, 0xd1, 0x29,
	0x84, 0xc1, 0x73, 0x53, 0x9c, 0x82, 0xdd, 0xfc, 0x24, 0x8c, 0x77, 0xd5, 0x41,
	0xf3, 0xe8, 0x1e, 0x42, 0xa1, 0xad, 0x2d, 0x9e, 0xff, 0x5b, 0x10, 0x26, 0xce,
	0x9d, 0x57, 0x17, 0x73, 0x16, 0x23, 0x38, 0xc8, 0xd6, 0xf1, 0xba, 0xa3, 0x96,
	0x5b, 0x16, 0x67, 0x4a, 0x4f, 0x73, 0x97, 0x3a, 0x4d, 0x14, 0xa4, 0xf4, 0xe2,
	0x3f, 0x8b, 0x05, 0x83, 0x42, 0xd1, 0xd0, 0xdc, 0x2f, 0x7a, 0xe5, 0xb6, 0x10,
	0xb2, 0x11, 0xc0, 0xdc, 0x21, 0x2a, 0x90, 0xff, 0xae, 0x97, 0x71, 0x5a, 0x49,
	0x81, 0xac, 0x40, 0xf3, 0x3b, 0xb8, 0x59, 0xb2, 0x4f, 0x02, 0x03, 0x01, 0x00,
	0x01, 0xa3, 0x82, 0x03, 0x21, 0x30, 0x82, 0x03, 0x1d, 0x30, 0x09, 0x06, 0x03,
	0x55, 0x1d, 0x13, 0x04, 0x02, 0x30, 0x00, 0x30, 0x11, 0x06, 0x09, 0x60, 0x86,
	0x48, 0x01, 0x86, 0xf8, 0x42, 0x01, 0x01, 0x04, 0x04, 0x03, 0x02, 0x06, 0x40,
	0x30, 0x0b, 0x06, 0x03, 0x55, 0x1d, 0x0f, 0x04, 0x04, 0x03, 0x02, 0x03, 0xf8,
	0x30, 0x13, 0x06, 0x03, 0x55, 0x1d, 0x25, 0x04, 0x0c, 0x30, 0x0a, 0x06, 0x08,
	0x2b, 0x06, 0x01, 0x05, 0x05, 0x07, 0x03, 0x01, 0x30, 0x1d, 0x06, 0x03, 0x55,
	0x1d, 0x0e, 0x04, 0x16, 0x04, 0x14, 0x61, 0x8f, 0x61, 0x34, 0x43, 0x55, 0x14,
	0x7f, 0x27, 0x09, 0xce, 0x4c, 0x8b, 0xea, 0x9b, 0x7b, 0x19, 0x25, 0xbc, 0x6e,
	0x30, 0x1f, 0x06, 0x03, 0x55, 0x1d, 0x23, 0x04, 0x18, 0x30, 0x16, 0x80, 0x14,
	0x0e, 0x07, 0x60, 0xd4, 0x39, 0xc9, 0x1b, 0x5b, 0x5d, 0x90, 0x7b, 0x23, 0xc8,
	0xd2, 0x34, 0x9d, 0x4a, 0x9a, 0x46, 0x39, 0x30, 0x09, 0x06, 0x03, 0x55, 0x1d,
	0x11, 0x04, 0x02, 0x30, 0x00, 0x30, 0x1c, 0x06, 0x03, 0x55, 0x1d, 0x12, 0x04,
	0x15, 0x30, 0x13, 0x81, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x40,
	0x69, 0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x30, 0x72, 0x06, 0x09,
	0x60, 0x86, 0x48, 0x01, 0x86, 0xf8, 0x42, 0x01, 0x0d, 0x04, 0x65, 0x16, 0x63,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x4e,
	0x4f, 0x54, 0x20, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x44, 0x2e,
	0x20, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41, 0x31, 0x20, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x20, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x20, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x70,
	0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x2f, 0x06, 0x09, 0x60,
	0x86, 0x48, 0x01, 0x86, 0xf8, 0x42, 0x01, 0x02, 0x04, 0x22, 0x16, 0x20, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x70,
	0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61,
	0x32, 0x30, 0x30, 0x32, 0x2f, 0x30, 0x43, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01,
	0x86, 0xf8, 0x42, 0x01, 0x04, 0x04, 0x36, 0x16, 0x34, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x70, 0x73, 0x63, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61, 0x32, 0x30, 0x30,
	0x32, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61, 0x32, 0x30, 0x30, 0x32, 0x43, 0x4c,
	0x41, 0x53, 0x45, 0x41, 0x31, 0x2e, 0x63, 0x72, 0x6c, 0x30, 0x46, 0x06, 0x09,
	0x60, 0x86, 0x48, 0x01, 0x86, 0xf8, 0x42, 0x01, 0x03, 0x04, 0x39, 0x16, 0x37,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69,
	0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63,
	0x61, 0x32, 0x30, 0x30, 0x32, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41, 0x31, 0x2e, 0x68, 0x74,
	0x6d, 0x6c, 0x3f, 0x30, 0x43, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x86, 0xf8,
	0x42, 0x01, 0x07, 0x04, 0x36, 0x16, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x70, 0x73, 0x63, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61, 0x32, 0x30, 0x30, 0x32, 0x2f,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41,
	0x31, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x3f, 0x30, 0x41, 0x06, 0x09, 0x60, 0x86,
	0x48, 0x01, 0x86, 0xf8, 0x42, 0x01, 0x08, 0x04, 0x34, 0x16, 0x32, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x70, 0x73,
	0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61, 0x32,
	0x30, 0x30, 0x32, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x4c, 0x41,
	0x53, 0x45, 0x41, 0x31, 0x2e, 0x68, 0x74, 0x6d, 0x6c, 0x30, 0x81, 0x83, 0x06,
	0x03, 0x55, 0x1d, 0x1f, 0x04, 0x7c, 0x30, 0x7a, 0x30, 0x39, 0xa0, 0x37, 0xa0,
	0x35, 0x86, 0x33, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x69, 0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70,
	0x73, 0x63, 0x61, 0x32, 0x30, 0x30, 0x32, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61,
	0x32, 0x30, 0x30, 0x32, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41, 0x31, 0x2e, 0x63,
	0x72, 0x6c, 0x30, 0x3d, 0xa0, 0x3b, 0xa0, 0x39, 0x86, 0x37, 0x68, 0x74, 0x74,
	0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x69,
	0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x70, 0x73, 0x63,
	0x61, 0x32, 0x30, 0x30, 0x32, 0x2f, 0x69, 0x70, 0x73, 0x63, 0x61, 0x32, 0x30,
	0x30, 0x32, 0x43, 0x4c, 0x41, 0x53, 0x45, 0x41, 0x31, 0x2e, 0x63, 0x72, 0x6c,
	0x30, 0x32, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x05, 0x05, 0x07, 0x01, 0x01, 0x04,
	0x26, 0x30, 0x24, 0x30, 0x22, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x05, 0x05, 0x07,
	0x30, 0x01, 0x86, 0x16, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6f, 0x63,
	0x73, 0x70, 0x2e, 0x69, 0x70, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x05,
	0x05, 0x00, 0x03, 0x81, 0x81, 0x00, 0x68, 0xee, 0x79, 0x97, 0x97, 0xdd, 0x3b,
	0xef, 0x16, 0x6a, 0x06, 0xf2, 0x14, 0x9a, 0x6e, 0xcd, 0x9e, 0x12, 0xf7, 0xaa,
	0x83, 0x10, 0xbd, 0xd1, 0x7c, 0x98, 0xfa, 0xc7, 0xae, 0xd4, 0x0e, 0x2c, 0x9e,
	0x38, 0x05, 0x9d, 0x52, 0x60, 0xa9, 0x99, 0x0a, 0x81, 0xb4, 0x98, 0x90, 0x1d,
	0xae, 0xbb, 0x4a, 0xd7, 0xb9, 0xdc, 0x88, 0x9e, 0x37, 0x78, 0x41, 0x5b, 0xf7,
	0x82, 0xa5, 0xf2, 0xba, 0x41, 0x25, 0x5a, 0x90, 0x1a, 0x1e, 0x45, 0x38, 0xa1,
	0x52, 0x58, 0x75, 0x94, 0x26, 0x44, 0xfb, 0x20, 0x07, 0xba, 0x44, 0xcc, 0xe5,
	0x4a, 0x2d, 0x72, 0x3f, 0x98, 0x47, 0xf6, 0x26, 0xdc, 0x05, 0x46, 0x05, 0x07,
	0x63, 0x21, 0xab, 0x46, 0x9b, 0x9c, 0x78, 0xd5, 0x54, 0x5b, 0x3d, 0x0c, 0x1e,
	0xc8, 0x64, 0x8c, 0xb5, 0x50, 0x23, 0x82, 0x6f, 0xdb, 0xb8, 0x22, 0x1c, 0x43,
	0x96, 0x07, 0xa8, 0xbb,
}

var stringSliceTestData = [][]string{
	{"foo", "bar"},
	{"foo", "\\bar"},
	{"foo", "\"bar\""},
	{"foo", "åäö"},
}

func TestStringSlice(t *testing.T) {
	for _, test := range stringSliceTestData {
		bs, err := Marshal(test)
		if err != nil {
			t.Error(err)
		}

		var res []string
		_, err = Unmarshal(bs, &res)
		if err != nil {
			t.Error(err)
		}

		if fmt.Sprintf("%v", res) != fmt.Sprintf("%v", test) {
			t.Errorf("incorrect marshal/unmarshal; %v != %v", res, test)
		}
	}
}

type explicitTaggedTimeTest struct {
	Time time.Time `asn1:"explicit,tag:0"`
}

var explicitTaggedTimeTestData = []struct {
	in  []byte
	out explicitTaggedTimeTest
}{
	{[]byte{0x30, 0x11, 0xa0, 0xf, 0x17, 0xd, '9', '1', '0', '5', '0', '6', '1', '6', '4', '5', '4', '0', 'Z'},
		explicitTaggedTimeTest{time.Date(1991, 05, 06, 16, 45, 40, 0, time.UTC)}},
	{[]byte{0x30, 0x17, 0xa0, 0xf, 0x18, 0x13, '2', '0', '1', '0', '0', '1', '0', '2', '0', '3', '0', '4', '0', '5', '+', '0', '6', '0', '7'},
		explicitTaggedTimeTest{time.Date(2010, 01, 02, 03, 04, 05, 0, time.FixedZone("", 6*60*60+7*60))}},
}

func TestExplicitTaggedTime(t *testing.T) {
	// Test that a time.Time will match either tagUTCTime or
	// tagGeneralizedTime.
	for i, test := range explicitTaggedTimeTestData {
		var got explicitTaggedTimeTest
		_, err := Unmarshal(test.in, &got)
		if err != nil {
			t.Errorf("Unmarshal failed at index %d %v", i, err)
		}
		if !got.Time.Equal(test.out.Time) {
			t.Errorf("#%d: got %v, want %v", i, got.Time, test.out.Time)
		}
	}
}

type implicitTaggedTimeTest struct {
	Time time.Time `asn1:"tag:24"`
}

func TestImplicitTaggedTime(t *testing.T) {
	// An implicitly tagged time value, that happens to have an implicit
	// tag equal to a GENERALIZEDTIME, should still be parsed as a UTCTime.
	// (There's no "timeType" in fieldParameters to determine what type of
	// time should be expected when implicitly tagged.)
	der := []byte{0x30, 0x0f, 0x80 | 24, 0xd, '9', '1', '0', '5', '0', '6', '1', '6', '4', '5', '4', '0', 'Z'}
	var result implicitTaggedTimeTest
	if _, err := Unmarshal(der, &result); err != nil {
		t.Fatalf("Error while parsing: %s", err)
	}
	if expected := time.Date(1991, 05, 06, 16, 45, 40, 0, time.UTC); !result.Time.Equal(expected) {
		t.Errorf("Wrong result. Got %v, want %v", result.Time, expected)
	}
}

type truncatedExplicitTagTest struct {
	Test int `asn1:"explicit,tag:0"`
}

func TestTruncatedExplicitTag(t *testing.T) {
	// This crashed Unmarshal in the past. See #11154.
	der := []byte{
		0x30, // SEQUENCE
		0x02, // two bytes long
		0xa0, // context-specific, tag 0
		0x30, // 48 bytes long
	}

	var result truncatedExplicitTagTest
	if _, err := Unmarshal(der, &result); err == nil {
		t.Error("Unmarshal returned without error")
	}
}

type invalidUTF8Test struct {
	Str string `asn1:"utf8"`
}

func TestUnmarshalInvalidUTF8(t *testing.T) {
	data := []byte("0\x05\f\x03a\xc9c")
	var result invalidUTF8Test
	_, err := Unmarshal(data, &result)

	const expectedSubstring = "UTF"
	if err == nil {
		t.Fatal("Successfully unmarshaled invalid UTF-8 data")
	} else if !strings.Contains(err.Error(), expectedSubstring) {
		t.Fatalf("Expected error to mention %q but error was %q", expectedSubstring, err.Error())
	}
}

func TestMarshalNilValue(t *testing.T) {
	nilValueTestData := []interface{}{
		nil,
		struct{ v interface{} }{},
	}
	for i, test := range nilValueTestData {
		if _, err := Marshal(test); err == nil {
			t.Fatal("#%d: successfully marshaled nil value", i)
		}
	}
}
