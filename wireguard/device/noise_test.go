/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/bepass-org/warp-plus/wireguard/conn"
	"github.com/bepass-org/warp-plus/wireguard/tun/tuntest"
)

// TestCurveWrappers tests the Curve wrappers by creating two private keys,
// deriving their public keys, and checking if shared secrets are equal.
func TestCurveWrappers(t *testing.T) {
	sk1, err := newPrivateKey()
	assertNil(t, err) // asserNil checks if err is nil, if not it will fatalf with err

	sk2, err := newPrivateKey()
	assertNil(t, err)

	pk1 := sk1.publicKey()
	pk2 := sk2.publicKey()

	ss1, err1 := sk1.sharedSecret(pk2)
	ss2, err2 := sk2.sharedSecret(pk1)

	// Check if shared secrets are equal and if there were no errors during
	// computation.
	if ss1 != ss2 || err1 != nil || err2 != nil {
		t.Fatal("Failed to compute shared secret")
	}
}

// randDevice creates a new Device with a random private key, a tuntest.TUN,
// and a NewLogger. It then sets the private key for the Device.
func randDevice(t *testing.T) *Device {
	sk, err := newPrivateKey()
	if err != nil {
		t.Fatal(err) // fatalf stops the test and prints err
	}
	tun := tuntest.NewChannelTUN()
	logger := NewLogger(LogLevelError, "")
	device := NewDevice(tun.TUN(), conn.NewDefaultBind(), logger)
	device.SetPrivateKey(sk)
	return device
}

// assertNil checks if an error is nil, if not it will fatalf with the error.
func assertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// assertEqual checks if two byte slices are equal, if not it will fatalf with
// a message containing the two byte slices.
func assertEqual(t *testing.T, a, b []byte) {
	if !bytes.Equal(a, b) {
		t.Fatal(a, "!=", b)
	}
}

// TestNoiseHandshake tests the Noise handshake by creating two Devices,
// performing the handshake, and checking if keys are derived correctly.
func TestNoiseHandshake(t *testing.T) {
	dev1 := randDevice(t)
	dev2 := randDevice(t)

	defer dev1.Close()
	defer dev2.Close()

	// Create a new peer for each Device with the other Device's public key.
	peer1, err := dev2.NewPeer(dev1.staticIdentity.privateKey.publicKey())
	if err != nil {
		t.Fatal(err)
	}
	peer2, err := dev1.NewPeer(dev2.staticIdentity.privateKey.publicKey())
	if err != nil {
		t.Fatal(err)
	}

	// Start the handshake for both peers.
	peer1.Start()
	peer2.Start()

	// Check if precomputedStaticStatic keys are equal.
	assertEqual(
		t,
		peer1.handshake.precomputedStaticStatic[:],
		peer2.handshake.precomputedStaticStatic[:],
	)

	// Simulate the handshake by exchanging initiation and response messages.

	// Exchange initiation message.
	t.Log("exchange initiation message")
	msg1, err := dev1.CreateMessageInitiation(peer2)
	assertNil(t, err)

	packet := make([]byte, 0, 256)
	writer := bytes.NewBuffer(packet)
	err = binary.Write(writer, binary.LittleEndian, msg1)
	assertNil(t, err)

	// Consume the initiation message by dev2 and check if a peer was created.
	peer := dev2.ConsumeMessageInitiation(msg1)
	if peer == nil {
		t.Fatal("handshake failed at initiation message")
	}

	// Check if chainKey and hash are equal for both peers.
	assertEqual(
		t,
		peer1.handshake.chainKey[:],
		peer2.handshake.chainKey[:],
	)

	assertEqual(
		t,
		peer1.handshake.hash[:],
		peer2.handshake.hash[:],
	)

	// Exchange response message.
	t.Log("exchange response message")
	msg2, err := dev2.CreateMessageResponse(peer1)
	assertNil(t, err)

	// Consume the response message by dev1.
	peer = dev1.ConsumeMessageResponse(msg2)
	if peer == nil {
		t.Fatal("handshake failed at response message")
	}

	// Check if chainKey and hash are equal for both peers.
	assertEqual(
		t,
		peer1.handshake.chainKey[:],
		peer2.handshake.chainKey[:],
	)

	assertEqual(
		t,
		peer1.handshake.hash[:],
		peer2.handshake.hash[:],
	)

	// Derive keys for both peers
