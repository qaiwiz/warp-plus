package ping

import (
	"errors"
	"fmt"
	"net/netip"

	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// Ping struct represents a ping object with Options.
// The Options field is a pointer to a statute.ScannerOptions struct,
// which contains various configuration options for the ping operation.
type Ping struct {
	Options *statute.ScannerOptions
}

// DoPing performs a ping on the given IP address using the selected ping operation(s).
// It returns IPInfo and error.
// The function first checks if the HTTP ping operation is selected by performing a bitwise AND
// operation with the SelectedOps field and the statute.HTTPPing constant.
// If the result is greater than 0, it means that the HTTP ping operation is selected.
// The function then calls the httpPing method and returns the result and any error encountered.
// If the HTTP ping operation is not selected, the function checks if the TLS ping operation is selected
// by performing a bitwise AND operation with the SelectedOps field and the statute.TLSPing constant.
// If the result is greater than 0, it means that the TLS ping operation is selected.
// The function then calls the tlsPing method and returns the result and any error encountered.
func (p *Ping) DoPing(ip netip.Addr) (statute.IPInfo, error) {
	// Check if HTTP ping operation is selected.
	if p.Options.SelectedOps&statute.HTTPPing > 0 {
		res, err := p.httpPing(ip)
		if err != nil {
			return statute.IPInfo{}, err
	
