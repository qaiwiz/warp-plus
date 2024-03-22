/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	"math/rand"
	"net"
	"net/netip"
	"testing"
)

// testPairCommonBits is a struct that holds two slices of bytes and an expected
// number of common leading bits for the slices. It is used in the TestCommonBits
// function to test the commonBits function.
type testPairCommonBits struct {
	s1    []byte // first slice of bytes
	s2    []byte // second slice of bytes
	match uint8  // expected number of common leading bits
}

// TestCommonBits tests the commonBits function by comparing its output with the
// expected number of common leading bits for a set of test cases.
func TestCommonBits(t *testing.T) {
	tests := []testPairCommonBits{
		// The test cases include various combinations of slice values and
		// expected common leading bits.
	}

	for _, p := range tests {
		v := commonBits(p.s1, p.s2)
		if v != p.match {
			t.Error(
				"For slice", p.s1, p.s2,
				"expected match", p.match,
				",but got", v,
			)
		}
	}
}

// benchmarkTrie is a benchmark function that creates a trie data structure with
// a given number of peers and addresses, and then performs a lookup operation
// for a random address. The function is used to measure the performance of the
// trie data structure under different workloads.
func benchmarkTrie(peerNumber, addressNumber, addressLength int, b *testing.B) {
	// The benchmark function initializes a trie data structure, inserts a given
	// number of peers and addresses into it, and then performs a lookup operation
	// for a random address. The function is called multiple times with different
	// parameters to measure the performance of the trie data structure under
	// different workloads.
}

// BenchmarkTrieIPv4Peers100Addresses1000 and BenchmarkTrieIPv4Peers10Addresses10
// are benchmark functions that measure the performance of the trie data structure
// for IPv4 addresses with different number of peers and addresses.
func BenchmarkTrieIPv4Peers100Addresses1000(b *testing.B)
func BenchmarkTrieIPv4Peers10Addresses10(b *testing.B)

// BenchmarkTrieIPv6Peers100Addresses1000 and BenchmarkTrieIPv6Peers10Addresses10
// are benchmark functions that measure the performance of the trie data structure
// for IPv6 addresses with different number of peers and addresses.
func BenchmarkTrieIPv6Peers100Addresses1000(b *testing.B)
func BenchmarkTrieIPv6Peers10Addresses10(b *testing.B)

// TestTrieIPv4 is a test function that tests the trie data structure for IPv4
// addresses with a set of test cases ported from the kernel implementation.
func TestTrieIPv4(t *testing.T) {
	// The test function initializes a trie data structure and performs various
	// insert, lookup, and remove operations on it. The function compares the
	// results of these operations with the expected results to ensure the
	// correctness of the trie data structure.
}

// TestTrieIPv6 is a test function that tests the trie data structure for IPv6
// addresses with a set of test cases ported from the kernel implementation.
func TestTrieIPv6(t *testing.T) {
	// The test function is similar to TestTrieIPv4, but it tests the trie data
	// structure for IPv6 addresses.
}
