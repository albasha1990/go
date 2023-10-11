// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is called cgo_unix.go, but to allow syscalls-to-libc-based
// implementations to share the code, it does not use cgo directly.
// Instead of C.foo it uses _C_foo, which is defined in either
// cgo_unix_cgo.go or cgo_unix_syscall.go

//go:build !netgo && ((cgo && unix) || darwin)

package net

import (
	"context"
	"errors"
	"net/netip"
	"syscall"
	"unsafe"

	"golang.org/x/net/dns/dnsmessage"
)

// cgoAvailable set to true to indicate that the cgo resolver
// is available on this system.
const cgoAvailable = true

// An addrinfoErrno represents a getaddrinfo, getnameinfo-specific
// error number. It's a signed number and a zero value is a non-error
// by convention.
type addrinfoErrno int

func (eai addrinfoErrno) Error() string   { return _C_gai_strerror(_C_int(eai)) }
func (eai addrinfoErrno) Temporary() bool { return eai == _C_EAI_AGAIN }
func (eai addrinfoErrno) Timeout() bool   { return false }

// isAddrinfoErrno is just for testing purposes.
func (eai addrinfoErrno) isAddrinfoErrno() {}

// doBlockingWithCtx executes a blocking function in a separate goroutine when the provided
// context is cancellable. It is intended for use with calls that don't support context
// cancellation (cgo, syscalls). blocking func may still be running after this function finishes.
func doBlockingWithCtx[T any](ctx context.Context, blocking func() (T, error)) (T, error) {
	if ctx.Done() == nil {
		return blocking()
	}

	type result struct {
		res T
		err error
	}

	res := make(chan result, 1)
	go func() {
		var r result
		r.res, r.err = blocking()
		res <- r
	}()

	select {
	case r := <-res:
		return r.res, r.err
	case <-ctx.Done():
		var zero T
		return zero, mapErr(ctx.Err())
	}
}

func cgoLookupHost(ctx context.Context, name string) (hosts []string, err error) {
	addrs, err := cgoLookupIP(ctx, "ip", name)
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		hosts = append(hosts, addr.String())
	}
	return hosts, nil
}

func cgoLookupPort(ctx context.Context, network, service string) (port int, err error) {
	var hints _C_struct_addrinfo
	switch network {
	case "ip": // no hints
	case "tcp", "tcp4", "tcp6":
		*_C_ai_socktype(&hints) = _C_SOCK_STREAM
		*_C_ai_protocol(&hints) = _C_IPPROTO_TCP
	case "udp", "udp4", "udp6":
		*_C_ai_socktype(&hints) = _C_SOCK_DGRAM
		*_C_ai_protocol(&hints) = _C_IPPROTO_UDP
	default:
		return 0, &DNSError{Err: "unknown network", Name: network + "/" + service}
	}
	switch ipVersion(network) {
	case '4':
		*_C_ai_family(&hints) = _C_AF_INET
	case '6':
		*_C_ai_family(&hints) = _C_AF_INET6
	}

	return doBlockingWithCtx(ctx, func() (int, error) {
		return cgoLookupServicePort(&hints, network, service)
	})
}

func cgoLookupServicePort(hints *_C_struct_addrinfo, network, service string) (port int, err error) {
	cservice, err := syscall.ByteSliceFromString(service)
	if err != nil {
		return 0, &DNSError{Err: err.Error(), Name: network + "/" + service}
	}
	// Lowercase the C service name.
	for i, b := range cservice[:len(service)] {
		cservice[i] = lowerASCII(b)
	}
	var res *_C_struct_addrinfo
	gerrno, err := _C_getaddrinfo(nil, (*_C_char)(unsafe.Pointer(&cservice[0])), hints, &res)
	if gerrno != 0 {
		isTemporary := false
		switch gerrno {
		case _C_EAI_SYSTEM:
			if err == nil { // see golang.org/issue/6232
				err = syscall.EMFILE
			}
		case _C_EAI_SERVICE, _C_EAI_NONAME: // Darwin returns EAI_NONAME.
			return 0, &DNSError{Err: "unknown port", Name: network + "/" + service, IsNotFound: true}
		default:
			err = addrinfoErrno(gerrno)
			isTemporary = addrinfoErrno(gerrno).Temporary()
		}
		return 0, &DNSError{Err: err.Error(), Name: network + "/" + service, IsTemporary: isTemporary}
	}
	defer _C_freeaddrinfo(res)

	for r := res; r != nil; r = *_C_ai_next(r) {
		switch *_C_ai_family(r) {
		case _C_AF_INET:
			sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
			p := (*[2]byte)(unsafe.Pointer(&sa.Port))
			return int(p[0])<<8 | int(p[1]), nil
		case _C_AF_INET6:
			sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
			p := (*[2]byte)(unsafe.Pointer(&sa.Port))
			return int(p[0])<<8 | int(p[1]), nil
		}
	}
	return 0, &DNSError{Err: "unknown port", Name: network + "/" + service, IsNotFound: true}
}

