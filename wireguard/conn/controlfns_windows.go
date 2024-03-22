/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package conn

import (
	"syscall"

	"golang.org/x/sys/windows"
)

// init is a package-level initialization function that configures socket buffer sizes
// for Windows platforms by appending a new control function to the controlFns slice.
// This function sets both receive and send buffer sizes to the value of socketBufferSize.
func init() {
	// Append a new control function to the controlFns slice.
	controlFns = append(controlFns,
		func(network, address string, c syscall.RawConn) error {
			// The control function takes a file descriptor (fd) and sets the receive and send buffer sizes
			// using the Windows-specific SetsockoptInt function.
			return c.Control(func(fd uintptr) {
				// Set the receive buffer size to socketBufferSize.
				_ = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_RCVBUF, socketBufferSize)
				// Set the send buffer size to socketBufferSize.
				_ = windows.SetsockoptInt(windows.Handle(fd), windows.SOL_SOCKET, windows.SO_SNDBUF, socketBufferSize)
			})
		},
	)
}

