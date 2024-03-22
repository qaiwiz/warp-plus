/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */
package conn

// PeekLookAtSocketFd4 returns the file descriptor (fd) of the IPv4 socket connection.
// It first gets a SyscallConn from the IPv4 connection and then uses the Control function
// to get the file descriptor. If there is an error, it returns -1 and the error.
func (s *StdNetBind) PeekLookAtSocketFd4() (fd int, err error) {
	sysconn, err := s.ipv4.SyscallConn()
	if err != nil {
		return -1, err
	}
	err = sysconn.Control(func(f uintptr) {
		fd = int(f)
	})
	if err != nil {
	
