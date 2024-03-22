/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"crypto/rand"
	"encoding/binary"
	"sync"
)

// IndexTableEntry represents an entry in the IndexTable, which maps a 32-bit
// index to a peer, handshake, and keypair.
type IndexTableEntry struct {
	peer      *Peer
	handshake *Handshake
	keypair   *Keypair
}

// IndexTable is a thread-safe map that stores IndexTableEntry instances.
type IndexTable struct {
	sync.RWMutex // Provides read-write locking for safe concurrent access.
	table        map[uint32]IndexTableEntry
}

// randUint32 generates a random 32-bit unsigned integer with crypto-grade
// randomness.
func randUint32() (uint32, error) {
	var integer [4]byte
	_, err := rand.Read(integer[:])
	// Arbitrary endianness; both are intrinsified by the Go compiler.
	return binary.LittleEndian.Uint32(integer[:]), err
}

// Init initializes the IndexTable by creating a new, empty map.
func (table *IndexTable) Init() {
	table.Lock()
	defer table.Unlock()
	table.table = make(map[uint32]IndexTableEntry)
}

// Delete removes an entry from the IndexTable by index.
func (table *IndexTable) Delete(index uint32) {
	table.Lock()
	defer table.Unlock()
	delete(table.table, index)
}

// SwapIndexForKeypair replaces the keypair associated with an index in the
// IndexTable.
func (table *IndexTable) SwapIndexForKeypair(index uint32, keypair *Keypair) {
	table.Lock()
	defer table.Unlock()
	entry, ok := table.table[index]
	if !ok {
		return
	}
	table.table[index] = IndexTableEntry{
		peer:      entry.peer,
		keypair:   keypair,
		handshake: nil,
	}
}

// NewIndexForHandshake generates a new, unique 32-bit index and associates it
// with a peer and handshake in the IndexTable.
func (table *IndexTable) NewIndexForHandshake(peer *Peer, handshake *Handshake) (uint32, error) {
	for {
		// generate random index
		index, err := randUint32()
		if err != nil {
			return index, err
	
