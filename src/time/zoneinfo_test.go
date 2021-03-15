// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time_test

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func init() {
	if time.ZoneinfoForTesting() != nil {
		panic(fmt.Errorf("zoneinfo initialized before first LoadLocation"))
	}
}

func TestEnvVarUsage(t *testing.T) {
	time.ResetZoneinfoForTesting()

	const testZoneinfo = "foo.zip"
	const env = "ZONEINFO"

	defer os.Setenv(env, os.Getenv(env))
	os.Setenv(env, testZoneinfo)

	// Result isn't important, we're testing the side effect of this command
	time.LoadLocation("Asia/Jerusalem")
	defer time.ResetZoneinfoForTesting()

	if zoneinfo := time.ZoneinfoForTesting(); testZoneinfo != *zoneinfo {
		t.Errorf("zoneinfo does not match env variable: got %q want %q", *zoneinfo, testZoneinfo)
	}
}

func TestBadLocationErrMsg(t *testing.T) {
	time.ResetZoneinfoForTesting()
	loc := "Asia/SomethingNotExist"
	want := errors.New("unknown time zone " + loc)
	_, err := time.LoadLocation(loc)
	if err.Error() != want.Error() {
		t.Errorf("LoadLocation(%q) error = %v; want %v", loc, err, want)
	}
}

func TestLoadLocationValidatesNames(t *testing.T) {
	time.ResetZoneinfoForTesting()
	const env = "ZONEINFO"
	defer os.Setenv(env, os.Getenv(env))
	os.Setenv(env, "")

	bad := []string{
		"/usr/foo/Foo",
		"\\UNC\foo",
		"..",
		"a..",
	}
	for _, v := range bad {
		_, err := time.LoadLocation(v)
		if err != time.ErrLocation {
			t.Errorf("LoadLocation(%q) error = %v; want ErrLocation", v, err)
		}
	}
}

func TestVersion3(t *testing.T) {
	time.ForceZipFileForTesting(true)
	defer time.ForceZipFileForTesting(false)
	_, err := time.LoadLocation("Asia/Jerusalem")
	if err != nil {
		t.Fatal(err)
	}
}

// Test that we get the correct results for times before the first
// transition time. To do this we explicitly check early dates in a
// couple of specific timezones.
func TestFirstZone(t *testing.T) {
	time.ForceZipFileForTesting(true)
	defer time.ForceZipFileForTesting(false)

	const format = "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"
	var tests = []struct {
		zone  string
		unix  int64
		want1 string
		want2 string
	}{
		{
			"PST8PDT",
			-1633269601,
			"Sun, 31 Mar 1918 01:59:59 -0800 (PST)",
			"Sun, 31 Mar 1918 03:00:00 -0700 (PDT)",
		},
		{
			"Pacific/Fakaofo",
			1325242799,
			"Thu, 29 Dec 2011 23:59:59 -1100 (-11)",
			"Sat, 31 Dec 2011 00:00:00 +1300 (+13)",
		},
	}

	for _, test := range tests {
		z, err := time.LoadLocation(test.zone)
		if err != nil {
			t.Fatal(err)
		}
		s := time.Unix(test.unix, 0).In(z).Format(format)
		if s != test.want1 {
			t.Errorf("for %s %d got %q want %q", test.zone, test.unix, s, test.want1)
		}
		s = time.Unix(test.unix+1, 0).In(z).Format(format)
		if s != test.want2 {
			t.Errorf("for %s %d got %q want %q", test.zone, test.unix, s, test.want2)
		}
	}
}

func TestLocationNames(t *testing.T) {
	if time.Local.String() != "Local" {
		t.Errorf(`invalid Local location name: got %q want "Local"`, time.Local)
	}
	if time.UTC.String() != "UTC" {
		t.Errorf(`invalid UTC location name: got %q want "UTC"`, time.UTC)
	}
}

func TestLoadLocationFromTZData(t *testing.T) {
	time.ForceZipFileForTesting(true)
	defer time.ForceZipFileForTesting(false)

	const locationName = "Asia/Jerusalem"
	reference, err := time.LoadLocation(locationName)
	if err != nil {
		t.Fatal(err)
	}

	tzinfo, err := time.LoadTzinfo(locationName, time.OrigZoneSources[len(time.OrigZoneSources)-1])
	if err != nil {
		t.Fatal(err)
	}
	sample, err := time.LoadLocationFromTZData(locationName, tzinfo)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(reference, sample) {
		t.Errorf("return values of LoadLocationFromTZData and LoadLocation don't match")
	}
}

