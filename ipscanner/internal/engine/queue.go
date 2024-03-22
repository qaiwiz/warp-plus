package engine

import (
	"log/slog"
	"sort"
	"sync"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// IPQueue represents a queue of IPInfo structs with a limited size and RTT-based sorting.
type IPQueue struct {
	queue        []statute.IPInfo       // The main IPInfo queue
	maxQueueSize int                    // Maximum queue size
	mu           sync.Mutex             // Mutex for thread safety
	available    chan struct{}           // Channel for signaling available queue slots
	maxTTL       time.Duration          // Maximum Time To Live for IPInfo entries
	rttThreshold time.Duration          // RTT threshold for queue membership
	inIdealMode  bool                   // Flag indicating if the queue is in ideal mode
	log          *slog.Logger            // Logger for the queue
	reserved     statute.IPInfQueue     // Queue for reserved IPInfos
}

// NewIPQueue creates a new IPQueue instance with the given options.
func NewIPQueue(opts *statute.ScannerOptions) *IPQueue {
	// Initialize a new IPQueue instance
	return &IPQueue{
		queue:        make([]statute.IPInfo, 0),
		maxQueueSize: opts.IPQueueSize,
		maxTTL:       opts.IPQueueTTL,
	
