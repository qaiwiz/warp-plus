// HttpPingResult struct holds the result of an HTTP ping
type HttpPingResult struct {
	// AddrPort is the IP address and port of the target
	AddrPort netip.AddrPort
	// Proto is the protocol used for the HTTP ping
	Proto    string
	// Status is the HTTP status code of the response
	Status   int
	// Length is the length of the response
	Length   int
	// RTT is the round-trip time in duration
	RTT      time.Duration
	// Err is any error encountered during the HTTP ping
	Err      error
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
	// Method is the HTTP method to use for the ping
	Method string
	// URL is the target URL for the HTTP ping
	URL    string
	// IP is the target IP address for the HTTP ping
	IP     netip.Addr

	// opts is an instance of ScannerOptions for customizing the HTTP client
	opts statute.ScannerOptions
}

// Ping method performs an HTTP ping with the default context
func (h *HttpPing) Ping() statute.IPingResult {
	return h.PingContext(context.Background())
}

// PingContext method performs an HTTP ping with a given context
func (h *HttpPing) PingContext(ctx context.Context) statute.IPingResult {
	// ... (rest of the function)
}

