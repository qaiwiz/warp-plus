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
	
