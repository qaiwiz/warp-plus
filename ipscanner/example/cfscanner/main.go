package main

import (
	"context"
	"github.com/bepass-org/warp-plus/ipscanner" // Import the ipscanner package
)

func main() {
	// Create a new IP scanner using the ipscanner.NewScanner function
	scanner := ipscanner.NewScanner(
		ipscanner.WithHTTPPing(), // Enable HTTP ping (ICMP echo)
		ipscanner.WithUseIPv6(true), // Enable IPv6 support
	)

	// Create a new context with cancellation and defer cancellation when the function returns
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the scanner as a goroutine
	go scanner.Run(ctx)

	// Wait for the context to be done (either due to cancellation or an error)
	<-ctx.Done()
}