func cgoLookupHostIP(network, name string) (addrs []IPAddr, cname string, err error) {
	acquireThread()
	defer releaseThread()

	var hints _C_struct_addrinfo
	*_C_ai_flags(&hints) = cgoAddrInfoFlags
	*_C_ai_socktype(&hints) = _C_SOCK_STREAM
	*_C_ai_family(&hints) = _C_AF_UNSPEC
	switch ipVersion(network) {
	case '4':
		*_C_ai_family(&hints) = _C_AF_INET
	case '6':
		*_C_ai_family(&hints) = _C_AF_INET6
	}

	h, err := syscall.BytePtrFromString(name)
	if err != nil {
		return nil, "", &DNSError{Err: err.Error(), Name: name}
	}
	var res *_C_struct_addrinfo
	gerrno, err := _C_getaddrinfo((*_C_char)(unsafe.Pointer(h)), nil, &hints, &res)
	if gerrno != 0 {
		isErrorNoSuchHost := false
		isTemporary := false
		switch gerrno {
		case _C_EAI_SYSTEM:
			if err == nil {
				// err should not be nil, but sometimes getaddrinfo returns
				// gerrno == _C_EAI_SYSTEM with err == nil on Linux.
				// The report claims that it happens when we have too many
				// open files, so use syscall.EMFILE (too many open files in system).
				// Most system calls would return ENFILE (too many open files),
				// so at the least EMFILE should be easy to recognize if this
				// comes up again. golang.org/issue/6232.
				err = syscall.EMFILE
			}
		case _C_EAI_NONAME, _C_EAI_NODATA:
			err = errNoSuchHost
			isErrorNoSuchHost = true
		default:
			err = addrinfoErrno(gerrno)
			isTemporary = addrinfoErrno(gerrno).Temporary()
		}

		return nil, "", &DNSError{Err: err.Error(), Name: name, IsNotFound: isErrorNoSuchHost, IsTemporary: isTemporary}
	}
	defer _C_freeaddrinfo(res)

	if res != nil {
		cname = _C_GoString(*_C_ai_canonname(res))
		if cname == "" {
			cname = name
		}
		if len(cname) > 0 && cname[len(cname)-1] != '.' {
			cname += "."
		}
	}

	for r := res; r != nil; r = *_C_ai_next(r) {
		// We only asked for SOCK_STREAM, but check anyhow.
		if *_C_ai_socktype(r) != _C_SOCK_STREAM {
			continue
		}
		switch *_C_ai_family(r) {
		case _C_AF_INET:
			sa := (*syscall.RawSockaddrInet4)(unsafe.Pointer(*_C_ai_addr(r)))
			addr := IPAddr{IP: copyIP(sa.Addr[:])}
			addrs = append(addrs, addr)
		case _C_AF_INET6:
			sa := (*syscall.RawSockaddrInet6)(unsafe.Pointer(*_C_ai_addr(r)))
			addr := IPAddr{IP: copyIP(sa.Addr[:]), Zone: zoneCache.name(int(sa.Scope_id))}
			addrs = append(addrs, addr)
		}
	}
	return addrs, cname, nil
}

func cgoLookupIP(ctx context.Context, network, name string) (addrs []IPAddr, err error) {
	return doBlockingWithCtx(ctx, func() ([]IPAddr, error) {
		addrs, _, err := cgoLookupHostIP(network, name)
		return addrs, err
	})
}

