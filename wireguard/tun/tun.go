/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package tun

import (
	"os" // For working with files
)

// Event represents different types of events that can occur on a Device
type Event int

// Constants for each type of Event
const (
	EventUp Event = 1 << iota // EventUp represents the device being brought up
	EventDown                 // EventDown represents the device being brought down
	EventMTUUpdate            // EventMTUUpdate represents an update to the device's MTU
)

// Device is an interface that represents a network device
type Device interface {
	// File returns the file descriptor of the device
	File() *os.File

	// Read reads one or more packets from the Device
	Read(bufs [][]byte, sizes []int, offset int) (n int, err error)

	// Write writes one or more packets to the device
	Write(bufs [][]byte, offset int) (int, error)

	// MTU returns the MTU of the Device
	MTU() (int, error)

	// Name returns the current name of the Device
	Name() (string, error)

	// Events returns a channel of type Event, which is fed Device events
	Events() <-chan Event

	// Close stops the Device and closes the Event channel
	Close() error

	// BatchSize returns the preferred/max number of packets that can be read or
	// written in a single read/write call. BatchSize must not change over the
	// lifetime of a Device.
	BatchSize() int
}

