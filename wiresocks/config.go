package wiresocks

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/netip"

	"github.com/go-ini/ini"
)

// PeerConfig struct represents the configuration for a peer in the WireGuard network.
type PeerConfig struct {
	// PublicKey is the base64-encoded public key of the peer.
	PublicKey    string `ini:"PublicKey"`
	// PreSharedKey is the base64-encoded pre-shared key of the peer.
	PreSharedKey string `ini:"PreSharedKey"`
	// Endpoint is the endpoint (IP address and port) of the peer.
	Endpoint     string `ini:"Endpoint"`
	// KeepAlive is the persistent keepalive interval for the peer.
	KeepAlive    int    `ini:"KeepAlive"`
	// AllowedIPs are the allowed IP addresses for the peer.
	AllowedIPs   []netip.Prefix `ini:"AllowedIPs"`
	// Trick is a flag indicating if the peer should be treated as a local peer.
	Trick        bool   `ini:"Trick"`
}

// InterfaceConfig struct represents the configuration for a WireGuard interface.
type InterfaceConfig struct {
	// PrivateKey is the base64-encoded private key of the interface.
	PrivateKey string `ini:"PrivateKey"`
	// Addresses are the IP addresses assigned to the interface.
	Addresses  []netip.Addr `ini:"Address"`
	// DNS are the DNS servers assigned to the interface.
	DNS        []netip.Addr `ini:"DNS"`
	// MTU is the Maximum Transmission Unit for the interface.
	MTU        int    `ini:"MTU"`
}

// Configuration struct represents the overall configuration for the WireGuard network.
type Configuration struct {
	// Interface is the configuration for the WireGuard interface.
	Interface *InterfaceConfig
	// Peers are the configurations for the peers in the WireGuard network.
	Peers     []PeerConfig
}

// encodeBase6
