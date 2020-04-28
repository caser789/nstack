package stack

import (
	"github.com/caser789/nstack/tcpip"
	"github.com/caser789/nstack/tcpip/buffer"
)

// TransportEndpointID is the identifier of a transport layer protocol endpoint.
type TransportEndpointID struct {
	// LocalPort is the local port associated with the endpoint.
	LocalPort uint16

	// LocalAddress is the local [network layer] address associated with
	// the endpoint.
	LocalAddress tcpip.Address

	// RemotePort is the remote port associated with the endpoint.
	RemotePort uint16

	// RemoteAddress is the remote [network layer] address associated with
	// the endpoint.
	RemoteAddress tcpip.Address
}

// TransportEndpoint is the interface that needs to be implemented by transport
// protocol (e.g., tcp, udp) endpoints that can handle packets.
type TransportEndpoint interface {
	// HandlePacket is called by the stack when new packets arrive to
	// this transport endpoint.
	HandlePacket(r *Route, id TransportEndpointID, v buffer.View)
}
