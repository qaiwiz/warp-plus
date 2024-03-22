// This code block is only built for non-Windows, non-Linux, and non-WASM platforms.
//go:build !windows && !linux && !wasm

/*
 SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn // The package name is 'conn'.

import (
	"syscall" // The syscall package is imported for accessing lower-level system calls.

	"golang.org/x/sys/unix" // The unix package from the golang.org/x/sys/unix repository is imported.
)

// init is the initializer function for the package. It is called automatically when the package is imported.
func init() {
	controlFns = append(controlFns, // The controlFns slice is appended with new functions.

		// This function sets the receive and send buffer sizes for the socket.
		func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_RCVBUF, socketBufferSize) // Set the receive buffer size.
				_ = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_SNDBUF, socketBufferSize) // Set the send buffer size.
			})
		},

		// This function sets the IPV6_V6ONLY socket option for UDP version 6 (udp6) networks.
		func(network, address string, c syscall.RawConn) error {
			var err error
			if network == "udp6" {
				c.Control(func(fd uintptr) {
					err = unix.SetsockoptInt(int(fd), unix.IPPROTO_IPV6, unix.IPV
