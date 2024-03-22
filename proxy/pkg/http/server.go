package http

import (
	"bufio"
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strconv"

	"github.com/bepass-org/warp-plus/proxy/pkg/statute"
)

// Server struct represents an HTTP proxy server with customizable options.
type Server struct {
	// bind is the address to listen on
	Bind string

	Listener net.Listener // The listener for incoming connections

	// ProxyDial specifies the optional proxyDial function for
	// establishing the transport connection.
	ProxyDial statute.ProxyDialFunc
	// UserConnectHandle gives the user control to handle the TCP CONNECT requests
	UserConnectHandle statute.UserConnectHandler
	// Logger error log
	Logger *slog.Logger
	// Context is default context
	Context context.Context
	// BytesPool getting and returning temporary bytes for use by io.CopyBuffer
	BytesPool statute.BytesPool
}

// NewServer creates a new Server instance with default options.
func NewServer(options ...ServerOption) *Server {
	s := &Server{
		Bind:      statute.DefaultBindAddress,
		ProxyDial: statute.DefaultProxyDial(),
		Logger:    slog.Default(),
		Context:   statute.DefaultContext(),
	}

	// Apply the provided options to the Server instance
	for _, option := range options {
		option(s)
	}

	return s
}

// ServerOption is a functional option for configuring the Server.
type ServerOption func(*Server)

// ListenAndServe starts the HTTP proxy server and handles incoming connections.
func (s *Server) ListenAndServe() error {
	// Create a new listener if not provided
	if s.Listener == nil {
		ln, err := net.Listen("tcp", s.Bind)
		if err != nil {
			return err // Return error if binding was unsuccessful
		}
		s.Listener = ln
	}

	// Log the listening address
	s.Logger.Debug("started proxy", "address", s.Bind)

	// Ensure the listener will be closed
	defer func() {
		_ = s.Listener.Close()
	}()

	// Create a cancelable context based on s.Context
	ctx, cancel := context.WithCancel(s.Context)
	defer cancel() // Ensure resources are cleaned up

	// Accept and serve connections concurrently
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			conn, err := s.Listener.Accept()
			if err != nil {
				s.Logger.Error(err.Error())
				continue
			}

			// Handle each connection in a new goroutine
			go func() {
				err := s.ServeConn(conn)
				if err != nil {
					s.Logger.Error(err.Error()) // Log errors from ServeConn
				}
			}()
		}
	}
}

// WithLogger sets the Logger for the Server.
func WithLogger(logger *slog.Logger) ServerOption {
	return func(s *Server) {
		s.Logger = logger
	}
}

// WithBind sets the bind address for the Server.
func WithBind(bindAddress string) ServerOption {
	return func(s *Server) {
		s.Bind = bindAddress
	}
}

// WithConnectHandle sets the UserConnectHandle for the Server.
func WithConnectHandle(handler statute.UserConnectHandler) ServerOption {
	return func(s *Server) {
		s.UserConnectHandle = handler
	}
}

// WithProxyDial sets the ProxyDial for the Server.
func WithProxyDial(proxyDial statute.ProxyDialFunc) ServerOption {
	return func(s *Server) {
		s.ProxyDial = proxyDial
	}
}

// WithContext sets the Context for the Server.
func WithContext(ctx context.Context) ServerOption {
	return func(s *Server) {
		s.Context = ctx
	}
}

// WithBytesPool sets the BytesPool for the Server.
func WithBytesPool(bytesPool statute.BytesPool) ServerOption {
	return func(s *Server) {
		s.BytesPool = bytesPool
	}
}

// ServeConn handles an individual connection.
func (s *Server) ServeConn(conn net.Conn) error {
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		return err
	}

	return s.handleHTTP(conn, req, req.Method == http.MethodConnect)
}

// handleHTTP handles HTTP requests and CONNECT requests differently.
func (s *Server) handleHTTP(conn net.Conn, req *http.Request, isConnectMethod bool) error {
	// If UserConnectHandle is not provided, use the default handling
	if s.UserConnectHandle == nil {
		return s.embedHandleHTTP(conn, req, isConnectMethod)
	}

	// Handle CONNECT requests
	if isConnectMethod {
		_, err := conn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
		if err != nil {
			return err
		}
	} else {
		// Handle non-CONNECT requests
		cConn := &customConn{
			Conn: conn,
			req:  req,
		}
		conn = cConn
	}

	// Parse the target address
	targetAddr := req.URL.Host
	host, portStr, err := net.SplitHostPort(targetAddr)
	if err != nil {
		host = targetAddr
		if req.URL.Scheme == "https" || isConnectMethod {