// Issue 30099.
func TestEarlyLocation(t *testing.T) {
	time.ForceZipFileForTesting(true)
	defer time.ForceZipFileForTesting(false)

	const locName = "America/New_York"
	loc, err := time.LoadLocation(locName)
	if err != nil {
		t.Fatal(err)
	}

	d := time.Date(1900, time.January, 1, 0, 0, 0, 0, loc)
	tzName, tzOffset := d.Zone()
	if want := "EST"; tzName != want {
		t.Errorf("Zone name == %s, want %s", tzName, want)
	}
	if want := -18000; tzOffset != want {
		t.Errorf("Zone offset == %d, want %d", tzOffset, want)
	}
}

func TestMalformedTZData(t *testing.T) {
	// The goal here is just that malformed tzdata results in an error, not a panic.
	issue29437 := "TZif\x00000000000000000\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0000"
	_, err := time.LoadLocationFromTZData("abc", []byte(issue29437))
	if err == nil {
		t.Error("expected error, got none")
	}
}

var slimTests = []struct {
	zoneName   string
	tzData     string
	wantName   string
	wantOffset int
}{
	{
		// 2020b slim tzdata for Europe/Berlin.
		zoneName:   "Europe/Berlin",
		tzData:     "TZif2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00TZif2\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00<\x00\x00\x00\x04\x00\x00\x00\x12\xff\xff\xff\xffo\xa2a\xf8\xff\xff\xff\xff\x9b\f\x17`\xff\xff\xff\xff\x9b\xd5\xda\xf0\xff\xff\xff\xff\x9cٮ\x90\xff\xff\xff\xff\x9d\xa4\xb5\x90\xff\xff\xff\xff\x9e\xb9\x90\x90\xff\xff\xff\xff\x9f\x84\x97\x90\xff\xff\xff\xff\xc8\tq\x90\xff\xff\xff\xff\xcc\xe7K\x10\xff\xff\xff\xffͩ\x17\x90\xff\xff\xff\xff\u03a2C\x10\xff\xff\xff\xffϒ4\x10\xff\xff\xff\xffЂ%\x10\xff\xff\xff\xff\xd1r\x16\x10\xff\xff\xff\xffѶ\x96\x00\xff\xff\xff\xff\xd2X\xbe\x80\xff\xff\xff\xffҡO\x10\xff\xff\xff\xff\xd3c\x1b\x90\xff\xff\xff\xff\xd4K#\x90\xff\xff\xff\xff\xd59\xd1 \xff\xff\xff\xff\xd5g\xe7\x90\xff\xff\xff\xffըs\x00\xff\xff\xff\xff\xd6)\xb4\x10\xff\xff\xff\xff\xd7,\x1a\x10\xff\xff\xff\xff\xd8\t\x96\x10\xff\xff\xff\xff\xd9\x02\xc1\x90\xff\xff\xff\xff\xd9\xe9x\x10\x00\x00\x00\x00\x13MD\x10\x00\x00\x00\x00\x143\xfa\x90\x00\x00\x00\x00\x15#\xeb\x90\x00\x00\x00\x00\x16\x13ܐ\x00\x00\x00\x00\x17\x03͐\x00\x00\x00\x00\x17\xf3\xbe\x90\x00\x00\x00\x00\x18㯐\x00\x00\x00\x00\x19Ӡ\x90\x00\x00\x00\x00\x1aÑ\x90\x00\x00\x00\x00\x1b\xbc\xbd\x10\x00\x00\x00\x00\x1c\xac\xae\x10\x00\x00\x00\x00\x1d\x9c\x9f\x10\x00\x00\x00\x00\x1e\x8c\x90\x10\x00\x00\x00\x00\x1f|\x81\x10\x00\x00\x00\x00 lr\x10\x00\x00\x00\x00!\\c\x10\x00\x00\x00\x00\"LT\x10\x00\x00\x00\x00#<E\x10\x00\x00\x00\x00$,6\x10\x00\x00\x00\x00%\x1c'\x10\x00\x00\x00\x00&\f\x18\x10\x00\x00\x00\x00'\x05C\x90\x00\x00\x00\x00'\xf54\x90\x00\x00\x00\x00(\xe5%\x90\x00\x00\x00\x00)\xd5\x16\x90\x00\x00\x00\x00*\xc5\a\x90\x00\x00\x00\x00+\xb4\xf8\x90\x00\x00\x00\x00,\xa4\xe9\x90\x00\x00\x00\x00-\x94ڐ\x00\x00\x00\x00.\x84ː\x00\x00\x00\x00/t\xbc\x90\x00\x00\x00\x000d\xad\x90\x00\x00\x00\x001]\xd9\x10\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x03\x01\x02\x01\x02\x01\x03\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x00\x00\f\x88\x00\x00\x00\x00\x1c \x01\x04\x00\x00\x0e\x10\x00\t\x00\x00*0\x01\rLMT\x00CEST\x00CET\x00CEMT\x00\nCET-1CEST,M3.5.0,M10.5.0/3\n",
		wantName:   "CET",
		wantOffset: 3600,
	},
	{
		// 2021a slim tzdata for America/Nuuk.
		zoneName:   "America/Nuuk",
		tzData:     "TZif3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00TZif3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\"\x00\x00\x00\x03\x00\x00\x00\f\xff\xff\xff\xff\x9b\x80h\x00\x00\x00\x00\x00\x13M|P\x00\x00\x00\x00\x143\xfa\x90\x00\x00\x00\x00\x15#\xeb\x90\x00\x00\x00\x00\x16\x13ܐ\x00\x00\x00\x00\x17\x03͐\x00\x00\x00\x00\x17\xf3\xbe\x90\x00\x00\x00\x00\x18㯐\x00\x00\x00\x00\x19Ӡ\x90\x00\x00\x00\x00\x1aÑ\x90\x00\x00\x00\x00\x1b\xbc\xbd\x10\x00\x00\x00\x00\x1c\xac\xae\x10\x00\x00\x00\x00\x1d\x9c\x9f\x10\x00\x00\x00\x00\x1e\x8c\x90\x10\x00\x00\x00\x00\x1f|\x81\x10\x00\x00\x00\x00 lr\x10\x00\x00\x00\x00!\\c\x10\x00\x00\x00\x00\"LT\x10\x00\x00\x00\x00#<E\x10\x00\x00\x00\x00$,6\x10\x00\x00\x00\x00%\x1c'\x10\x00\x00\x00\x00&\f\x18\x10\x00\x00\x00\x00'\x05C\x90\x00\x00\x00\x00'\xf54\x90\x00\x00\x00\x00(\xe5%\x90\x00\x00\x00\x00)\xd5\x16\x90\x00\x00\x00\x00*\xc5\a\x90\x00\x00\x00\x00+\xb4\xf8\x90\x00\x00\x00\x00,\xa4\xe9\x90\x00\x00\x00\x00-\x94ڐ\x00\x00\x00\x00.\x84ː\x00\x00\x00\x00/t\xbc\x90\x00\x00\x00\x000d\xad\x90\x00\x00\x00\x001]\xd9\x10\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\xff\xffπ\x00\x00\xff\xff\xd5\xd0\x00\x04\xff\xff\xe3\xe0\x01\bLMT\x00-03\x00-02\x00\n<-03>3<-02>,M3.5.0/-2,M10.5.0/-1\n",
		wantName:   "-03",
		wantOffset: -10800,
	},
	{
		// 2021a slim tzdata for Asia/Gaza.
		zoneName:   "Asia/Gaza",
		tzData:     "TZif3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00TZif3\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00s\x00\x00\x00\x05\x00\x00\x00\x15\xff\xff\xff\xff}\xbdJ\xb0\xff\xff\xff\xff\xc8Y\xcf\x00\xff\xff\xff\xff\xc8\xfa\xa6\x00\xff\xff\xff\xff\xc98\x9c\x80\xff\xff\xff\xff\xcc\xe5\xeb\x80\xff\xff\xff\xffͬ\xfe\x00\xff\xff\xff\xff\xce\xc7\x1f\x00\xff\xff\xff\xffϏ\x83\x00\xff\xff\xff\xffЩ\xa4\x00\xff\xff\xff\xffф}\x00\xff\xff\xff\xffҊ׀\xff\xff\xff\xff\xd3e\xb0\x80\xff\xff\xff\xff\xd4l\v\x00\xff\xff\xff\xff\xe86c`\xff\xff\xff\xff\xe8\xf4-P\xff\xff\xff\xff\xea\v\xb9`\xff\xff\xff\xff\xea\xd5`\xd0\xff\xff\xff\xff\xeb\xec\xfa\xf0\xff\xff\xff\xff\xec\xb5m\x00\xff\xff\xff\xff\xed\xcf\u007f\xf0\xff\xff\xff\xff\xee\x97\xf2\x00\xff\xff\xff\xffﰳp\xff\xff\xff\xff\xf0y%\x80\xff\xff\xff\xff\xf1\x91\xe6\xf0\xff\xff\xff\xff\xf2ZY\x00\xff\xff\xff\xff\xf3s\x1ap\xff\xff\xff\xff\xf4;\x8c\x80\xff\xff\xff\xff\xf5U\x9fp\xff\xff\xff\xff\xf6\x1e\x11\x80\xff\xff\xff\xff\xf76\xd2\xf0\xff\xff\xff\xff\xf7\xffE\x00\xff\xff\xff\xff\xf9\x18\x06p\xff\xff\xff\xff\xf9\xe1\xca\x00\xff\xff\xff\xff\xfa\xf99\xf0\xff\xff\xff\xff\xfb'BP\x00\x00\x00\x00\b|\x8b\xe0\x00\x00\x00\x00\b\xfd\xb0\xd0\x00\x00\x00\x00\t\xf6\xea`\x00\x00\x00\x00\n\xa63\xd0\x00\x00\x00\x00\x13\xe9\xfc`\x00\x00\x00\x00\x14![`\x00\x00\x00\x00\x1a\xfa\xc6`\x00\x00\x00\x00\x1b\x8en`\x00\x00\x00\x00\x1c\xbe\xf8\xe0\x00\x00\x00\x00\x1dw|\xd0\x00\x00\x00\x00\x1e\xcc\xff`\x00\x00\x00\x00\x1f`\x99P\x00\x00\x00\x00 \x82\xb1`\x00\x00\x00\x00!I\xb5\xd0\x00\x00\x00\x00\"^\x9e\xe0\x00\x00\x00\x00# ]P\x00\x00\x00\x00$Z0`\x00\x00\x00\x00%\x00?P\x00\x00\x00\x00&\v\xed\xe0\x00\x00\x00\x00&\xd6\xe6\xd0\x00\x00\x00\x00'\xeb\xcf\xe0\x00\x00\x00\x00(\xc0\x03P\x00\x00\x00\x00)\xd4\xec`\x00\x00\x00\x00*\xa9\x1f\xd0\x00\x00\x00\x00+\xbbe\xe0\x00\x00\x00\x00,\x89\x01\xd0\x00\x00\x00\x00-\x9bG\xe0\x00\x00\x00\x00._\xa9P\x00\x00\x00\x00/{)\xe0\x00\x00\x00\x000H\xc5\xd0\x00\x00\x00\x000\xe7\a\xe0\x00\x00\x00\x001dF`\x00\x00\x00\x002A\xc2`\x00\x00\x00\x003D(`\x00\x00\x00\x004!\xa4`\x00\x00\x00\x005$\n`\x00\x00\x00\x006\x01\x86`\x00\x00\x00\x007\x16a`\x00\x00\x00\x008\x06DP\x00\x00\x00\x008\xff}\xe0\x00\x00\x00\x009\xef`\xd0\x00\x00\x00\x00:\xdf_\xe0\x00\x00\x00\x00;\xcfB\xd0\x00\x00\x00\x00<\xbfA\xe0\x00\x00\x00\x00=\xaf$\xd0\x00\x00\x00\x00>\x9f#\xe0\x00\x00\x00\x00?\x8f\x06\xd0\x00\x00\x00\x00@\u007f\x05\xe0\x00\x00\x00\x00A\\\x81\xe0\x00\x00\x00\x00B^\xe7\xe0\x00\x00\x00\x00CA\xb7\xf0\x00\x00\x00\x00D-\xa6`\x00\x00\x00\x00E\x12\xfdP\x00\x00\x00\x00F\x0e\xd9\xe0\x00\x00\x00\x00F\xe8op\x00\x00\x00\x00G\xec\x18\xe0\x00\x00\x00\x00H\xb7\x11\xd0\x00\x00\x00\x00I\xcb\xfa\xe0\x00\x00\x00\x00J\xa0<`\x00\x00\x00\x00K\xad.\x9c\x00\x00\x00\x00La\xbd\xd0\x00\x00\x00\x00M\x94\xf9\x9c\x00\x00\x00\x00N5\xc2P\x00\x00\x00\x00Ot\xdb`\x00\x00\x00\x00P[\x91\xe0\x00\x00\x00\x00QT\xbd`\x00\x00\x00\x00RD\xa0P\x00\x00\x00\x00S4\x9f`\x00\x00\x00\x00TIlP\x00\x00\x00\x00U\x15\xd2\xe0\x00\x00\x00\x00V)\\`\x00\x00\x00\x00V\xf5\xc2\xf0\x00\x00\x00\x00X\x13\xca`\x00\x00\x00\x00Xդ\xf0\x00\x00\x00\x00Y\xf3\xac`\x00\x00\x00\x00Z\xb5\x86\xf0\x00\x00\x00\x00[ӎ`\x00\x00\x00\x00\\\x9dC\xe0\x00\x00\x00\x00]\xb3bP\x00\x00\x00\x00^~w`\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x03\x04\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x02\x01\x00\x00 P\x00\x00\x00\x00*0\x01\x04\x00\x00\x1c \x00\t\x00\x00*0\x01\r\x00\x00\x1c \x00\x11LMT\x00EEST\x00EET\x00IDT\x00IST\x00\nEET-2EEST,M3.4.4/48,M10.4.4/49\n",
		wantName:   "EET",
		wantOffset: 7200,
	},
}

