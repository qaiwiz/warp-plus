//go:build linux && !android

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn // Package for connection-related functionality

import (
	"context"
	"net"
	"net/netip" // For IP address manipulation
	"runtime"   // For runtime-specific operations
	"testing"  // For testing functionality
	"unsafe"    // For low-level memory manipulation

	"golang.org/x/sys/unix" // For system calls and constants
)

