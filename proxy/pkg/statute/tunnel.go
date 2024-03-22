package statute

import (
	"context"
	"io"
	"net"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// isClosedConnError reports whether the error (err) is related to the use of a
// closed network connection.
func isClosedConnError(err error) bool {
	if err == nil {
		return false
	}

	str := err.Error()
	if strings.Contains(str, "use of closed network connection") {
	
