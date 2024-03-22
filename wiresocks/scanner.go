package wiresocks

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner"
	"github.com/bepass-org/warp-plus/warp"
	"github.com/go-ini/ini"
)

// ScanOptions struct holds the options for the IP scan
type ScanOptions struct {
	V4     bool   // IPv4 scan enabled
	V6     bool   // IPv6 scan enabled
	MaxRTT time.Duration
	// MaxRTT is the maximum round-trip time for the scan
}

// RunScan function initiates an IP scan with the given options
func RunScan(ctx context.Context, l *slog.Logger, opts ScanOptions) (result []ipscanner.IPInfo, err error) {
	// Load the configuration file
	cfg, err := ini.Load("./primary/wgcf-profile.ini")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Read the private key from the 'Interface' section
	privateKey := cfg.Section("Interface").Key("PrivateKey").String()

	// Read the public key from the 'Peer' section
	publicKey := cfg.Section("Peer").Key("PublicKey").String()

	// Initialize a new IP scanner
	scanner := ipscanner.NewScanner(
		ipscanner.WithLogger(l.With(slog.String("subsystem", "scanner"))),
		ipscanner.WithWarpPing(),
		ipscanner.WithWarpPrivateKey(privateKey),
		ipscanner.WithWarpPeerPublicKey(publicKey),
		ipscanner.WithUseIPv4(opts.V4),
		ipscanner.WithUseIPv6(opts.V6),
		ipscanner.WithMaxDesirableRTT(opts.MaxRTT),
		ipscanner.WithCidrList(warp.WarpPrefixes()),
	)

	// Set a timeout context for the scan
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	// Start the IP scan
	scanner.Run(ctx)

	// Set up a ticker to check the scan progress
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	// Continuously check the scan progress until it's done
	for {
		ipList := scanner.GetAvailableIPs()
		if len(ipList) > 1 {
