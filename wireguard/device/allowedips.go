/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	// ... (imported packages)
)

// parentIndirection is a helper struct used to keep track of a parent trieEntry
// and the bit index of the child trieEntry within the parent's child array.
type parentIndirection struct {
	parentBit     **trieEntry // Pointer to the parent trieEntry.
	parentBitType uint8       // The bit index of the child trieEntry within the parent's child array.
}

// trieEntry represents a node in the IP address prefix trie.
type trieEntry struct {
	peer        *Peer        // The associated Peer for this trieEntry.
	child       [2]*trieEntry // The child trieEntries for this node.
	parent      parentIndirection  // The parent trieEntry and the bit index of this trieEntry within the parent's child array.
	cidr        uint8          // The CIDR value for this trieEntry.
	bitAtByte   uint8          // The byte index of the bit used for this trieEntry.
	bitAtShift  uint8          // The bit shift value for this trieEntry.
	bits        []byte         // The IP address bits for this trieEntry.
	perPeerElem *list.Element // A doubly-linked list element used to keep track of the trieEntries for each peer.
}

// ... (function implementations)

// insert inserts a new trieEntry for the given IP address prefix and peer.
func (trie parentIndirection) insert(ip []byte, cidr uint8, peer *Peer) {
	// ... (function implementation)
}

// lookup finds the trieEntry for the given IP address.
func (node *trieEntry) lookup(ip []byte) *Peer {
	// ... (function implementation)
}

// AllowedIPs is a struct that holds two IP address prefix tries, one for IPv4 and one for IPv6.
type AllowedIPs struct {
	IPv4  *trieEntry // The IPv4 address prefix trie.
	IPv6  *trieEntry // The IPv6 address prefix trie.
	mutex sync.RWMutex // A mutex to synchronize access to the tries.
}

// ... (function implementations)

// Insert inserts a new IP address prefix and peer into the appropriate trie.
func (table *AllowedIPs) Insert(prefix netip.Prefix, peer *Peer) {
	// ... (function implementation)
}

// Lookup finds the peer associated with the given IP address.
func (table *AllowedIPs) Lookup(ip []byte) *Peer {
	// ... (function implementation)
}

