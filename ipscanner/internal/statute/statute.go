package statute

import (
	// Importing necessary packages

	"context"
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"
	"net/netip"
	"time"

	// Importing third-party packages

	"github.com/quic-go/quic-go"
)

// TIPQueueChangeCallback is a type for a function that handles changes in the IP queue
type TIPQueueChangeCallback func(ips []IPInfo)

// Type definitions
type (
	TDialerFunc     func(ctx context.Context, network, addr string) (net.Conn, error)
	TQuicDialerFunc func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (quic.EarlyConnection, error)
	THTTPClientFunc func(rawDialer TDialerFunc, tlsDialer TDialerFunc, quicDialer TQuicDialerFunc, targetAddr ...string) *http.Client
)

// Constants for various network scanning operations
const (
	HTTPPing = 1 << 1
	TLSPing  = 1 << 2
	TCPPing  = 1 << 3
	QUICPing = 1 << 4
	WARPPing = 1 << 5
)

// IPInfo struct contains information about an IP address
type IPInfo struct {
	AddrPort  netip.AddrPort // Combination of IP address and port number
	RTT       time.Duration  // Round-trip time
	CreatedAt time.Time      // Timestamp when the IP was added to the queue
}

// ScannerOptions struct holds configuration options for the network scanner
type ScannerOptions struct {
	// Scanner options
	UseIPv4               bool
	UseIPv6               bool
	CidrList              []netip.Prefix // List of CIDR ranges to scan
	SelectedOps           int
	Logger                *slog.Logger
	InsecureSkipVerify    bool
	RawDialerFunc         TDialerFunc
	TLSDialerFunc         TDialerFunc
	QuicDialerFunc        TQuicDialerFunc
	HttpClientFunc        THTTPClientFunc
	UseHTTP3              bool
	UseHTTP2              bool
	DisableCompression    bool
	HTTPPath              string
	Referrer              string
	UserAgent             string
	Hostname              string
	WarpPrivateKey        string
	WarpPeerPublicKey     string
	WarpPresharedKey      string
	Port                  uint16
	IPQueueSize           int
	IPQueueTTL            time.Duration
	MaxDes
