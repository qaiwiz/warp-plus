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
	PublicKey    string
	// PreSharedKey is the base64-encoded pre-shared key of the peer.
	PreSharedKey string
	// Endpoint is the endpoint (IP address and port) of the peer.
	Endpoint     string
	// KeepAlive is the persistent keepalive interval for the peer.
	KeepAlive    int
	// AllowedIPs are the allowed IP addresses for the peer.
	AllowedIPs   []netip.Prefix
	// Trick is a flag indicating if the peer should be treated as a local peer.
	Trick        bool
}

// InterfaceConfig struct represents the configuration for a WireGuard interface.
type InterfaceConfig struct {
	// PrivateKey is the base64-encoded private key of the interface.
	PrivateKey string
	// Addresses are the IP addresses assigned to the interface.
	Addresses  []netip.Addr
	// DNS are the DNS servers assigned to the interface.
	DNS        []netip.Addr
	// MTU is the Maximum Transmission Unit for the interface.
	MTU        int
}

// Configuration struct represents the overall configuration for the WireGuard network.
type Configuration struct {
	// Interface is the configuration for the WireGuard interface.
	Interface *InterfaceConfig
	// Peers are the configurations for the peers in the WireGuard network.
	Peers     []PeerConfig
}

// encodeBase64ToHex takes a base64-encoded string and returns its hexadecimal representation.
func encodeBase64ToHex(key string) (string, error) {
	// Decode the base64-encoded string.
	decoded, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", fmt.Errorf("invalid base64 string: %s", key)
	}
	// Check if the decoded string is 32 bytes long.
	if len(decoded) != 32 {
		return "", fmt.Errorf("key should be 32 bytes: %s", key)
	}
	// Encode the decoded string to hexadecimal representation.
	return hex.EncodeToString(decoded), nil
}

// ParseInterface parses the [Interface] section of the configuration file and returns the InterfaceConfig.
func ParseInterface(cfg *ini.File) (InterfaceConfig, error) {
	// Initialize the InterfaceConfig struct.
	device := InterfaceConfig{}
	// Get the [Interface] section of the configuration file.
	interfaces, err := cfg.SectionsByName("Interface")
	if len(interfaces) != 1 || err != nil {
		return InterfaceConfig{}, errors.New("only one [Interface] is expected")
	}
	// Get the first (and only) [Interface] section.
	iface := interfaces[0]

	// Parse the Addresses field.
	key := iface.Key("Address")
	if key != nil {
		var addresses []netip.Addr
		// Split the string by ',' and parse each IP address.
		for _, str := range key.StringsWithShadows(",") {
			prefix, err := netip.ParsePrefix(str)
			if err != nil {
				return InterfaceConfig{}, err
			}
			addresses = append(addresses, prefix.Addr())
		}
		device.Addresses = addresses
	}

	// Parse the PrivateKey field.
	key = iface.Key("PrivateKey")
	if key == nil {
		return InterfaceConfig{}, errors.New("PrivateKey should not be empty")
	}
	// Encode the base64-encoded string to hexadecimal representation.
	privateKeyHex, err := encodeBase64ToHex(key.String())
	if err != nil {
		return InterfaceConfig{}, err
	}
	device.PrivateKey = privateKeyHex

	// Parse the DNS field.
	key = iface.Key("DNS")
	if key != nil {
		var addresses []netip.Addr
		// Split the string by ',' and parse each IP address.
		for _, str := range key.StringsWithShadows(",") {
			ip, err := netip.ParseAddr(str)
			if err != nil {
				return InterfaceConfig{}, err
			}
			addresses = append(addresses, ip)
		}
		device.DNS = addresses
	}

	// Parse the MTU field.
	if sectionKey, err := iface.GetKey("MTU"); err == nil {
		value, err := sectionKey.Int()
		if err != nil {
			return InterfaceConfig{}, err
		}
		device.MTU = value
	}

	// Return the parsed InterfaceConfig.
	return device, nil
}

// ParsePeers parses the [Peer] sections of the configuration file and returns a slice of PeerConfig.
func ParsePeers(cfg *ini.File) ([]PeerConfig, error) {
	// Get the [Peer] sections of the configuration file.
	sections, err := cfg.SectionsByName("Peer")
	if len(sections) < 1 || err != nil {
		return nil, errors.New("at least one [Peer] is expected")
	}

	// Initialize a slice of PeerConfig.
	peers := make([]PeerConfig, len(sections))
	// Parse each [Peer] section.
	for i, section := range sections {
		peer := PeerConfig{
			PreSharedKey: "0000000000000000000000000000000000000000
