//go:build darwin || freebsd || openbsd

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package ipc

import (
	"errors"
	"net"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

// UAPIListener is a wrapper around net.Listener that provides additional
// functionality to watch for the deletion of the Unix domain socket.
type UAPIListener struct {
	listener net.Listener // The underlying net.Listener.
	connNew  chan net.Conn   // Channel to receive new connections.
	connErr  chan error      // Channel to receive errors.
	kqueueFd int            // File descriptor for the kqueue.
	keventFd int            // File descriptor for the socket directory.
}

// Accept implements the net.Listener interface and returns the next
// connection on the listener.
func (l *UAPIListener) Accept() (net.Conn, error) {
	// Wait for a new connection or an error on the connNew or connErr channels.
	for {
		select {
		case conn := <-l.connNew:
			// A new connection is available.
			return conn, nil

		case err := <-l.connErr:
			// An error occurred.
			return nil, err
		}
	}
}

// Close implements the net.Listener interface and closes the listener.
func (l *UAPIListener) Close() error {
	// Close the kqueue and socket directory file descriptors.
	err1 := unix.Close(l.kqueueFd)
	err2 := unix.Close(l.keventFd)

	// Close the underlying net.Listener.
	err3 := l.listener.Close()

	// Return the first error encountered.
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return err3
}

// Addr implements the net.Listener interface and returns the listener's
// address.
func (l *UAPIListener) Addr() net.Addr {
	// Return the address of the underlying net.Listener.
	return l.listener.Addr()
}

// UAPIListen creates a new net.Listener that listens on the Unix domain
// socket at the specified path. It also sets up a kqueue to monitor the
// socket directory for the deletion of the socket. If the socket is
// deleted, the kqueue will trigger an event, and the listener's
// connErr channel will receive an error indicating that the socket has
// been deleted.
func UAPIListen(name string, file *os.File) (net.Listener, error) {
	// Wrap the file in a net.Listener.
	listener, err := net.FileListener(file)
	if err != nil {
		return nil, err
	}

	// Create a new UAPIListener.
	uapi := &UAPIListener{
		listener: listener,
		connNew:  make(chan net.Conn, 1),
		connErr:  make(