// These are roughly enough for the following:
//
//	 Source		Encoding			Maximum length of single name entry
//	 Unicast DNS		ASCII or			<=253 + a NUL terminator
//				Unicode in RFC 5892		252 * total number of labels + delimiters + a NUL terminator
//	 Multicast DNS	UTF-8 in RFC 5198 or		<=253 + a NUL terminator
//				the same as unicast DNS ASCII	<=253 + a NUL terminator
//	 Local database	various				depends on implementation
const (
	nameinfoLen    = 64
	maxNameinfoLen = 4096
)

func cgoLookupPTR(ctx context.Context, addr string) (names []string, err error) {
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		return nil, &DNSError{Err: "invalid address", Name: addr}
	}
	sa, salen := cgoSockaddr(IP(ip.AsSlice()), ip.Zone())
	if sa == nil {
		return nil, &DNSError{Err: "invalid address " + ip.String(), Name: addr}
	}

	return doBlockingWithCtx(ctx, func() ([]string, error) {
		return cgoLookupAddrPTR(addr, sa, salen)
	})
}

func cgoLookupAddrPTR(addr string, sa *_C_struct_sockaddr, salen _C_socklen_t) (names []string, err error) {
	acquireThread()
	defer releaseThread()

	var gerrno int
	var b []byte
	for l := nameinfoLen; l <= maxNameinfoLen; l *= 2 {
		b = make([]byte, l)
		gerrno, err = cgoNameinfoPTR(b, sa, salen)
		if gerrno == 0 || gerrno != _C_EAI_OVERFLOW {
			break
		}
	}
	if gerrno != 0 {
		isErrorNoSuchHost := false
		isTemporary := false
		switch gerrno {
		case _C_EAI_SYSTEM:
			if err == nil { // see golang.org/issue/6232
				err = syscall.EMFILE
			}
		case _C_EAI_NONAME:
			err = errNoSuchHost
			isErrorNoSuchHost = true
		default:
			err = addrinfoErrno(gerrno)
			isTemporary = addrinfoErrno(gerrno).Temporary()
		}
		return nil, &DNSError{Err: err.Error(), Name: addr, IsTemporary: isTemporary, IsNotFound: isErrorNoSuchHost}
	}
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			b = b[:i]
			break
		}
	}
	return []string{absDomainName(string(b))}, nil
}

func cgoSockaddr(ip IP, zone string) (*_C_struct_sockaddr, _C_socklen_t) {
	if ip4 := ip.To4(); ip4 != nil {
		return cgoSockaddrInet4(ip4), _C_socklen_t(syscall.SizeofSockaddrInet4)
	}
	if ip6 := ip.To16(); ip6 != nil {
		return cgoSockaddrInet6(ip6, zoneCache.index(zone)), _C_socklen_t(syscall.SizeofSockaddrInet6)
	}
	return nil, 0
}

// cgoLookupCanonicalName returns the host canonical name.
func cgoLookupCanonicalName(ctx context.Context, network string, name string) (cname string, err error) {
	return doBlockingWithCtx(ctx, func() (string, error) {
		_, cname, err := cgoLookupHostIP(network, name)
		return cname, err
	})
}

// cgoLookupCNAME queries the CNAME resource using cgo resSearch.
// It returns the last CNAME found in the entire CNAME chain or the queried name when
// query returns with no answer resources.
func cgoLookupCNAME(ctx context.Context, name string) (cname string, err error) {
	msg, err := resSearch(ctx, name, int(dnsmessage.TypeCNAME), int(dnsmessage.ClassINET))

	noData := false
	if err != nil {
		var dnsErr *DNSError
		if !errors.As(err, &dnsErr) {
			// Not a DNS error.
			return "", err
		} else if dnsErr.isNoData && msg != nil {
			// DNS query succeeded, without error code (like NXDOMAIN),
			// but it has zero answer records.
			noData = true
		} else {
			return "", err
		}
	}

	var p dnsmessage.Parser
	_, err = p.Start(msg)
	if err != nil {
		return "", &DNSError{Err: errCannotUnmarshalDNSMessage.Error(), Name: name}
	}

	q, err := p.Question()
	if err != nil {
		return "", &DNSError{Err: errCannotUnmarshalDNSMessage.Error(), Name: name}
	}

	// Multiple questions, this should never happen.
	if err := p.SkipQuestion(); err != dnsmessage.ErrSectionDone {
		return "", &DNSError{Err: errCannotUnmarshalDNSMessage.Error(), Name: name}
	}

	if noData {
		return q.Name.String(), nil
	}

	// Using name from question, not the one provided in function arguments,
	// because of possible search domain in resolv.conf.
	cname, err = lastCNAMEinChain(q.Name, p)
	if err != nil {
		return "", &DNSError{
			Err:  err.Error(),
			Name: name,
		}
	}

	return cname, nil
}

