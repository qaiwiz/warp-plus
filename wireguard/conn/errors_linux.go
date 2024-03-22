// SPDX-License-Identifier: MIT
// Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.

package conn

import (
	"errors"
	"os"

	"golang.org/x/sys/unix" // For using unix syscall constants
)

// Function: errShouldDisableUDPGSO(err error) bool
// Purpose: Check if the given error indicates that UDP GSO should be disabled.
//
// This function takes an error as an input and returns a boolean value indicating
// whether the error suggests that UDP GSO should be disabled.
func errShouldDisableUDPGSO(err error) bool {
	// Declare a new variable 'serr' of type '*os.SyscallError' and assign the
	// result of 'errors.As(err, &serr)'. This checks if the error 'err' can be
	// type-asserted as '*os.SyscallError' and stores the result in 'serr'.
	var serr *os.SyscallError
	if errors.As(err, &serr) {
		// If 'errors.As(err, &serr)' returns true, it means that 'err' is of type
		// '*os.SyscallError'. Now, we check if the error code is 'unix.EIO'.
		return serr.Err == unix.EIO
	}
	// If 'errors.As(err, &serr)' returns false or 'serr.Err' is not equal to
	// 'unix.EIO', we return false, indicating that UDP GSO should not be disabled.
	return false
}
