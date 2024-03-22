package mixed

import (
	"bufio"
	"context"
	"log/slog"
	"net"

	"github.com/bepass-org/warp-plus/proxy/pkg/http"
	"github.com/bepass-org/warp-plus/proxy/pkg/socks4"
	"github.com/bepass-org/warp-plus/proxy/pkg/socks5"
	"github.com/bepass-org/warp-plus/proxy/pkg/statute"
)

// userHandler is a type for user-defined request handler function
type userHandler func(request *statute.ProxyRequest) error

// Proxy represents a proxy server with SOCKS4, SOCKS5, and HTTP support
type Proxy struct {
	// bind is the address to listen on
	bind string

	listener net.Listener // The listener for incoming connections

	// socks5Proxy is a SOCKS5 server with TCP and UDP support
	socks5Proxy *socks5.Server
	// socks4Proxy is a SOCKS4 server with TCP support
	socks4Proxy *socks4.Server
	// httpProxy is an HTTP proxy server with HTTP and HTTP-CONNECT support
	httpProxy *http.Server
	// userConnectHandle is a user handler for TCP and UDP requests (its general handler)
	userHandler userHandler
	// if user doesn't set userHandler, it can specify userTCPHandler for manual handling of TCP requests
	userTCPHandler userHandler
	// if user doesn't set userHandler, it can specify userUDPHandler for manual handling of UDP requests
	userUDPHandler userHandler
	// overwrite dial functions of http, socks4, socks5
	userDialFunc statute.ProxyDialFunc
	// logger error log
	logger *slog.Logger
	// ctx is default context
	ctx context.Context
}

// NewProxy creates a new Proxy instance with default settings and applies given options
func NewProxy(options ...Option) *Proxy {
	p := &Proxy{
		bind:         statute.DefaultBindAddress,
		socks5Proxy:  socks5.NewServer(),
		socks4Proxy:  socks4.NewServer(),
		httpProxy:    http.NewServer(),
		userDialFunc: statute.DefaultProxyDial(),
		logger:       slog.Default(),
		ctx:          statute.DefaultContext(),
	}

	for _, option := range options {
		option(p)
	}

	return p
}

// Option is a functional option for configuring a Proxy instance
type Option func(*Proxy)

// SwitchConn wraps a net.Conn and a bufio.Reader
type SwitchConn struct {
	net.Conn
	reader *bufio.Reader
}

// NewSwitchConn creates a new SwitchConn
func NewSwitchConn(conn net.Conn) *SwitchConn {
	return &SwitchConn{
		Conn:   conn,
		reader: bufio.NewReader(conn),
	}
}

// Read reads data into p, first from the bufio.Reader, then from the net.Conn
func (c *SwitchConn) Read(p []byte) (n int, err error) {
	return c.reader.Read(p)
}

// ListenAndServe starts the proxy server and listens for incoming connections
func (p *Proxy) ListenAndServe() error {
	// Create a new listener
	if p.listener == nil {
		ln, err := net.Listen("tcp", p.bind)
		if err != nil {
			return err // Return error if binding was unsuccessful
		}
		p.listener = ln
	}

	p.bind = p.listener.Addr().(*net.TCPAddr).String()
	p.logger.Debug("started proxy", "address", p.bind)

	// ensure listener will be closed
	defer func() {
		_ = p.listener.Close()
	}()

	// Create a cancelable context based on p.Context
	ctx, cancel := context.WithCancel(p.ctx)
	defer cancel() // Ensure resources are cleaned up

	// Start to accept connections and serve them
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := p.listener.Accept()
			if err != nil {
				p.logger.Error(err.Error())
				continue
			}

			// Start a new goroutine to handle each connection
			// This way, the server can handle multiple connections concurrently
			go func() {
				err := p.handleConnection(conn)
				if err != nil {
					p.logger.Error(err.Error()) // Log errors from ServeConn
				}
			}()
		}
	}
}

// handleConnection handles an incoming connection based on the proxy protocol
func (p *Proxy) handleConnection(conn net.Conn) error {
	// Create a SwitchConn
	switchConn := NewSwitchConn(conn)

	// Read one byte to determine the protocol
	buf := make([]byte, 1)
	_, err := switchConn.Read(buf)
	if err != nil {
		return err
	}

	// Unread the byte so it's available for the next read
	err = switchConn.reader.UnreadByte()
	if err != nil {
		return err
	}

	switch {
	case buf[0] == 5:
		// SOCKS5 protocol
		err = p.socks5Proxy.ServeConn(switchConn)
	case buf[0] == 4:
		// SOCKS4 protocol
		err = p.socks4Proxy.ServeConn(switchConn)
	default:
		// HTTP protocol
		err = p.httpProxy
