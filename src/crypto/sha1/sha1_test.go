// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SHA-1 hash algorithm. See RFC 3174.

package sha1

import (
	"bytes"
	"crypto/internal/boring"
	"crypto/rand"
	"encoding"
	"fmt"
	"io"
	"testing"
)

type sha1Test struct {
	out       string
	in        string
	halfState string // marshaled hash state after first half of in written, used by TestGoldenMarshal
}

var golden = []sha1Test{
	{"76245dbf96f661bd221046197ab8b9f063f11bad", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n", "sha\x01\v\xa0)I\xdeq(8h\x9ev\xe5\x88[\xf8\x81\x17\xba4Daaaaaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x96"},
	{"da39a3ee5e6b4b0d3255bfef95601890afd80709", "", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"86f7e437faa5a7fce15d1ddcb9eaeaea377667b8", "a", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x00"},
	{"da23614e02469a0d7c7bd1bdab5c9c474b1904dc", "ab", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"a9993e364706816aba3e25717850c26c9cd0d89d", "abc", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x01"},
	{"81fe8bfe87576c3ecb22426f8e57847382917acf", "abcd", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"03de6c570bfe24bfc328ccd7ca46b76eadaf4334", "abcde", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x02"},
	{"1f8ac10f23c5b5bc1167bda84b833e5c057a77d2", "abcdef", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdeaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"2fb5e13419fc89246865e7a324f476ec624e8740", "abcdefg", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdefaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x03"},
	{"425af12a0743502b322e93a015bcf868e324d56a", "abcdefgh", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdefgaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1", "abcdefghi", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdefghaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x04"},
	{"d68c19a0a345b7eab78d5e11e991c026ec60db63", "abcdefghij", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0abcdefghiaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x05"},
	{"ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7", "Discard medicine more than two years old.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Discard medicine moraaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x14"},
	{"e5dea09392dd886ca63531aaa00571dc07554bb6", "He who has a shady past knows that nice guys finish last.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0He who has a shady past knowwo years old.aaa\naaaaaaaaaaaaaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"45988f7234467b94e3e9494434c96ee3609d8f8f", "I wouldn't marry him with a ten foot pole.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0I wouldn't marry him st knows that nice guys finish last.aaaaaaa\x00\x00\x00\x00\x00\x00\x00\x15"},
	{"55dee037eb7460d5a692d1ce11330b260e40c988", "Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Free! Free!/A trip/to Mars/ften foot pole.ys finish last.aaaaaaa\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"b7bc5fb91080c7de6b582ea281f8a396d7c0aee8", "The days of the digital watch are numbered.  -Tom Stoppard", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0The days of the digital watchr 900/empty jars/Burma Shaveaaaaaaa\x00\x00\x00\x00\x00\x00\x00\x1d"},
	{"c3aed9358f7c77f523afe86135f06b95b3999797", "Nepal premier won't resign.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Nepal premierhe digital watch are numbered.  -Tom Stoppardaaaaaa\x00\x00\x00\x00\x00\x00\x00\r"},
	{"6e29d302bf6e3a5e4305ff318d983197d6906bb9", "For every action there is an equal and opposite government program.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0For every action there is an equa numbered.  -Tom Stoppardaaaaaa\x00\x00\x00\x00\x00\x00\x00!"},
	{"597f6a540010f94c15d71806a99a2c8710e747bd", "His money is twice tainted: 'taint yours and 'taint mine.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0His money is twice tainted:  equal and opposite government progr\x00\x00\x00\x00\x00\x00\x00\x1c"},
	{"6859733b2590a8a091cecf50086febc5ceef1e80", "There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0There is no reason for any individual to hav 'taint mine.t progr\x00\x00\x00\x00\x00\x00\x00,"},
	{"514b2630ec089b8aee18795fc0cf1f4860cdacad", "It's a tiny change to the code and not completely disgusting. - Bob Manchek", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0It's a tiny change to the code and no to have a computer in thei\x00\x00\x00\x00\x00\x00\x00%"},
	{"c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69", "size:  a.out:  bad magic", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0size:  a.outchange to the code and not completely disgusting. - \x00\x00\x00\x00\x00\x00\x00\f"},
	{"74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a", "The major problem is with sendmail.  -Mark Horton", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0The major problem is wite code and not completely disgusting. - \x00\x00\x00\x00\x00\x00\x00\x18"},
	{"0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66", "Give me a rock, paper and scissors and I will move the world.  CCFestoon", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Give me a rock, paper and scissors a -Mark Horton disgusting. - \x00\x00\x00\x00\x00\x00\x00$"},
	{"3ae7937dd790315beb0f48330e8642237c61550a", "If the enemy is within range, then so are you.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0If the enemy is within nd scissors and I will move the world.  C\x00\x00\x00\x00\x00\x00\x00\x17"},
	{"410a2b296df92b9a47412b13281df8f830a9f44b", "It's well we cannot hear the screams/That we create in others' dreams.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0It's well we cannot hear the screamso are you.move the world.  C\x00\x00\x00\x00\x00\x00\x00#"},
	{"841e7c85ca1adcddbdd0187f1289acb5c642f7f5", "You remind me of a TV show, but that's all right: I watch it anyway.", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0You remind me of a TV show, but thms/That we create in others' d\x00\x00\x00\x00\x00\x00\x00\""},
	{"163173b825d03b952601376b25212df66763e1db", "C is as portable as Stonehedge!!", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0C is as portable a TV show, but that's all right: I watch it any\x00\x00\x00\x00\x00\x00\x00\x10"},
	{"32b0377f2687eb88e22106f133c586ab314d5279", "Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0Even if I could be Shakespeare, I think I shight: I watch it any\x00\x00\x00\x00\x00\x00\x00,"},
	{"0885aaf99b569542fd165fa44e322718f4a984e0", "The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule", "sha\x01x}\xf4\r\xeb\xf2\x10\x87\xe8[\xb2JA$D\xb7غx8eme Faraday. - A. Huxleyspeare, I think I should still choose to\x00\x00\x00\x00\x00\x00\x00B"},
	{"6627d6904d71420b0bf3886ab629623538689f45", "How can you write a big system without C++?  -Paul Glick", "sha\x01gE#\x01\xefͫ\x89\x98\xba\xdc\xfe\x102Tv\xc3\xd2\xe1\xf0How can you write a big systo its mole fraction.  Lewis-Randall \x00\x00\x00\x00\x00\x00\x00\x1c"},
}

func TestGolden(t *testing.T) {
	for i := 0; i < len(golden); i++ {
		g := golden[i]
		s := fmt.Sprintf("%x", Sum([]byte(g.in)))
		if s != g.out {
			t.Fatalf("Sum function: sha1(%s) = %s want %s", g.in, s, g.out)
		}
		c := New()
		for j := 0; j < 4; j++ {
			var sum []byte
			switch j {
			case 0, 1:
				io.WriteString(c, g.in)
				sum = c.Sum(nil)
			case 2:
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.Sum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
				sum = c.Sum(nil)
			case 3:
				if boring.Enabled {
					continue
				}
				io.WriteString(c, g.in[0:len(g.in)/2])
				c.(*digest).ConstantTimeSum(nil)
				io.WriteString(c, g.in[len(g.in)/2:])
				sum = c.(*digest).ConstantTimeSum(nil)
			}
			s := fmt.Sprintf("%x", sum)
			if s != g.out {
				t.Fatalf("sha1[%d](%s) = %s want %s", j, g.in, s, g.out)
			}
			c.Reset()
		}
	}
}

func TestGoldenMarshal(t *testing.T) {
	h := New()
	h2 := New()
	for _, g := range golden {
		h.Reset()
		h2.Reset()

		io.WriteString(h, g.in[:len(g.in)/2])

		state, err := h.(encoding.BinaryMarshaler).MarshalBinary()
		if err != nil {
			t.Errorf("could not marshal: %v", err)
			continue
		}

		if string(state) != g.halfState {
			t.Errorf("sha1(%q) state = %q, want %q", g.in, state, g.halfState)
			continue
		}

		if err := h2.(encoding.BinaryUnmarshaler).UnmarshalBinary(state); err != nil {
			t.Errorf("could not unmarshal: %v", err)
			continue
		}

		io.WriteString(h, g.in[len(g.in)/2:])
		io.WriteString(h2, g.in[len(g.in)/2:])

		if actual, actual2 := h.Sum(nil), h2.Sum(nil); !bytes.Equal(actual, actual2) {
			t.Errorf("sha1(%q) = 0x%x != marshaled 0x%x", g.in, actual, actual2)
		}
	}
}

func TestSize(t *testing.T) {
	c := New()
	if got := c.Size(); got != Size {
		t.Errorf("Size = %d; want %d", got, Size)
	}
}

func TestBlockSize(t *testing.T) {
	c := New()
	if got := c.BlockSize(); got != BlockSize {
		t.Errorf("BlockSize = %d; want %d", got, BlockSize)
	}
}

// Tests that blockGeneric (pure Go) and block (in assembly for some architectures) match.
func TestBlockGeneric(t *testing.T) {
	if boring.Enabled {
		t.Skip("BoringCrypto doesn't expose digest")
	}
	for i := 1; i < 30; i++ { // arbitrary factor
		gen, asm := New().(*digest), New().(*digest)
		buf := make([]byte, BlockSize*i)
		rand.Read(buf)
		blockGeneric(gen, buf)
		block(asm, buf)
		if *gen != *asm {
			t.Errorf("For %#v block and blockGeneric resulted in different states", buf)
		}
	}
}

var bench = New()
var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

func BenchmarkHash320Bytes(b *testing.B) {
	benchmarkSize(b, 320)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}
