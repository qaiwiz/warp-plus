/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package bindtest

import (
	"fmt"
	"math/rand"
	"net"
	"net/netip"
	"os"

	"github.com/bepass-org/warp-plus/wireguard/conn" // Importing conn package for implementing Bind and Endpoint interfaces
)

// ChannelBind struct represents a channel bind object with various channels and endpoints
type ChannelBind struct {
	rx4, tx4         *chan []byte // Channels for receiving and transmitting IPv4 packets
	rx6, tx6         *chan []byte // Channels for receiving and transmitting IPv6 packets
	closeSignal      chan bool   // Channel for signaling closure
	source4, source6 ChannelEndpoint
	target4, target6 ChannelEndpoint
}

// ChannelEndpoint type represents a channel endpoint
type ChannelEndpoint uint16

// Interfaces are implemented for ChannelBind and ChannelEndpoint types
var (
	_ conn.Bind     = (*ChannelBind)(nil) // ChannelBind implements Bind interface
	_ conn.Endpoint = (*ChannelEndpoint)(nil) // ChannelEndpoint implements Endpoint interface
)

// NewChannelBinds function returns a slice of two ChannelBind objects
func NewChannelBinds() [2]conn.Bind {
	// Creating channels for receiving and transmitting packets
	arx4 := make(chan []byte, 8192)
	brx4 := make(chan []byte, 8192)
	arx6 := make(chan []byte, 8192)
	brx6 := make(chan []byte, 8192)

	// Initializing ChannelBind objects and assigning channels and endpoints
	var binds [2]ChannelBind
	binds[0].rx4 = &arx4
	binds[0].tx4 = &brx4
	binds[1].rx4 = &brx4
	binds[1].tx4 = &arx4
	binds[0].rx6 = &arx6
	binds[0].tx6 = &brx6
	binds[1].rx6 = &brx6
	binds[1].tx6 = &arx6
	binds[0].target4 = ChannelEndpoint(1)
	binds[1].target4 = ChannelEndpoint(2)
	binds[0].target6 = ChannelEndpoint(3)
	binds[1].target6 = ChannelEndpoint(4)
	binds[0].source4 = binds[1].target4
	binds[0].source6 = binds[1].target6
	binds[1].source4 = binds[0].target4
	binds[1].source6 = binds[0].target6

	// Returning ChannelBind objects
	return [2]conn.Bind{&binds[0], &binds[1]}
}

// ClearSrc method for ChannelEndpoint type
func (c ChannelEndpoint) ClearSrc() {}

// SrcToString method for ChannelEndpoint type
func (c ChannelEndpoint) SrcToString() string { return "" }

// DstToString method for ChannelEndpoint type
func (c ChannelEndpoint) DstToString() string {
	return fmt.Sprintf("127.0.0.1:%d", c) // Returns the endpoint as a string in the format of "127.0.0.1:<port>"
}

// DstToBytes method for ChannelEndpoint type
func (c ChannelEndpoint) DstToBytes() []byte {
	return []byte{byte(c)} // Returns the endpoint as a byte slice
}

// DstIP method for ChannelEndpoint type
func (c ChannelEndpoint) DstIP() netip.Addr {
	return netip.AddrFrom4([4]byte{127, 0, 0, 1}) // Returns the endpoint's IP address as a netip.Addr object
}

// SrcIP method for ChannelEndpoint type
func (c ChannelEndpoint) SrcIP() netip.Addr {
	return netip.Addr{} // Returns an empty netip.Addr object for the source IP address
}

// Open method for ChannelBind type
func (c *ChannelBind) Open(port uint16) (fns []conn.ReceiveFunc, actualPort uint16, err error) {
	// Initializing closeSignal channel
	c.closeSignal = make(chan bool)

	// Appending receive functions to the fns slice
	fns = append(fns, c.makeReceiveFunc(*c.rx4))
	fns = append(fns, c.makeReceiveFunc(*c.rx6))

	// Generating a random number to determine the actual port
	if rand.Uint32()&1 == 0 {
		return fns, uint16(c.source4), nil
	} else {
		return fns, uint16(c.source6), nil
	}
}

// Close method for ChannelBind type
func (c *ChannelBind) Close() error {
	// Checking if closeSignal channel is initialized
	if c.closeSignal != nil {
		// Closing the closeSignal channel
		select {
		case <-c.closeSignal:
		default:
			close(c.closeSignal)
	
