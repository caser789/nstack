package stack

import (
	"github.com/caser789/nstack/tcpip"
	"github.com/caser789/nstack/tcpip/buffer"
)

// NetworkEndpointID is the identifier of a network layer protocol endpoint.
// Currently the local address is sufficient because all supported protocols
// (i.e., IPv4 and IPv6) have different sizes for their addresses.
type NetworkEndpointID struct {
	LocalAddress tcpip.Address
}

// NetworkEndpoint is the interface that needs to be implemented by endpoints
// of network layer protocols (e.g. ipv4, ipv6).
type NetworkEndpoint interface {
	// MTU returns the maximum transmission unit for this endpoint. This is
	// generally calculated as the MTU of the underlying data link endpoint
	// minus the network endpoint max header length.
	MTU() uint32

	// MaxHeaderLength returns the maximum size the network (and lower
	// level layers combined) headers can have. Higher levels use this
	// information to reserve space in the front of the packets they're
	// building.
	MaxHeaderLength() uint16

	// ID returns the network protocol endpoint ID.
	ID() *NetworkEndpointID

	// NICID returns the id of the NIC this endpoint belongs to.
	NICID() tcpip.NICID

	// WritePacket writes a packet to the given destination address and
	// protocol.
	// TODO
	WritePacket(r *Route, hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error

	// HandlePacket is the interface that needs to be implemented by network
	// Protocols (e.g., IPv4, ipv6) that want to be part of the networking stack.
	HandlePacket(r *Route, v buffer.View)
}
