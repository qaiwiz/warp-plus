package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/bepass-org/warp-plus/proxy/pkg/mixed" // Import the mixed package for creating the proxy
	"github.com/bepass-org/warp-plus/proxy/pkg/statute" // Import the statute package for ProxyRequest type
)

// The entry point of the application
func main() {
	// Create a new proxy with custom configurations
	proxy := mixed.NewProxy(
		mixed.WithBindAddress("127.0.0.1:1080"), // Set the bind address of the proxy
		mixed.WithUserHandler(generalHandler), // Set the user handler function
	)
	// Start the proxy and listen for incoming connections
	_ = proxy.ListenAndServe()
}

// generalHandler is the user handler function that processes incoming requests
func generalHandler(req *statute.ProxyRequest) error {
	// Print the destination of the request
	fmt.Println("handling request to", req.Destination)
	// Dial the destination using the provided network
	conn, err := net.Dial(req.Network, req.Destination)
	if err != nil {
		return err // Return the error if any
	}
	// Copy the request data to the destination connection
	go func() {
		_, err := io.Copy(conn, req.Conn)
		if err != nil {
			log.Println(err) // Log the error if any
		}
	}()
	// Copy the response data from the destination connection to the client
	_, err = io.Copy(req.Conn, conn)
	return err // Return the error if any
}

