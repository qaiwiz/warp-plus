// This Go file contains a single function, errShouldDisableUDPGSO,
// which takes an error as an input and returns a boolean value.
// The function is only built for non-Linux systems due to the build constraint: go:build !linux

// SPDX-License-Identifier: MIT
// This identifies the license for the code that follows, which is the MIT License.
// The copyright notice indicates that the code is owned by WireGuard LLC.

package conn // The package name is conn, which likely contains connection-related functions.

// The function errShouldDisableUDPGSO takes an error (err) as an input and returns
// a boolean value (bool). It always returns false in this implementation.
func errShouldDisableUDPGSO(err error) bool {
	return false
}

