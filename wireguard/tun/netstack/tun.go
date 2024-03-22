/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package netstack

import (
	// ... (import statements)
)

// netTun represents a network TUN device.
type netTun struct {
	ep             *channel.Endpoint
	stack          *stack.Stack
	events         chan tun.Event
	incomingPacket chan *buffer.View
	mtu            int
	dnsServers     []netip.Addr
	hasV4, hasV6   bool // hasV4 and hasV6 indicate whether the device has IPv4 and IPv6 support, respectively.
}

// Net is a type alias for netTun.
type Net netTun

// CreateNetTUN creates a new netTun device with the given localAddresses and dnsServers.
func CreateNetTUN(localAddresses, dnsServers []netip.Addr, mtu int) (tun.Device, *Net, error) {
	// ... (function body)
}

// Name returns the name of the device.
func (tun *netTun) Name() (string, error) {
	// ... (function body)
}

// File returns an os.File for the device.
func (tun *netTun) File() *os.File {
	// ... (function body)
}

// Events returns a channel for receiving device events.
func (tun *netTun) Events() <-chan tun.Event {
	// ... (function body)
}

// Read reads packets from the device.
func (tun *netTun) Read(buf [][]byte, sizes []int, offset int) (int, error) {
	// ... (function body)
}

// Write writes packets to the device.
func (tun *netTun) Write(buf [][]byte, offset int) (int, error) {
	// ... (function body)
}

// WriteNotify notifies the device to read incoming packets.
func (tun *netTun) WriteNotify() {
	// ... (function body)
}

// Close closes the device.
func (tun *netTun) Close() error {
	// ... (function body)
}

// MTU returns the MTU of the device.
func (tun *netTun) MTU() (int, error) {
	// ... (function body)
}

// BatchSize returns the batch size of the device.
func (tun *netTun) BatchSize() int {
	// ... (function body)
}

// convertToFullAddr converts an netip.AddrPort to a tcpip.FullAddress and a tcpip.NetworkProtocolNumber.
func convertToFullAddr(endpoint netip.AddrPort) (tcpip.FullAddress, tcpip.NetworkProtocolNumber) {
	// ... (function body)
}

// DialContextTCPAddrPort dials a TCP connection to the given addr.
func (net *Net) DialContextTCPAddrPort(ctx context.Context, addr netip.AddrPort) (*gonet.TCPConn, error) {
	// ... (function body)
}

// DialContextTCP dials a TCP connection to the given *net.TCPAddr.
func (net *Net) DialContextTCP(ctx context.Context, addr *net.TCPAddr) (*gonet.TCPConn, error) {
	// ... (function body)
}

// DialTCPAddrPort dials a TCP connection to the given addr without context.
func (net *Net) DialTCPAddrPort(addr netip.AddrPort) (*gonet.TCPConn, error) {
	// ... (function body)
}

// DialTCP dials a TCP connection to the given *net.TCPAddr without context.
func (net *Net) DialTCP(addr *net.TCPAddr) (*gonet.TCPConn, error) {
	// ... (function body)
}

// ListenTCPAddrPort listens for incoming TCP connections on the given addr.
func (net *Net) ListenTCPAddrPort(addr netip.AddrPort) (*gonet.TCPListener, error) {
	// ... (function body)
}

// ListenTCP listens for incoming TCP connections on the given *net.TCPAddr.
func (net *Net) ListenTCP(addr *net.TCPAddr) (*gonet.TCPListener, error) {
	// ... (function body)
}

// DialUDPAddrPort dials a UDP connection to the given laddr and raddr.
func (net *Net) DialUDPAddrPort(laddr, raddr netip.AddrPort) (*gonet.UDPConn, error) {
	// ... (function body)
}

// ListenUDPAddrPort listens for incoming UDP connections on the given laddr.
func (net *Net) ListenUDPAddrPort(laddr netip.AddrPort) (*gonet.UDPConn, error) {
	// ... (function body)
}

// DialUDP dials a UDP connection to the given *net.UDPAddr.
func (net *Net) DialUDP(laddr, raddr *net.UDPAddr) (*gonet.UDPConn, error) {
	// ... (function body)
}

// ListenUDP listens for incoming UDP connections on the given *net.UDPAddr.
func (net *Net) ListenUDP(laddr *net.UDPAddr) (*gonet.UDPConn, error) {
	// ... (function body)
}

// ... (other function and type definitions)

