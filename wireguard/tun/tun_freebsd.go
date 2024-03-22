/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package tun

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Constants for various ioctl commands and flags
const (
	_TUNSIFHEAD = 0x80047460
	_TUNSIFMODE = 0x8004745e
	_TUNGIFNAME = 0x4020745d
	_TUNSIFPID  = 0x2000745f

	_SIOCGIFINFO_IN6        = 0xc048696c
	_SIOCSIFINFO_IN6        = 0xc048696d
	_ND6_IFF_AUTO_LINKLOCAL = 0x20
	_ND6_IFF_NO_DAD         = 0x100
)

// Structs for various ioctl requests
type ifreqName struct { // Represents an ioctl request with just the name
	Name [unix.IFNAMSIZ]byte
	_    [16]byte
}
type ifreqPtr struct { // Represents an ioctl request with a pointer
	Name [unix.IFNAMSIZ]byte
	Data uintptr
	_    [16 - unsafe.Sizeof(uintptr(0))]byte
}
type ifreqMtu struct { // Represents an ioctl request with MTU
	Name [unix.IFNAMSIZ]byte
	MTU  uint32
	_    [12]byte
}
type nd6Req struct { // Represents ND6 flag manipulation
	Name          [unix.IFNAMSIZ]byte
	Linkmtu       uint32
	Maxmtu        uint32
	Basereachable uint32
	Reachable     uint32
	Retrans       uint32
	Flags         uint32
	Recalctm      int
	Chlim         uint8
	Initialized   uint8
	Randomseed0   [8]byte
	Randomseed1   [8]byte
	Randomid      [8]byte
}

// NativeTun struct represents a TUN device
type NativeTun struct {
	name        string   // Name of the TUN device
	tunFile     *os.File  // File handle for the TUN device
	events      chan Event // Channel for TUN events
	errors      chan error // Channel for TUN errors
	routeSocket int      // File descriptor for the route socket
	closeOnce   sync.Once // Ensures Close() is called only once
}

// routineRouteListener listens for TUN events and sends them to the events channel
func (tun *NativeTun) routineRouteListener(tunIfindex int) {
	// ... (rest of the function remains unchanged)
}

// tunName returns the name of a TUN device given its file descriptor
func tunName(fd uintptr) (string, error) {
	// ... (rest of the function remains unchanged)
}

// tunDestroy destroys a TUN device with the given name
func tunDestroy(name string) error {
	// ... (rest of the function remains unchanged)
}

// CreateTUN creates a new TUN device with the given name and MTU
func CreateTUN(name string, mtu int) (Device, error) {
	// ... (rest of the function remains unchanged)
}

// CreateTUNFromFile creates a new TUN device from an existing file handle and MTU
func CreateTUNFromFile(file *os.File, mtu int) (Device, error) {
	// ... (rest of the function remains unchanged)
}

// Name returns the name of the TUN device
func (tun *NativeTun) Name() (string, error) {
	// ... (rest of the function remains unchanged)
}

// File returns the file handle for the TUN device
func (tun *NativeTun) File() *os.File {
	// ... (rest of the function remains unchanged)
}

// Events returns the channel for TUN events
func (tun *NativeTun) Events() <-chan Event {
	// ... (rest of the function remains unchanged)
}

// Read reads data from the TUN device and sends it to the given buffer
func (tun *NativeTun) Read(bufs [][]byte, sizes []int, offset int) (int, error) {
	// ... (rest of the function remains unchanged)
}

// Write writes data to the TUN device from the given buffer
func (tun *NativeTun) Write(bufs [][]byte, offset int) (int, error) {
	// ... (rest of the function remains unchanged)
}

// Close closes the TUN device and releases any associated resources
func (tun *NativeTun) Close() error {
	// ... (rest of the function remains unchanged)
}

// setMTU sets the MTU of the TUN device
func (tun *NativeTun) setMTU(n int) error {
	// ... (rest of the function remains unchanged)
}

// MTU returns the MTU of the TUN device
func (tun *NativeTun) MTU() (int, error) {
	// ... (rest of the function remains unchanged)
}

// BatchSize returns the batch size for the TUN device
func (tun *NativeTun) BatchSize() int {
	// ... (rest of the function remains unchanged)
}
