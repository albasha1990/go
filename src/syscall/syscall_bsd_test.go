// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd openbsd

package syscall_test

import (
	"syscall"
	"testing"
)

const MNT_WAIT = 1

func TestGetfsstat(t *testing.T) {
	n, err := syscall.Getfsstat(nil, MNT_WAIT)
	t.Logf("Getfsstat(nil, MNT_WAIT) = (%v, %v)", n, err)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]syscall.Statfs_t, n)
	n, err = syscall.Getfsstat(data, MNT_WAIT)
	t.Logf("Getfsstat([]syscall.Statfs_t, MNT_WAIT) = (%v, %v)", n, err)
	if err != nil {
		t.Fatal(err)
	}

	empty := syscall.Statfs_t{}
	for i, stat := range data {
		if stat == empty {
			t.Errorf("index %v is an empty Statfs_t struct", i)
		}
	}
}
