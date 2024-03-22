// checksumNoFold calculates the IP checksum of a given byte slice without folding.
// It takes the byte slice 'b' and an initial 64-bit accumulator value 'initial' as input.
// The function processes the byte slice in chunks of 32-bit unsigned integers in big-endian byte order.
// It returns the calculated 64-bit checksum value.
func checksumNoFold(b []byte, initial uint64) uint64 {
	ac := initial // Initialize the accumulator with the given initial value

	// Process the byte slice in chunks of 128 bytes
	for len(b) >= 128 {
		ac += uint64(binary.BigEndian.Uint32(b[:4])) // Add the first 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[4:8])) // Add the second 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[8:12])) // Add the third 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[12:16])) // Add the fourth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[16:20])) // Add the fifth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[20:24])) // Add the sixth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[24:28])) // Add the seventh 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[28:32])) // Add the eighth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[32:36])) // Add the ninth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[36:40])) // Add the tenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[40:44])) // Add the eleventh 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[44:48])) // Add the twelfth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[48:52])) // Add the thirteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[52:56])) // Add the fourteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[56:60])) // Add the fifteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[60:64])) // Add the sixteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[64:68])) // Add the seventeenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[68:72])) // Add the eighteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[72:76])) // Add the nineteenth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[76:80])) // Add the twentieth 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[80:84])) // Add the twenty-first 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(b[84:88])) // Add the twenty-second 32-bit unsigned integer in big-endian byte order
		ac += uint64(binary.BigEndian.Uint32(
