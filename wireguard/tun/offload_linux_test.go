package tun

import (
	"net/netip"
	"testing"

	"github.com/bepass-org/warp-plus/wireguard/conn"
	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/header"
)

// SPDX-License-Identifier: MIT
// Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.

const (
	offset = virtioNetHdrLen // Offset constant for the packet data.
)

