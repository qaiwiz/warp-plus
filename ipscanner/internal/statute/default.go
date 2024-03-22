package statute

import (
	// ... (import statements)
)

// FinalOptions is a global variable that holds the configuration for the scanner.
var FinalOptions *ScannerOptions

// DefaultHTTPClientFunc creates an HTTP client with custom dialers and options.
func DefaultHTTPClientFunc(rawDialer TDialerFunc, tlsDialer TDialerFunc, quicDialer TQuicDialerFunc, targetAddr ...string) *http.Client {
	// ... (code)

	// Create a new http.RoundTripper based on the user's preferences.
	var transport http.RoundTripper
	if FinalOptions.UseHTTP3 {
		// Create a new http3.RoundTripper if HTTP/3 is enabled.
		transport = &http3.RoundTripper{
			// ... (configuration)
		}
	} else {
		// Create a new http.Transport if HTTP/3 is disabled.
		trans := &http.Transport{
			// ... (configuration)
		}
		transport = trans
	}

	// Return a new http.Client with the custom RoundTripper.
	return &http.Client{
		Transport: transport,
		Timeout:   FinalOptions.ConnectionTimeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

// DefaultDialerFunc creates a new dialer with a custom connection timeout.
func DefaultDialerFunc(ctx context.Context, network, addr string) (net.Conn, error) {
	// ... (code)

	// Create a new Dialer with the custom connection timeout.
	d := &net.Dialer{
		Timeout: FinalOptions.ConnectionTimeout, // Connection timeout
		// ... (other custom settings)
	}
	return d.DialContext(ctx, network, addr)
}

// getServerName extracts the server name from the given address.
func getServerName(address string) (string, error) {
	// ... (code)
}

// defaultTLSConfig creates a new TLS config based on the user's preferences.
func defaultTLSConfig(addr string) *tls.Config {
	// ... (code)

	// Create a new TLS config with the user's preferences.
	return &tls.Config{
		InsecureSkipVerify: allowInsecure || FinalOptions.InsecureSkipVerify,
		ServerName:         sni,
		MinVersion:         FinalOptions.TlsVersion,
		MaxVersion:         FinalOptions.TlsVersion,
		NextProtos:         alpnProtocols,
	}
}

// DefaultTLSDialerFunc creates a new TLS dialer with a custom handshake timeout.
func DefaultTLSDialerFunc(ctx context.Context, network, addr string) (net.Conn, error) {
	// ... (code)

	// Create a new TLS client connection with the custom TLS config.
	tlsClientConn := tls.Client(rawConn, defaultTLSConfig(addr))

	// Perform the handshake with a timeout.
	err = tlsClientConn.SetDeadline(time.Now().Add(FinalOptions.HandshakeTimeout))
	if err != nil {
		// ... (error handling)
	}

	// Perform the handshake.
	err = tlsClientConn.Handshake()
	if err != nil {
		// ... (error handling)
	}

	// Reset the deadline for future I/O operations.
	err = tlsClientConn.SetDeadline(time.Time{})
	if err != nil {
		// ... (error handling)
	}

	// Return the established TLS connection.
	return tlsClientConn, nil
}

// DefaultQuicDialerFunc creates a new QUIC dialer with custom timeout options.
func DefaultQuicDialerFunc(ctx context.Context, addr string, _ *tls.Config, _ *quic.Config) (quic.EarlyConnection, error) {
	// ... (code)

	// Create a new QUIC config with the user's preferences.
	quicConfig := &quic.Config{
		MaxIdleTimeout:       FinalOptions.ConnectionTimeout,
		HandshakeIdleTimeout: FinalOptions.HandshakeTimeout,
	}

	// Dial the QUIC address with the custom QUIC config.
	return quic.DialAddrEarly(ctx, addr, defaultTLSConfig(addr), quicConfig)
}

// DefaultCFRanges returns the default Cloudflare IP ranges.
func DefaultCFRanges() []netip.Prefix {
	// ... (code)

	// Return the default Cloudflare IP ranges.
	return []netip.Prefix{
		// ... (IP ranges)
	}
}
