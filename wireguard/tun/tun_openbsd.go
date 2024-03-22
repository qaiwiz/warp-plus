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

	"golang.org/x/sys/unix" // Importing unix package from golang.org/x/sys for system calls
)

// Structure for iface mtu get/set ioctls
type ifreq_mtu struct { // Define a new type 'ifreq_mtu' for iface mtu get/set ioctls
	Name [unix.IFNAMSIZ]byte // Name field to store interface name
	MTU  uint32              // MTU field to store MTU value
	Pad0 [12]byte            // Padding bytes to ensure proper alignment
}

const _TUNSIFMODE = 0x8004745d // Const for _TUNSIFMODE

type NativeTun struct { // Define a new type 'NativeTun' for native tun device
	name        string        // name field to store tun device name
	tunFile     *os.File      // tunFile field to store tun device file
	events      chan Event    // events channel to send events
	errors      chan error    // errors channel to send errors
	routeSocket int           // routeSocket field to store route socket file descriptor
	closeOnce   sync.Once      // closeOnce field to ensure only one close operation
}

// ... (rest of the code)

// Function to create a new TUN device with the given name and mtu
func CreateTUN(name string, mtu int) (Device, error) {
	// ... (code)
}

// Function to create a new TUN device from an existing file
func CreateTUNFromFile(file *os.File, mtu int) (Device, error) {
	// ... (code)
}

// Function to get the name of the tun device
func (tun *NativeTun) Name() (string, error) {
	// ... (code)
}

// Function to get the file descriptor of the tun device
func (tun *NativeTun) File() *os.File {
	// ... (code)
}

// Function to get the events channel of the tun device
func (tun *NativeTun) Events() <-chan Event {
	// ... (code)
}

// Function to read data from the tun device
func (tun *NativeTun) Read(bufs [][]byte, sizes []int, offset int) (int, error) {
	// ... (code)
}

// Function to write data to the tun device
func (tun *NativeTun) Write(bufs [][]byte, offset int) (int, error) {
	// ... (code)
}

// Function to close the tun device
func (tun *NativeTun) Close() error {
	// ... (code)
}

// Function to set the MTU of the tun device
func (tun *NativeTun) setMTU(n int) error {
	// ... (code)
}

// Function to get the MTU of the tun device
func (tun *NativeTun) MTU() (int, error) {
	// ... (code)
}

// Function to get the batch size of the tun device
func (tun *NativeTun) BatchSize() int {
	// ... (code)
}
