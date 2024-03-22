/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package tuntest

import (
	"encoding/binary"
	"io"
	"net/netip"
	"os"

	"github.com/bepass-org/warp-plus/wireguard/tun" // Importing tun package for tun.Event and tun.Device
)

// Ping generates an ICMPv4 echo request packet for the specified destination and source addresses.
func Ping(dst, src netip.Addr) []byte {
	localPort := uint16(1337)
	seq := uint16(0)

	payload := make([]byte, 4)
	binary.BigEndian.PutUint16(payload[0:], localPort)
	binary.BigEndian.PutUint16(payload[2:], seq)

	return genICMPv4(payload, dst, src)
}

// checksum calculates the "internet checksum" as per RFC1071.
func checksum(buf []byte, initial uint16) uint16 {
	v := uint32(initial)
	for i := 0; i < len(buf)-1; i += 2 {
		v += uint32(binary.BigEndian.Uint16(buf[i:]))
	}
	if len(buf)%2 == 1 {
		v += uint32(buf[len(buf)-1]) << 8
	}
	for v > 0xffff {
		v = (v >> 16) + (v & 0xffff)
	}
	return ^uint16(v)
}

// genICMPv4 generates an ICMPv4 echo request packet with the given payload, destination, and source addresses.
func genICMPv4(payload []byte, dst, src netip.Addr) []byte {
	// Constants for ICMPv4 and IPv4 header fields.
	const (
		icmpv4ProtocolNumber = 1
		icmpv4Echo           = 8
		icmpv4ChecksumOffset = 2
		icmpv4Size           = 8
		ipv4Size             = 20
		ipv4TotalLenOffset   = 2
		ipv4ChecksumOffset   = 10
		ttl                  = 65
		headerSize           = ipv4Size + icmpv4Size
	)

	pkt := make([]byte, headerSize+len(payload))

	ip := pkt[0:ipv4Size]
	icmpv4 := pkt[ipv4Size : ipv4Size+icmpv4Size]

	// Encoding the ICMPv4 header fields.
	icmpv4[0] = icmpv4Echo // type
	icmpv4[1] = 0          // code
	chksum := ^checksum(icmpv4, checksum(payload, 0))
	binary.BigEndian.PutUint16(icmpv4[icmpv4ChecksumOffset:], chksum)

	// Encoding the IPv4 header fields.
	length := uint16(len(pkt))
	ip[0] = (4 << 4) | (ipv4Size / 4)
	binary.BigEndian.PutUint16(ip[ipv4TotalLenOffset:], length)
	ip[8] = ttl
	ip[9] = icmpv4ProtocolNumber
	copy(ip[12:], src.AsSlice())
	copy(ip[16:], dst.AsSlice())
	chksum = ^checksum(ip[:], 0)
	binary.BigEndian.PutUint16(ip[ipv4ChecksumOffset:], chksum)

	// Copying the payload to the packet.
	copy(pkt[headerSize:], payload)
	return pkt
}

// ChannelTUN represents a channel-based TUN device.
type ChannelTUN struct {
	Inbound  chan []byte // incoming packets, closed on TUN close
	Outbound chan []byte // outbound packets, blocks forever on TUN close

	closed chan struct{}
	events chan tun.Event
	tun    chTun
}

// NewChannelTUN creates a new ChannelTUN instance.
func NewChannelTUN() *ChannelTUN {
	c := &ChannelTUN{
		Inbound:  make(chan []byte),
		Outbound: make(chan []byte),
		closed:   make(chan struct{}),
		events:   make(chan tun.Event, 1),
	}
	c.tun.c = c
	c.events <- tun.EventUp
	return c
}

// TUN returns the tun.Device interface for the ChannelTUN.
func (c *ChannelTUN) TUN() tun.Device {
	return &c.tun
}

// chTun is the internal struct for ChannelTUN.tun.
type chTun struct {
	c *ChannelTUN
}

// File returns nil for chTun.
func (t *chTun) File() *os.File { return nil }

// Read reads packets from the TUN device.
func (t *chTun) Read(packets [][]byte, sizes []int, offset int) (int, error) {
	select {
	case <-t.c.closed:
		return 0, os.ErrClosed
	case msg := <-t.c.Outbound:
		n := copy(packets[0][offset:], msg)
		sizes[0] = n
		return 1, nil
	}
}

// Write is called by the wireguard device to deliver a packet for routing.
func (t *chTun
