/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	"crypto/hmac"
	"crypto/rand"
	"sync"
	"time"

	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20poly1305"
)

// CookieChecker struct holds state for checking WireGuard cookies.
type CookieChecker struct {
	sync.RWMutex
	mac1 struct {
		key [blake2s.Size]byte // Holds the key for MAC1 computation.
	}
	mac2 struct {
		secret        [blake2s.Size]byte // Holds the secret for MAC2 computation.
		secretSet     time.Time          // Tracks when the secret was last set.
		encryptionKey [chacha20poly1305.KeySize]byte // Holds the encryption key for MAC2.
	}
}

// CookieGenerator struct holds state for generating WireGuard cookies.
type CookieGenerator struct {
	sync.RWMutex
	mac1 struct {
		key [blake2s.Size]byte // Holds the key for MAC1 computation.
	}
	mac2 struct {
		cookie        [blake2s.Size128]byte // Holds the cookie for MAC2 computation.
		cookieSet     time.Time          // Tracks when the cookie was last set.
		hasLastMAC1   bool              // Tracks if a valid MAC1 is available.
		lastMAC1      [blake2s.Size128]byte // Holds the last MAC1.
		encryptionKey [chacha20poly1305.KeySize]byte // Holds the encryption key for MAC2.
	}
}

// Init initializes the CookieChecker with a NoisePublicKey.
func (st *CookieChecker) Init(pk NoisePublicKey) {
	st.Lock()
	defer st.Unlock()

	// Initialize mac1 state
	func() {
		hash, _ := blake2s.New256(nil)
		hash.Write([]byte(WGLabelMAC1))
		hash.Write(pk[:])
		hash.Sum(st.mac1.key[:0])
	}()

	// Initialize mac2 state
	func() {
		hash, _ := blake2s.New256(nil)
		hash.Write([]byte(WGLabelCookie))
		hash.Write(pk[:])
		hash.Sum(st.mac2.encryptionKey[:0])
	}()

	st.mac2.secretSet = time.Time{}
}

// CheckMAC1 checks the MAC1 of a WireGuard message.
func (st *CookieChecker) CheckMAC1(msg []byte) bool {
	st.RLock()
	defer st.RUnlock()

	size := len(msg)
	smac2 := size - blake2s.Size128
	smac1 := smac2 - blake2s.Size128

	var mac1 [blake2s.Size128]byte

	mac, _ := blake2s.New128(st.mac1.key[:])
	mac.Write(msg[:smac1])
	mac.Sum(mac1[:0])

	return hmac.Equal(mac1[:], msg[smac1:smac2])
}

// CheckMAC2 checks the MAC2 of a WireGuard message.
func (st *CookieChecker) CheckMAC2(msg, src []byte) bool {
	st.RLock()
	defer st.RUnlock()

	if time.Since(st.mac2.secretSet) > CookieRefreshTime {
		return false
	}

	// Derive cookie key
	var cookie [blake2s.Size128]byte
	func() {
		mac, _ := blake2s.New128(st.mac2.secret[:])
		mac.Write(src)
		mac.Sum(cookie[:0])
	}()

	// Calculate MAC of packet (including MAC1)
	smac2 := len(msg) - blake2s.Size128

	var mac2 [blake2s.Size128]byte
	func() {
		mac, _ := blake2s.New128(cookie[:])
		mac.Write(msg[:smac2])
		mac.Sum(mac2[:0])
	}()

	return hmac.Equal(mac2[:], msg[smac2:])
}

// CreateReply creates a reply message with a new MAC2 for a WireGuard message.
func (st *CookieChecker) CreateReply(
	msg []byte,
	recv uint32,
	src []byte,
) (*MessageCookieReply, error) {
	st.RLock()

	// Refresh cookie secret
	if time.Since(st.mac2.secretSet) > CookieRefreshTime {
		st.RUnlock()
		st.Lock()
		_, err := rand.Read(st.mac2.secret[:])
		if err != nil {
			st.Unlock()
			return nil, err
		}
		st.mac2.secretSet = time.Now()
		st.Unlock()
		st.RLock()
	}

	// Derive cookie
	var cookie [blake2s.Size128]byte
	func() {
		mac, _ := blake2s.New128(st.mac2.secret[:])
		mac.Write(src)
		mac.Sum(cookie[:0])
	}()

	// Encrypt cookie
	size := len(msg)

	smac2 := size - blake2s.Size128
	smac1 := smac2 - blake2s.Size128

	reply := new(MessageCookieReply)
	