func TestLoadLocationFromTZDataSlim(t *testing.T) {
	for _, test := range slimTests {
		reference, err := time.LoadLocationFromTZData(test.zoneName, []byte(test.tzData))
		if err != nil {
			t.Fatal(err)
		}

		d := time.Date(2020, time.October, 29, 15, 30, 0, 0, reference)
		tzName, tzOffset := d.Zone()
		if tzName != test.wantName {
			t.Errorf("Zone name == %s, want %s", tzName, test.wantName)
		}
		if tzOffset != test.wantOffset {
			t.Errorf("Zone offset == %d, want %d", tzOffset, test.wantOffset)
		}
	}
}

func TestTzset(t *testing.T) {
	for _, test := range []struct {
		inStr string
		inEnd int64
		inSec int64
		name  string
		off   int
		start int64
		end   int64
		isDST bool
		ok    bool
	}{
		{"", 0, 0, "", 0, 0, 0, false, false},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2159200800, "PDT", -7 * 60 * 60, 2152173600, 2172733200, true, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2152173599, "PST", -8 * 60 * 60, 2145916800, 2152173600, false, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2152173600, "PDT", -7 * 60 * 60, 2152173600, 2172733200, true, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2152173601, "PDT", -7 * 60 * 60, 2152173600, 2172733200, true, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2172733199, "PDT", -7 * 60 * 60, 2152173600, 2172733200, true, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2172733200, "PST", -8 * 60 * 60, 2172733200, 2177452800, false, true},
		{"PST8PDT,M3.2.0,M11.1.0", 0, 2172733201, "PST", -8 * 60 * 60, 2172733200, 2177452800, false, true},
	} {
		name, off, start, end, isDST, ok := time.Tzset(test.inStr, test.inEnd, test.inSec)
		if name != test.name || off != test.off || start != test.start || end != test.end || ok != test.ok {
			t.Errorf("tzset(%q, %d, %d) = %q, %d, %d, %d, %t, %t, want %q, %d, %d, %d, %t, %t", test.inStr, test.inEnd, test.inSec, name, off, start, end, isDST, ok, test.name, test.off, test.start, test.end, test.isDST, test.ok)
		}
	}
}

