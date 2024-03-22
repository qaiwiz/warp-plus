package ping

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner/internal/statute" // Importing statute package for IPInfo, ScannerOptions, and related types
)

// TlsPingResult represents the result of a TLS ping operation
type TlsPingResult struct {
	AddrPort   netip.AddrPort   // The address and port that were pinged
	TLSVersion uint16           // The TLS version used during the ping
	RTT        time.Duration    // The round-trip time for the ping
	Err        error            // Any error that occurred during the ping
}

// Result returns the IPInfo representation of the TlsPingResult
func (t *TlsPingResult) Result() statute.IPInfo {
	return statute.IPInfo{AddrPort: t.AddrPort, RTT: t.RTT, CreatedAt: time.Now()}
}

// Error returns the error associated with the TlsPingResult
func (t *TlsPingResult) Error() error {
	return t.Err
}

// String returns a string representation of the TlsPingResult
func (t *TlsPingResult) String() string {
	if t.Err != nil {
		return fmt.Sprintf("%s", t.Err)
	}

	return fmt.Sprintf("%s: protocol=%s, time=%d ms", t.AddrPort, statute.TlsVersionToString(t.TLSVersion), t.RTT)
}

// TlsPing represents a TLS ping operation
type TlsPing struct {
	Host string            // The hostname associated with the IP address
	Port uint16            // The port number to ping
	IP   netip.Addr        // The IP address to ping
	opts *statute.ScannerOptions  // Options for the ping operation
}

// Ping performs a TLS ping operation without a context
func (t *TlsPing) Ping() statute.IPingResult {
	return t.PingContext(context.Background())
}

// PingContext performs a TLS ping operation with a given context
func (t *TlsPing) PingContext(ctx context.Context) statute.IPingResult {
	if !t.IP.IsValid() {
		return t.errorResult(errors.New("no IP specified"))
	}
	addr := netip.AddrPortFrom(t.IP, t.Port)
	t0 := time.Now()
	client, err := t.opts.TLSDialerFunc(ctx, "tcp", addr.String()) // Dial
