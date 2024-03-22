//go:build darwin || freebsd

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package tun

import (
	"fmt"
	"syscall" // sysconn is used to make system calls
)

// operateOnFd is a method on NativeTun struct that takes a function (fn) as an argument
// and applies it to the file descriptor (fd) of the tunFile field.
func (tun *NativeTun) operateOnFd(fn func(fd uintptr)) {
	// Create a syscallConn from tunFile
	sysconn, err := tun.tunFile.SyscallConn()
	if err != nil {
		// If there is an error, send it to the errors channel
	
