package tun

import "encoding/binary"

// checksumNoFold calculates the IP checksum of a given byte slice without folding.
// It takes the byte slice 'b' and an initial 64-bit accumulator value 'initial' as input.
// The function processes the byte slice in chunks of 32-bit unsigned integers in big-endian byte order.
// It returns the calculated 64-bit checksum value.
func checksumNoFold(b []byte, initial uint64) uint64 {
	ac := initial

	// Process the byte slice in chunks of 128 bytes.
	for len(b) >= 128 {
		ac += uint64(binary.BigEndian.Uint32(b[:4]))
		ac += uint64(binary.BigEndian.Uint32(b[4:8]))
		ac += uint64(binary.BigEndian.Uint32(b[8:12]))
		ac += uint64(binary.BigEndian.Uint32(b[12:16]))
		ac += uint64(binary.BigEndian.Uint32(b[16:20]))
		ac += uint64(binary.BigEndian.Uint32(b[20:24]))
		ac += uint64(binary.BigEndian.Uint32(b[24:28]))
		ac += uint64(binary.BigEndian.Uint32(b[28:32]))
		ac += uint64(binary.BigEndian.Uint32(b[32:36]))
		ac += uint64(binary.BigEndian.Uint32(b[36:40]))
		ac += uint64(binary.BigEndian.Uint32(b[40:44]))
		ac += uint64(binary.BigEndian.Uint32(b[44:48]))
		ac += uint64(binary.BigEndian.Uint32(b[48:52]))
		ac += uint64(binary.BigEndian.Uint32(b[52:56]))
		ac += uint64(binary.BigEndian.Uint32(b[56:60]))
		ac += uint64(binary.BigEndian.Uint32(b[60:64]))
		ac += uint64(binary.BigEndian.Uint32(b[64:68]))
		ac += uint64(binary.BigEndian.Uint32(b[68:72]))
		ac += uint64(binary.BigEndian.Uint32(b[72:76]))
		ac += uint64(binary.BigEndian.Uint32(b[76:80]))
		ac += uint64(binary.BigEndian.Uint32(b[80:84]))
		ac += uint64(binary.BigEndian.Uint32(b[84:88]))
		ac += uint64(binary.BigEndian.Uint32(b[88:92]))
		ac += uint64(binary.BigEndian.Uint32(b[92:96]))
		ac += uint64(binary.BigEndian.Uint32(b[96:100]))
		ac += uint64(binary.BigEndian.Uint32(b[100:104]))
		ac += uint64(binary.BigEndian.Uint32(b[104:108]))
		ac += uint64(binary.BigEndian.Uint32(b[108:112]))
		ac += uint64(binary.BigEndian.Uint32(b[112:116]))
		ac += uint64(binary.BigEndian.Uint32(b[116:120]))
		ac += uint64(binary.BigEndian.Uint32(b[120:124]))
		ac += uint64(binary.BigEndian.Uint32(b[124:128]))
		b = b[128:]
	}

	// Process the remaining bytes in smaller chunks.
	if len(b) >= 64 {
		ac += uint64(binary.BigEndian.Uint32(b[:4]))
		ac += uint64(binary.BigEndian.Uint32(b[4:8]))
		ac += uint64(binary.BigEndian.Uint32(b[8:12]))
		ac += uint64(binary.BigEndian.Uint32(b[12:16]))
		ac += uint64(binary.BigEndian.Uint32(b[16:20]))
		ac += uint64(binary.BigEndian.Uint32(b[20:24]))
		ac += uint64(binary.BigEndian.Uint32(b[24:28]))
		ac += uint64(binary.BigEndian.Uint32(b[28:32]))
		ac += uint64(binary.BigEndian.Uint32(b[32:36]))
		ac += uint64(binary.BigEndian.Uint32(b[36:40]))
		ac += uint64(binary.BigEndian.Uint32(b[40:44]))
		ac += uint64(binary.BigEndian.Uint32(b[44:48]))
		ac += uint64(binary.BigEndian.Uint32(b[48:52]))
		ac += uint64(binary.BigEndian.Uint32(b[52:56]))
		ac += uint64(binary.BigEndian.Uint32(b[56:60]))
		ac += uint64(binary.BigEndian.
