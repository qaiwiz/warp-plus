/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package replay

// Package replay implements an efficient anti-replay algorithm as specified in RFC 6479.
package replay

type block uint64

const (
	// blockBitLog is the logarithm base 2 of the number of bits in a block.
	blockBitLog = 6                // 1<<6 == 64 bits
	// blockBits is the number of bits in a block (must be a power of 2).
	blockBits   = 1 << blockBitLog // must be power of 2
	// ringBlocks is the number of blocks in the ring buffer (must be a power of 2).
	ringBlocks  = 1 << 7           // must be power of 2
	// windowSize is the size of the sliding window in bits.
	windowSize  = (ringBlocks - 1) * blockBits
	// blockMask is a mask to get the block index from a counter value.
	blockMask   = ringBlocks - 1
	// bitMask is a mask to get the bit index within a block from a counter value.
	bitMask     = blockBits - 1
)

// A Filter rejects replayed messages by checking if message counter value is
// within a sliding window of previously received messages.
// The zero value for Filter is an empty filter ready to use.
// Filters are unsafe for concurrent use.
type Filter struct {
	// last is the index of the last processed message counter.
	last uint64
	// ring is a ring buffer of blocks used to store the sliding window.
	ring [ringBlocks]block
}

// Reset resets the filter to empty state.
func (f *Filter) Reset() {
	f.last = 0 // Reset the last processed message counter index.
	f.ring[0] = 0 // Reset the ring buffer to all zero blocks.
}

// ValidateCounter checks if the counter should be accepted based on the sliding window.
// Overlimit counters (>= limit) are always rejected.
func (f *Filter) ValidateCounter(counter, limit uint64) bool {
	if counter >= limit { // Reject overlimit counters.
		return false
	}

	indexBlock := counter >> blockBitLog // Get the block index from the counter value.

	if counter > f.last { // Move the window forward.
		current := f.last >> blockBitLog // Get the current block index.
		diff := indexBlock - current   // Calculate the difference between the new and current block indices.
		if diff > ringBlocks {
			diff = ringBlocks // Cap diff to clear the whole ring.
		}
		for i := current +