func TestTzsetName(t *testing.T) {
	for _, test := range []struct {
		in   string
		name string
		out  string
		ok   bool
	}{
		{"", "", "", false},
		{"X", "", "", false},
		{"PST", "PST", "", true},
		{"PST8PDT", "PST", "8PDT", true},
		{"PST-08", "PST", "-08", true},
		{"<A+B>+08", "A+B", "+08", true},
	} {
		name, out, ok := time.TzsetName(test.in)
		if name != test.name || out != test.out || ok != test.ok {
			t.Errorf("tzsetName(%q) = %q, %q, %t, want %q, %q, %t", test.in, name, out, ok, test.name, test.out, test.ok)
		}
	}
}

func TestTzsetOffset(t *testing.T) {
	for _, test := range []struct {
		in  string
		off int
		out string
		ok  bool
	}{
		{"", 0, "", false},
		{"X", 0, "", false},
		{"+", 0, "", false},
		{"+08", 8 * 60 * 60, "", true},
		{"-01:02:03", -1*60*60 - 2*60 - 3, "", true},
		{"01", 1 * 60 * 60, "", true},
		{"100", 100 * 60 * 60, "", true},
		{"1000", 0, "", false},
		{"8PDT", 8 * 60 * 60, "PDT", true},
	} {
		off, out, ok := time.TzsetOffset(test.in)
		if off != test.off || out != test.out || ok != test.ok {
			t.Errorf("tzsetName(%q) = %d, %q, %t, want %d, %q, %t", test.in, off, out, ok, test.off, test.out, test.ok)
		}
	}
}

