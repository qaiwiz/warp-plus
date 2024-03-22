package ping

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner/internal/statute" // Importing the statute package for type definitions
)

// TcpPingResult struct holds the result of a TCP ping operation
type TcpPingResult struct {
	AddrPort netip.AddrPort   // The address and port being pinged
	RTT      time.Duration    // The round-trip time for the ping
	Err      error            // Any error encountered during the ping
}

// Result method returns the TcpPingResult as a statute.IPInfo struct
func (tp *TcpPingResult) Result() statute.IPInfo {
	return statute.IPInfo{AddrPort: tp.AddrPort, RTT: tp.RTT, CreatedAt: time.Now()}
}

// Error method returns the error encountered during the ping
func (tp *TcpPingResult) Error() error {
	return tp.Err
}

// String method returns a string representation of the TcpPingResult
func (tp *TcpPingResult) String() string {
	if tp.Err != nil {
		return fmt.Sprintf("%s", tp.Err)
	} else {
		return fmt.Sprintf("%s: time=%d ms", tp.AddrPort, tp.RTT)
	}
}

// TcpPing struct holds the configuration for a TCP ping operation
type TcpPing struct {
	host string            // The target host
	port uint16           // The target port
	ip   netip.Addr       // The target IP address
	opts statute.ScannerOptions  // Scanner options
}

// SetHost method sets the target host and resolves it to an IP address
func (tp *TcpPing) SetHost(host string) {
	tp.host = host
	tp.ip, _ = netip.ParseAddr(host) // Ignoring any errors for simplicity
}

// Host method returns the target host
func (tp *TcpPing) Host() string {
	return tp.host
}

// Ping method performs a TCP ping operation
func (tp *TcpPing) Ping() statute.IPingResult {
	return tp.PingContext(context.Background())
}

// PingContext method performs a TCP ping operation with a given context
func (tp *TcpPing) PingContext(ctx context.Context) statute.IPingResult {
	if !tp.ip.IsValid() {
		return &TcpPingResult{AddrPort: netip.AddrPort{}, RTT: 0, Err: errors.New("no IP specified")}
	}

