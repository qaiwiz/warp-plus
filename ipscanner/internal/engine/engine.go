package engine

import (
	"context"
	"log/slog"
	"net/netip"
	"time"

	"github.com/bepass-org/warp-plus/ipscanner/internal/iterator"
	"github.com/bepass-org/warp-plus/ipscanner/internal/ping"
	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// Engine represents the scanner engine that generates IP addresses, pings them,
// and manages the IP queue.
type Engine struct {
	generator *iterator.IpGenerator   // IP address generator
	ipQueue   *IPQueue                // IP address queue
	ping      func(netip.Addr) (statute.IPInfo, error) // Ping function to check IP availability
	log       *slog.Logger             // Logger for the engine
}

// NewScannerEngine initializes and returns a new ScannerEngine instance.
func NewScannerEngine(opts *statute.ScannerOptions) *Engine {
	queue := NewIPQueue(opts)

	p := ping.Ping{
		Options: opts,
	}
	return &Engine{
		ipQueue:   queue,
		ping:      p.DoPing,
		generator: iterator.NewIterator(opts),
		log:       opts.Logger.With(slog.String("subsystem", "scanner/engine")),
	}
}

// GetAvailableIPs returns a slice of available IPInfo instances based on the engine's IP queue.
func (e *Engine) GetAvailableIPs(desc bool) []statute.IPInfo {
	if e.ipQueue != nil {
		return e.ipQueue.AvailableIPs(desc)
	}
	return nil
}

// Run is the main function of the engine that runs the IP scanning process in a loop.
func (e *Engine) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-e.ipQueue.available:
	
