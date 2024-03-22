package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"strconv"
)

// Main function is the entry point of the program
func main() {
	proxyAddr := "127.0.0.1:1080" // SOCKS5 proxy address
	targetAddr := "<YOUR Netcat ip address>:4444" // Target address to connect

	// Connect to the SOCKS5 proxy
	conn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		panic(err) // Panic if there's an error during connection
	}
	defer conn.Close() // Close the connection when the function ends

	// Send greeting to the SOCKS5 proxy
	conn.Write([]byte{0x05, 0x01, 0x00})

	// Read greeting response
	response := make([]byte, 2)
	io.ReadFull(conn, response)

	// Send UDP ASSOCIATE request
	conn.Write([]byte{0x05, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	// Read UDP ASSOCIATE response
	response = make([]byte, 10)
	io.ReadFull(conn, response)

	// Extract the bind address and port
	bindIP := net.IP(response[4:8]) // IP address
	bindPort := binary.BigEndian.Uint16(response[8:10]) // Port number

	// Print the bind address
	fmt.Printf("Bind address: %s:%d\n", bindIP, bindPort)

	// Create UDP connection
	udpConn, err := net.Dial("udp", fmt.Sprintf("%s:%d", bindIP, bindPort))
	if err != nil {
		panic(err) // Panic if there's an error during connection
	}
	defer udpConn.Close() // Close the connection when the function ends

	// Extract target IP and port
	dstIP, dstPortStr, _ := net.SplitHostPort(targetAddr) // Split the target address into IP and port
	dstPort, _ := strconv.Atoi(dstPortStr) // Convert the port string to an integer

	// Construct the UDP packet with the target address and message
	packet := make([]byte, 0)
	packet = append(packet, 0x00, 0x00, 0x00) // RSV and FRAG
	packet = append(packet, 0
