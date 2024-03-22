/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package device

import (
	"crypto/subtle"
	"encoding/hex"
	"errors"
)

// NoisePublicKeySize is the size of a NoisePublicKey, which is 32 bytes.
const NoisePublicKeySize = 32

// NoisePrivateKeySize is the size of a NoisePrivateKey, which is 32 bytes.
const NoisePrivateKeySize = 32

// NoisePresharedKeySize is the size of a NoisePresharedKey, which is 32 bytes.
const NoisePresharedKeySize = 32

// NoiseNonce is a type representing a Noise nonce, which is a 64-bit value padded to 12 bytes.
type NoiseNonce uint6
