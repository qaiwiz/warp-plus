package ping

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"net/url"
	"time"

	// ipscanner package for internal use
	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// HttpPingResult struct holds the result of an HTTP ping
type HttpPingResult struct {
	AddrPort netip.AddrPort   // The IP address and port
	Proto    string           // The protocol used
	Status   int              // The HTTP status code
	Length   int              // The length of the response
	RTT      time.Duration    // The round-trip time
	Err      error            // Any error encountered
}

// Result method returns the IPInfo struct based on the AddrPort and RTT
func (h *HttpPingResult) Result() statute.IPInfo {
	return statute.IPInfo{AddrPort: h.AddrPort, RTT: h.RTT, CreatedAt: time.Now()}
}

// Error method returns the error encountered during the HTTP ping
func (h *HttpPingResult) Error() error {
	return h.Err
}

// String method returns a string representation of the HttpPingResult
func (h *HttpPingResult) String() string {
	if h.Err != nil {
		return fmt.Sprintf("%s", h.Err)
	}

	return fmt.Sprintf("%s: protocol=%s, status=%d, length=%d, time=%d ms", h.AddrPort, h.Proto, h.Status, h.Length, h.RTT)
}

// HttpPing struct holds the configuration for an HTTP ping
type HttpPing struct {
	Method string
	URL    string
	IP     netip.Addr

	// ScannerOptions for customizing the HTTP client
	opts statute.ScannerOptions
}

// Ping method performs an HTTP ping
func (h *HttpPing) Ping() statute.IPingResult {
	return h.PingContext(context.Background())
}

// PingContext method performs an HTTP ping with a given context
func (h *HttpPing) PingContext(ctx context.Context) statute.IPingResult {
	// Parse the URL and validate the IP address
	u, err := url.Parse(h.URL)
	if err != nil {
		return h.errorResult(err)
	}
	orighost := u.Host

	if !h.IP.IsValid() {
		return h.errorResult(errors.New("no IP specified"))
	}

	// Create a new HTTP request with the given method and URL
	req, err := http.NewRequestWithContext(ctx, h.Method, h.URL, nil)
	if err != nil {
		return h.errorResult(err)
	}

	// Set the User-Agent and Referer headers based on the options
	ua := "httping"
	if h.opts.UserAgent != "" {
		ua = h.opts.UserAgent
	}
	req.Header.Set("User-Agent", ua)
	if h.opts.Referrer != "" {
		req.Header.Set("Referer", h.opts.Referrer)
	}
	req.Host = orighost

	// Create a custom HTTP client based on the options
	addr := netip.AddrPortFrom(h.IP, h.opts.Port)
	client := h.opts.HttpClientFunc(h.opts.RawDialerFunc, h.opts.TLSDialerFunc, h.opts.QuicDialerFunc, addr.String())

	// Disable redirects
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// Measure the round-trip time
	t0 := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return h.errorResult(err)
	}

	defer resp.Body.
