package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/netip"
	"os"
	"path/filepath"

	"github.com/bepass-org/warp-plus/psiphon"
	"github.com/bepass-org/warp-plus/warp"
	"github.com/bepass-org/warp-plus/wiresocks"
)

const singleMTU = 1400
const doubleMTU = 1320

// WarpOptions holds the configuration options for running Warp.
type WarpOptions struct {
	Bind     netip.AddrPort
	Endpoint string
	License  string
	Psiphon  *PsiphonOptions
	Gool     bool
	Scan     *wiresocks.ScanOptions
}

// PsiphonOptions holds the configuration options for running Psiphon.
type PsiphonOptions struct {
	Country string
}

// RunWarp runs Warp with the given options.
func RunWarp(ctx context.Context, l *slog.Logger, opts WarpOptions) error {
	// Check if Psiphon and Gool are not set at the same time.
	if opts.Psiphon != nil && opts.Gool {
		return errors.New("can't use psiphon and gool at the same time")
	}

	// Check if a country is provided when using Psiphon.
	if opts.Psiphon != nil && opts.Psiphon.Country == "" {
		return errors.New("must provide country for psiphon")
	}

	// Create necessary directories.
	if err := makeDirs(); err != nil {
		return err
	}
	l.Debug("'primary' and 'secondary' directories are ready")

	// Change the current working directory to 'stuff'.
	if err := os.Chdir("stuff"); err != nil {
		return fmt.Errorf("error changing to 'stuff' directory: %w", err)
	}
	l.Debug("Changed working directory to 'stuff'")

	// Create primary and secondary identities.
	if err := createPrimaryAndSecondaryIdentities(l.With("subsystem", "warp/account"), opts.License); err != nil {
		return err
	}

	// Decide the working scenario based on the provided options.
	endpoints := []string{opts.Endpoint, opts.Endpoint}

	if opts.Scan != nil {
		res, err := wiresocks.RunScan(ctx, l, *opts.Scan)
		if err != nil {
			return err
		}

		l.Info("scan results", "endpoints", res)

		endpoints = make([]string, len(res))
		for i := 0; i < len(res); i++ {
			endpoints[i] = res[i].AddrPort.String()
		}
	}
	l.Info("using warp endpoints", "endpoints", endpoints)

	var warpErr error
	switch {
	case opts.Psiphon != nil:
		l.Info("running in Psiphon (cfon) mode")
		// Run primary warp on a random TCP port and run psiphon on bind address.
		warpErr = runWarpWithPsiphon(ctx, l, opts.Bind, endpoints[0], opts.Psiphon.Country)
	case opts.Gool:
		l.Info("running in warp-in-warp (gool) mode")
		// Run warp in warp.
		warpErr = runWarpInWarp(ctx, l, opts.Bind, endpoints)
	default:
		l.Info("running in normal warp mode")
		// Just run primary warp on bindAddress.
		warpErr = runWarp(ctx, l, opts.Bind, endpoints[0])
	}

	return warpErr
}

// runWarp runs primary warp on the given bind address and endpoint.
func runWarp(ctx context.Context, l *slog.Logger, bind netip.AddrPort, endpoint string) error {
	// Parse the configuration from the profile file.
	conf, err := wiresocks.ParseConfig("./primary/wgcf-profile.ini", endpoint)
	if err != nil {
		return err
	}
	conf.Interface.MTU = singleMTU

	// Update the keep-alive and trick settings for all peers.
	for i, peer := range conf.Peers {
		peer.Trick = true
		peer.KeepAlive = 3
		conf.Peers[i] = peer
	}

	// Start Wireguard with the given configuration.
	tnet, err := wiresocks.StartWireguard(ctx, l, conf)
	if err != nil {
		return err
	}

	// Start a proxy server on the given bind address.
	_, err = tnet.StartProxy(bind)
	if err != nil {
		return err
	}

	l.Info("serving proxy", "address", bind)

	return nil
}

// runWarpWithPsiphon runs primary warp on a random TCP port and runs psiphon on the bind address.
func runWarpWithPsiphon(ctx context.Context, l *slog.Logger, bind netip.AddrPort, endpoint string, country string) error {
	// Parse the configuration from the profile file.
	conf, err := wiresocks.ParseConfig("./primary/wgcf-profile.ini", endpoint)
	if err != nil {
		return err
	}
	conf.Interface.MTU = singleMTU

	// Update the keep-alive and trick settings for all peers.
	for i, peer := range conf.Peers {
		peer.Trick = true
		peer.KeepAlive = 3
		conf.Peers[i] = peer
	}


