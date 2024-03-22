/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	"encoding/hex"
	"testing"

	"golang.org/x/crypto/blake2s" // Importing the Blake2s hash function from the x/crypto module.
)

// KDFTest struct defines the test cases for Key Derivation Function (KDF) testing.
type KDFTest struct {
	key     string // The input key for the KDF.
	input   string // The input data for the KDF.
	t0, t1, t2 string // The expected output values for the KDF.
}

// assertEquals function checks if two strings are equal and fails the test if they are not.
func assertEquals(t *testing.T, a, b string) {
	if a != b {
		t.Fatal("expected", a, "=", b)
	}
}

// TestKDF function tests the Key Derivation Function (KDF) with different input keys and input data.
func TestKDF(t *testing.T) {
	tests := []KDFTest{
		// Test case 1:
		{
			key:   "746573742d6b6579", // Input key: "test-key"
			input: "746573742d696e707574", // Input data: "test-input"
			t0:    "6f0e5ad38daba1bea8a0d213688736f19763239305e0f58aba697f9ffc41c633", // Expected output t0
			t1:    "df1194df20802a4fe594cde27e92991c8cae66c366e8106aaa937a55fa371e8a", // Expected output t1
			t2:    "fac6e2745a325f5dc5d11a5b165aad08b0ada28e7b4e666b7c077934a4d76c24", // Expected output t2
		},
		// Test case 2:
		{
			key:   "776972656775617264", // Input key: "wireguard"
			input: "776972656775617264", // Input data: "wireguard"
			t0:    "491d43bbfdaa8750aaf535e334ecbfe5129967cd64635101c566d4caefda96e8", // Expected output t0
			t1:    "1e71a379baefd8a79aa4662212fcafe19a23e2b609a3db7d6bcba8f560e3d25f", // Expected output t1
			t2:    "31e1ae48bddfbe5de38f295e5452b1909a1b4e3
