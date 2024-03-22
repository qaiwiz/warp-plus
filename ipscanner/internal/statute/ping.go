package statute

import (
	"context"
	"crypto/tls"
	"fmt"
)

// IPingResult is an interface that defines the methods for an IPing result.
// It includes methods for getting the IPInfo result, the error (if any),
// and a String() method for formatting the result as a string.
type IPingResult interface {
	// Result returns the IPInfo associated with the ping result.
	Result() IPInfo

	// Error returns any error that occurred during the ping.
	Error() error

	// String implements the Stringer interface, providing a string
	// representation of the IPingResult.
	fmt.Stringer
}

// IPing is an interface for a type that can perform a ping operation.
// It includes methods for performing a ping with and without a context,
// both of which return an IPingResult.
type IPing interface {
	// Ping performs a ping operation and returns an IPingResult.
	Ping() IPingResult

	// PingContext performs a ping operation within the context of a given
	// context and returns an IPingResult.
	PingContext(context.Context) IPingResult
}

// TlsVersionToString is a function that converts a TLS version number
// (represented as a uint16) into a human-readable string.
func TlsVersionToString(ver uint16) string {
	// Use a switch statement to handle each possible TLS version.
	switch ver {
	case tls.VersionSSL30:
		return "SSL 3.0"
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		// If the version is not recognized, return "unknown".
		return "unknown"
	}
}
``
