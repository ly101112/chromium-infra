// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package iputil

import (
	"fmt"
	"math/big"
	"net"
)

// IsCanonicalIP returns true if the underlying IP object is 16 bytes long.
//
// Internally, net.IP can be 4 bytes or it can be 16 bytes long.
// IPv4 addresses also have a 16 byte representation with a specific prefix.
func IsCanonicalIP(ip net.IP) bool {
	return len(ip) == 16
}

// MustParseIP parses an IP address and panics if it's invalid.
func MustParseIP(x string) net.IP {
	ip := net.ParseIP(x)
	if ip == nil {
		panic(fmt.Sprintf("invalid ip address: %q", x))
	}
	return ip
}

// incrByte takes a byte, increments it, and returns a boolean indicating whether it overflowed or not.
func incrByte(x byte) (res byte, overflow bool) {
	return x + 1, x == 255
}

// RawIncr takes an IP address and increments it in an abstraction-breaking way. It doesn't respect submasks, for example.
func RawIncr(ip net.IP) (res net.IP, overflow bool) {
	overflow = true
	if len(ip) == 0 {
		return
	}
	res = make([]byte, len(ip))
	if n := copy(res, ip); n != len(ip) {
		panic("internal error in ../util/iputil/iptuil.go")
	}
	for i := -1 + len(ip); i >= 0; i-- {
		if !overflow {
			break
		}
		res[i], overflow = incrByte(ip[i])
	}
	return
}

// AddToIP adds an arbitrary integer to an IP and returns the empty IP if the result would be negative.
func AddToIP(ip net.IP, offset *big.Int) net.IP {
	ipAsInt := big.NewInt(0)
	ipAsInt.SetBytes(ip)
	ipAsInt.Add(ipAsInt, offset)
	if isNegative(ipAsInt) {
		return nil
	}
	return pad(ipAsInt.Bytes(), len(ip))
}

func pad(x []byte, n int) []byte {
	if len(x) == n {
		return x
	}
	if n <= 0 {
		return nil
	}
	out := make([]byte, n)

	for i := 0; i < n; i++ {
		j := n - i - 1
		k := len(x) - i - 1
		if k < 0 {
			return out
		}
		out[j] = x[k]
	}

	return out
}

func isNegative(item *big.Int) bool {
	return item.Sign() == -1
}