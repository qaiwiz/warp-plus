package ping

import (
	"errors"
	"fmt"
	"net/netip"

	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// Ping struct represents a ping object with Options.
type Ping struct {
	Options *statute.ScannerOptions
}

// DoPing performs a ping on the given IP address using the selected ping operation(s).
// It returns IPInfo and error.
func (p *Ping) DoPing(ip netip.Addr) (statute.IPInfo, error) {
	// Check if HTTP ping operation is selected.
	if p.Options.SelectedOps&statute.HTTPPing > 0 {
		res, err := p.httpPing(ip)
		if err != nil {
			return statute.IPInfo{}, err
		}

		return res, nil
	}
	// Check if TLS ping operation is selected.
	if p.Options.SelectedOps&statute.TLSPing > 0 {
		res, err := p.tlsPing(ip)
		if err != nil {
		
