/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	// math/rand package provides random number generation functionality.
	"math/rand"
	// runtime package provides access to the Go runtime.
	"runtime"
	// sync package provides synchronization primitives, such as WaitGroup.
	"sync"
	// sync/atomic package provides low-level atomic memory update support.
	"sync/atomic"
	// testing package provides support for automated testing.
	"testing"
	// time package provides time-related functionality.
	"time"
)

// TestWaitPool tests the WaitPool functionality. It is currently disabled.
func TestWaitPool(t *testing.T) {
	t.Skip("Currently disabled")
	// Initialize a WaitGroup and an atomic int32 variable to store the number of trials.
	var wg sync.WaitGroup
	var trials atomic.Int32
	startTrials := int32(100000)
	if raceEnabled {
		// This test can be very slow with -race.
		startTrials /= 10
	}
	trials.Store(startTrials)
	// Calculate the number of workers based on the available CPU cores.
	workers := runtime.NumCPU() + 2
	if workers-4 <= 0 {
		t.Skip("Not enough cores")
	}
	// Create a new WaitPool with the specified number of workers and a factory function.
	p := NewWaitPool(uint32(workers-4), func() any { return make([]byte, 16) })
	// Add the required number of workers to the WaitGroup.
	wg.Add(workers)
	// Initialize an atomic uint32 variable to store the maximum count.
	var max atomic.Uint32
	// Define a function to update the maximum count.
	updateMax := func() {
		count := p.count.Load()
		if count > p.max {
			t.Errorf("count (%d) > max (%d)", count, p.max)
		}
		for {
			old := max.Load()
			if count <= old {
				break
			}
			if max.CompareAndSwap(old, count) {
				break
			}
		}
	}
	// Create worker goroutines.
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for trials.Add(-1) > 0 {
				updateMax()
				// Get an item from the WaitPool.
				x := p.Get()
				updateMax()
				// Sleep for a random duration.
				time.Sleep(time.Duration(rand.Intn(100)) * time.Microsecond)
				updateMax()
				// Put the item back into the WaitPool.
				p.Put(x)
				updateMax()
			}
		}()
	}
	// Wait for all worker goroutines to finish.
	wg.Wait()
	// Compare the actual maximum count with the ideal maximum count.
	if max.Load() != p.max {
		t.Errorf("Actual maximum count (%d) != ideal maximum count (%d)", max, p.max)
	}
}

// BenchmarkWaitPool benchmarks the WaitPool functionality.
func BenchmarkWaitPool(b *testing.B) {
	var wg sync.WaitGroup
	var trials atomic.Int32
	trials.Store(int32(b.N))
	workers := runtime.NumCPU() + 2
	if workers-4 <= 0 {
		b.Skip("Not enough cores")
	}
	p := NewWaitPool(uint32(workers-4), func() any { return make([]byte, 16) })
	wg.Add(workers)
	// Reset the timer for benchmarking.
	b.ResetTimer()
	// Create worker goroutines.
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for trials.Add(-1) > 0 {
				// Get an item from the WaitPool.
				x := p.Get()
				// Sleep for a random duration.
				
