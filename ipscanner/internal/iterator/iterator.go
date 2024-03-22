package iterator

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net"
	"net/netip"

	"github.com/bepass-org/warp-plus/ipscanner/internal/cache"
	"github.com/bepass-org/warp-plus/ipscanner/internal/statute"
)

// LCG represents a linear congruential generator with full period.
type LCG struct {
	modulus    *big.Int // Modulus of the LCG
	multiplier *big.Int // Multiplier of the LCG
	increment  *big.Int // Increment of the LCG
	current    *big.Int // Current value of the LCG
}

// NewLCG creates a new LCG instance with a given size.
func NewLCG(size *big.Int) *LCG {
	// ... (code unchanged)
}

// checkHullDobell checks if the given parameters satisfy the Hull-Dobell Theorem.
func checkHullDobell(modulus, multiplier, increment *big.Int) bool {
	// ... (code unchanged)
}

// Next generates the next number in the sequence.
func (lcg *LCG) Next() *big.Int {
	// ... (code unchanged)
}

// ipRange represents a range of IP addresses with associated LCG and size.
type ipRange struct {
	lcg   *LCG   // LCG for generating IP addresses in the range
	start netip.Addr  // Starting IP address of the range
	stop  netip.Addr  // Ending IP address of the range
	size  *big.Int  // Size of the range
	index *big.Int  // Current index of the range
}

// newIPRange creates a new ipRange instance for a given CIDR prefix.
func newIPRange(cidr netip.Prefix) (ipRange, error) {
	// ... (code unchanged)
}

// lastIP calculates the last IP address in a given CIDR prefix.
func lastIP(prefix netip.Prefix) netip.Addr {
	// ... (code unchanged)
}

// ipToBigInt converts a netip.Addr to a *big.Int.
func ipToBigInt(ip netip.Addr) *big.Int {
	// ... (code unchanged)
}

// bigIntToIP converts a *big.Int to a netip.Addr.
func bigIntToIP(n *big.Int) netip.Addr {
	// ... (code unchanged)
}

// addIP adds a given big integer value to a netip.Addr.
func addIP(ip netip.Addr, num *big.Int) netip.Addr {
	// ... (code unchanged)
}

// ipRangeSize calculates the size of an IP range based on a given CIDR prefix.
func ipRangeSize(prefix netip.Prefix) *big.Int {
	// ... (code unchanged)
}

// IpGenerator generates IP addresses from a list of ipRanges.
type IpGenerator struct {
	ipRanges []ipRange // List of ipRanges to generate IP addresses from
}

// NextBatch generates a batch of IP addresses from the ipRanges.
func (g *IpGenerator) NextBatch() ([]netip.Addr, error) {
	// ... (code unchanged)
}

// shuffleSubnetsIpRange shuffles a slice of ipRange using crypto/rand.
func shuffleSubnetsIpRange(subnets []ipRange) error {
	// ... (code unchanged)
}

// NewIterator creates a new IpGenerator instance with a given statute.ScannerOptions.
func NewIterator(opts *statute.ScannerOptions) *IpGenerator {
	// ... (code unchanged)
}