func TestTzsetRule(t *testing.T) {
	for _, test := range []struct {
		in  string
		r   time.Rule
		out string
		ok  bool
	}{
		{"", time.Rule{}, "", false},
		{"X", time.Rule{}, "", false},
		{"J10", time.Rule{Kind: time.RuleJulian, Day: 10, Time: 2 * 60 * 60}, "", true},
		{"20", time.Rule{Kind: time.RuleDOY, Day: 20, Time: 2 * 60 * 60}, "", true},
		{"M1.2.3", time.Rule{Kind: time.RuleMonthWeekDay, Mon: 1, Week: 2, Day: 3, Time: 2 * 60 * 60}, "", true},
		{"30/03:00:00", time.Rule{Kind: time.RuleDOY, Day: 30, Time: 3 * 60 * 60}, "", true},
		{"M4.5.6/03:00:00", time.Rule{Kind: time.RuleMonthWeekDay, Mon: 4, Week: 5, Day: 6, Time: 3 * 60 * 60}, "", true},
		{"M4.5.7/03:00:00", time.Rule{}, "", false},
		{"M4.5.6/-04", time.Rule{Kind: time.RuleMonthWeekDay, Mon: 4, Week: 5, Day: 6, Time: -4 * 60 * 60}, "", true},
	} {
		r, out, ok := time.TzsetRule(test.in)
		if r != test.r || out != test.out || ok != test.ok {
			t.Errorf("tzsetName(%q) = %#v, %q, %t, want %#v, %q, %t", test.in, r, out, ok, test.r, test.out, test.ok)
		}
	}
}

