//go:build linux || openbsd || freebsd

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn // The package name is "conn".

import (
	"runtime"

	"golang.org/x/sys/unix" // The "unix" package is imported for system calls.
)

// fwmarkIoctl is a variable that stores the value of the IOCTL command.
var fwmarkIoctl int

// init is a special function that is called automatically when the package is imported.
// It sets the value of fwmarkIoctl based on the operating system.
func init() {
	switch runtime.GOOS {
	case "linux", "android":
		fwmarkIoctl = 36 /* unix.SO_MARK */
	case "freebsd":
		fwmarkIoctl = 0x1015 /* unix.SO_USER_COOKIE */
	case "openbsd":
		fwmarkIoctl = 0x1021 /* unix.SO_RTABLE */
	}
}

// The "StdNetBind" type is defined here, but its definition is not provided.

// SetMark is a method on the "StdNetBind" type that sets the firewall mark on the socket.
func (s *StdNetBind) SetMark(mark uint32) error {
	var operr error

	// If fwmarkIoctl is 0, then the operating system does not support setting the firewall mark,
	// so the method returns immediately.
	if fwmarkIoctl == 0 {
		return nil
	}

	// The method iterates over the IPv4 and IPv6 sockets and sets the firewall mark on each one.
	if s.ipv4 != nil {
		fd, err := s.ipv4.SyscallConn()
		if err != nil {
			return err
		}
		err = fd.Control(
