//go:build !windows

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

// NewDefaultBind creates a new default bind object.
// It returns an instance of Bind, specifically a StdNetBind.
// This function is only available on non-Windows systems.
func NewDefaultBind() Bind { return NewStdNetBind() }
