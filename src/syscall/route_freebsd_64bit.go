// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd,amd64

package syscall

import "unsafe"

func (any *anyMessage) parseRouteMessage(b []byte) *RouteMessage {
	p := (*RouteMessage)(unsafe.Pointer(any))
	return &RouteMessage{Header: p.Header, Data: b[rsaAlignOf(int(unsafe.Offsetof(p.Header.Rmx))+SizeofRtMetrics):any.Msglen]}
}