// errCgoDNSLookupFailed is returned from resSearch on systems with non thread safe h_errno.
var errCgoDNSLookupFailed = errors.New("res_nsearch lookup failed")

// resSearch will make a call to the 'res_nsearch' routine in the C library
// and parse the output as a slice of DNS resources.
// In case of an error, the msg might be populated with a raw DNS response (it might
// be partial or with junk after the DNS message).
func resSearch(ctx context.Context, hostname string, rtype, class int) (msg []byte, err error) {
	return doBlockingWithCtx(ctx, func() ([]byte, error) {
		return cgoResSearch(hostname, rtype, class)
	})
}

func cgoResSearch(hostname string, rtype, class int) ([]byte, error) {
	acquireThread()
	defer releaseThread()

	var state *_C_struct___res_state
	if unsafe.Sizeof(_C_struct___res_state{}) != 0 {
		state = (*_C_struct___res_state)(_C_malloc(unsafe.Sizeof(_C_struct___res_state{})))
		defer _C_free(unsafe.Pointer(state))
		*state = _C_struct___res_state{}
	}

	if err := _C_res_ninit(state); err != nil {
		return nil, errors.New("res_ninit failure: " + err.Error())
	}
	defer _C_res_nclose(state)

	bufSize := maxDNSPacketSize
	buf := (*_C_uchar)(_C_malloc(uintptr(bufSize)))
	defer _C_free(unsafe.Pointer(buf))

	s, err := syscall.BytePtrFromString(hostname)
	if err != nil {
		return nil, err
	}

	var size int
	for {
		var herrno int
		var err error
		size, herrno, err = _C_res_nsearch(state, (*_C_char)(unsafe.Pointer(s)), class, rtype, buf, bufSize)
		if size <= 0 || size > 0xffff {
			// Copy from c to go memory.
			msgC := unsafe.Slice((*byte)(unsafe.Pointer(buf)), bufSize)
			msg := make([]byte, len(msgC))
			copy(msg, msgC)

			// We use -1 to indicate that h_errno is available, -2 otherwise.
			if size == -1 {
				if herrno == _C_HOST_NOT_FOUND || herrno == _C_NO_DATA {
					return msg, &DNSError{
						Err:        errNoSuchHost.Error(),
						IsNotFound: true,
						isNoData:   herrno == _C_NO_DATA,
						Name:       hostname,
					}
				}

				if err != nil {
					return msg, &DNSError{
						Err:         "dns lookup failure: " + err.Error(),
						IsTemporary: herrno == _C_TRY_AGAIN,
						Name:        hostname,
					}
				}

				return msg, &DNSError{
					Err:         "dns lookup failure",
					IsTemporary: herrno == _C_TRY_AGAIN,
					Name:        hostname,
				}
			}

			return msg, errCgoDNSLookupFailed
		}

		if size <= bufSize {
			break
		}

		// Allocate a bigger buffer to fit the entire msg.
		_C_free(unsafe.Pointer(buf))
		bufSize = size
		buf = (*_C_uchar)(_C_malloc(uintptr(bufSize)))
	}

	// Copy from c to go memory.
	msgC := unsafe.Slice((*byte)(unsafe.Pointer(buf)), size)
	msg := make([]byte, len(msgC))
	copy(msg, msgC)
	return msg, nil
}
