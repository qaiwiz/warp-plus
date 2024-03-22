package main

import (
	"context"
	"log/slog"
	"net"
	"net/netip"
	"os"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner"
	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
	"github.com/bepass-org/warp-plus/warp"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// privKey and pubKey are the private and public keys for Warp.
var (
	privKey           = "yGXeX7gMyUIZmK5QIgC7+XX5USUSskQvBYiQ6LdkiXI="
	pubKey            = "bmXOC+F1FxEMF9dyiK2H5/1SUtzH0JuVo51h2wPfgyo="
	googlev6DNSAddr80 = netip.MustParseAddrPort("[2001:4860:4860::8888]:80")
)

// canConnectIPv6 checks if a remote IPv6 address is reachable.
// It takes a netip.AddrPort as an argument and returns a boolean value.
func canConnectIPv6(remoteAddr netip.AddrPort) bool {
	// Create a dialer with a timeout of 5 seconds.
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}

	// Dial the remote address and return true if successful.
	conn, err := dialer.Dial("tcp6", remoteAddr.String())
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

// RunScan scans for available IP addresses using Warp.
// It takes two strings as arguments, privKey and pubKey, and returns a slice of statute.IPInfo.
func RunScan(privKey, pubKey string) (result []statute.IPInfo) {
	// Create a new scanner with various configurations.
	scanner := ipscanner.NewScanner(
		ipscanner.WithLogger(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))),
		ipscanner.WithWarpPing(),
		ipscanner.WithWarpPrivateKey(privKey),
		ipscanner.WithWarpPeerPublicKey(pubKey),
		ipscanner.WithUseIPv6(canConnectIPv6(googlev6DNSAddr80)),
		ipscanner.WithUseIPv4(true),
		ipscanner.WithMaxDesirableRTT(500*time.Millisecond),
		ipscanner.WithCidrList(warp.WarpPrefixes()),
	)

	// Create a context with a timeout of 2 minutes.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Start the scanner and wait for the results.
	scanner.Run(ctx)

	// Set up a ticker to check for results every second.
	t := time.NewTicker(1 * time.Second)
	