func TestIsDST(t *testing.T) {
	time.ForceZipFileForTesting(true)
	defer time.ForceZipFileForTesting(false)

	tzWithDST, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		t.Error("could not load tz 'Australia/Sydney'")
	}
	tzWithoutDST, err := time.LoadLocation("Australia/Brisbane")
	if err != nil {
		t.Error("could not load tz 'Australia/Sydney'")
	}
	tzFixed := time.FixedZone("FIXED_TIME", 12345)

	for _, test := range []struct {
		loc   *time.Location
		t     time.Time
		isDST bool
	}{
		{time.UTC, time.Date(2009, 6, 1, 12, 0, 0, 0, time.UTC), false},
		{time.UTC, time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{tzWithDST, time.Date(2009, 6, 1, 12, 0, 0, 0, time.UTC), true},
		{tzWithDST, time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{tzWithoutDST, time.Date(2009, 6, 1, 12, 0, 0, 0, time.UTC), false},
		{tzWithoutDST, time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC), false},
		{tzFixed, time.Date(2009, 6, 1, 12, 0, 0, 0, time.UTC), false},
		{tzFixed, time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC), false},
	} {
		isDST := test.loc.IsDST(test.t)
		if isDST != test.isDST {
			t.Errorf("(%#v).IsDST(%#v) = %#v, want %#v", test.loc, test.t, isDST, test.isDST)
		}
	}
}
