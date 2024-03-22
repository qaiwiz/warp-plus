//go:build windows || wasm

// SPDX-License-Identifier: MIT

package rwcancel

// RWCancel struct represents a type that can be used to cancel read-write operations.
type RWCancel struct{}

// Cancel method cancels any ongoing read-write operations.
// This method does not return any value.
func (*RWCancel) Cancel() {}

