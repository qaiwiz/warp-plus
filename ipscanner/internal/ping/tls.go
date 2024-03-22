package ping

import (
	// Importing necessary packages
	"context"
	"errors"
	"fmt"
	"net/netip"
	"time"

	// Importing statute package for IPInfo, ScannerOptions, and related types
	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// TlsPingResult represents the result of a TLS ping operation
type TlsPingResult struct {
	// The address and port that were pinged
	AddrPort   netip.AddrPort   //
