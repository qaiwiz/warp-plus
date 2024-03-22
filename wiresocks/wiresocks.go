package wiresocks

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"

	"github.com/bepass-org/warp-plus/wireguard/conn"
	"github.com/bepass-org/warp-plus/wireguard/device"
	"github.com/bepass-org/warp-plus/wireguard/tun/netstack"
)

// StartWireguard creates a tun interface on netstack given a configuration
func StartWireguard(ctx context.Context, l *slog.Logger, conf *Configuration) (*VirtualTun, error) {
	// Initialize a new bytes buffer to store the wireguard configuration
	var request bytes.Buffer

	// Write the private key of the interface to the buffer
	request.WriteString(fmt.Sprintf("private_key=%s\n", conf.Interface.PrivateKey))

	// Loop through the peers and write their configuration to the buffer
	for _, peer := range conf.Peers {
		request.WriteString(fmt.Sprintf("public_key=%s\n", peer.PublicKey))
		request.WriteString(fmt.Sprintf("persistent_keepalive_interval=%d\n", peer.KeepAlive))
		request.WriteString(fmt.Sprintf("preshared_key=%s\n", peer.PreSharedKey))
		request.WriteString(fmt.Sprintf("endpoint=%s\n", peer.Endpoint))
		request.WriteString(fmt.Sprintf("trick=%t\n", peer.Trick))

		// Write the allowed IPs for the peer to the buffer
		for _, cidr := range peer.AllowedIPs {
			request.WriteString(fmt.Sprintf("allowed_ip=%s\n", cidr))
		}
	}

	// Create a new tun interface and network stack with the specified addresses, DNS, and MTU
	tun, tnet, err := netstack.CreateNetTUN(conf.Interface.Addresses, conf.Interface.DNS, conf.Interface.MTU)
	if err != nil {
		return nil, err
	}

	// Initialize a new wireguard device with the tun interface, a default bind, and a logger
	dev := device.NewDevice(tun, conn.NewDefaultBind(), device.NewSLogger(l.With("subsystem", "wireguard-go")))

	// Set the wireguard interface configuration
	err = dev.IpcSet(request.String())
