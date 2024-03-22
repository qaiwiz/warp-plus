// SPDX-License-Identifier: MIT
// Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.

package ratelimiter

import (
	"net/netip"
	"sync"
	"time"
)

// Constants
const (
	packetsPerSecond   = 20      // Number of packets allowed per second.
	packetsBurstable   = 5       // Burstable capacity for packets.
	garbageCollectTime = time.Second // Time duration for garbage collection.
	packetCost         = 1000000000 / packetsPerSecond // Cost of a packet in nanoseconds.
	maxTokens          = packetCost * packetsBurstable // Maximum number of tokens that can be stored.
)

// RatelimiterEntry represents a single entry in the rate limiter table.
type RatelimiterEntry struct {
	mu       sync.Mutex // Mutex to protect the entry.
	lastTime time.Time  // Last time the entry was updated.
	tokens   int64      // Number of available tokens.
}

// Ratelimiter represents the rate limiter.
type Ratelimiter struct {
	mu      sync.RWMutex      // Mutex to protect the rate limiter.
	timeNow func() time.Time  // Function to get the current time.
	stopReset chan struct{} // Channel to stop and reset the rate limiter.
	table     map[netip.Addr]*RatelimiterEntry // Table of entries.
}

// Close stops the rate limiter.
func (rate *Ratelimiter) Close() {
	rate.mu.Lock()
	defer rate.mu.Unlock()

	// Close the stopReset channel to stop the garbage collection routine.
	if rate.stopReset != nil {
		close(rate.stopReset)
	}
}

// Init initializes the rate limiter.
func (rate *Ratelimiter) Init() {
	rate.mu.Lock()
	defer rate.mu.Unlock()

	// Set the timeNow function if it is not already set.
	if rate.timeNow == nil {
		rate.timeNow = time.Now
	}

	// Stop any ongoing garbage collection routine.
	if rate.stopReset != nil {
		close(rate.stopReset)
	}

	// Create a new stopReset channel and table.
	rate.stopReset = make(chan struct{})
	rate.table = make(map[netip.Addr]*RatelimiterEntry)

	// Start a new garbage collection routine.
	go rate.garbageCollect()
}

// garbageCollect performs garbage collection on the table.
func (rate *Ratelimiter) garbageCollect() {
	ticker
