package statute

import (
	"context"
	"fmt"
	"io"
	"net"
)

// Logger interface for logging debug and error messages
type Logger interface {
	Debug(v ...interface{})   // Function to log debug messages
	Error(v ...interface{})   // Function to log error messages
}

// DefaultLogger struct for default logger implementation
type DefaultLogger struct{}

func (l DefaultLogger) Debug(v ...interface{}) { // Function to log debug messages
	fmt.Println(v...)
}

func (l DefaultLogger) Error(v ...interface{}) { // Function to log error messages
	fmt.Println(v...)
}

// ProxyRequest struct for proxy request parameters
type ProxyRequest struct {
	Conn        net.Conn            // Underlying connection
	Reader      io.Reader           // Reader for the connection
	Writer      io.Writer           // Writer for the connection
	Network     string              // Network type (e.g. tcp, udp)
	Destination string              // Destination address and port (e.g. 192.168.1.1:80)
	DestHost    string              // Destination host
	DestPort    int32               // Destination port
}

// UserConnectHandler type for user-defined connection handlers
type UserConnectHandler func(request *ProxyRequest) error // Function to handle user connection requests

// UserAssociateHandler type for user-defined associate handlers
type UserAssociateHandler func(request *ProxyRequest) error // Function to handle user associate requests

// ProxyDialFunc type for proxy dial functions
type ProxyDialFunc func(ctx context.Context, network string, address string) (net.Conn, error) // Function to establish a connection using a proxy

// DefaultProxyDial function for default ProxyDialFunc implementation
func DefaultProxyDial() ProxyDialFunc {
	var dialer net.Dialer
	return dialer.DialContext
}

// ProxyListenPacket type for proxy listen packet functions
type ProxyListenPacket func(ctx context.Context, network string, address string) (net.PacketConn, error) // Function to establish a packet connection using a proxy

// DefaultProxyListenPacket function for default ProxyListenPacket implementation
func DefaultProxyListenPacket() ProxyListenPacket {
	var listener net.ListenConfig
	return listener.ListenPacket
}

// PacketForwardAddress type for packet forwarding address functions
type PacketForwardAddress func(ctx context.Context, destinationAddr string,
	packet net.PacketConn, conn net.Conn) (net.IP, int, error) // Function to forward packets to the specified address

// BytesPool interface for managing temporary bytes for io.CopyBuffer
type BytesPool interface {
	Get() []byte // Function to get temporary bytes
	Put([]byte)  // Function to return temporary bytes
}

// DefaultContext function for default context.Context implementation
func DefaultContext() context.Context {
	return context.Background()
}

const DefaultBindAddress = "127.0.0.1:1080" // Default bind address for the proxy server
