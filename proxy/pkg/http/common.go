package http

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"sync"
)

// copyBuffer is a helper function to copy data between two net.Conn objects.
// It is currently commented out, but it can be used with io.CopyBuffer() to
// provide a custom buffer size.

type responseWriter struct {
	conn    net.Conn        // The underlying net.Conn object for the response.
	headers http.Header     // A http.Header object to store and manage response headers.
	status  int             // The HTTP status code for the response.
	written bool            // A flag to indicate if the headers and status code have been written.
}

// NewHTTPResponseWriter creates a new custom http.ResponseWriter that wraps a net.Conn.
func NewHTTPResponseWriter(conn net.Conn) http.ResponseWriter {
	return &responseWriter{
	
