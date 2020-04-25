package tcpip

import (
	"github.com/caser789/nstack/waiter"
)

// Stack represents a networking stack, with all supported protocols, NICs, and
// route table.
type Stack interface {
	// NewEndpoint creates a new transport layer endpoint of the given
	// protocol.
	NewEndpoint(transport TransportProtocolNumber, network NetworkProtocolNumber, waiterQueue *waiter.Queue) (Endpoint, error)

	// SetRouteTable assigns the route table to be used by this stack. It
	// specifies which NICs to use for given destination address ranges.
	SetRouteTable(table []Route)

	// CreateNIC creates a NIC with the provided id and link-layer sender.
	CreateNIC(id NICID, linkEndpoint LinkEndpointID) error

	// AddAddress adds a new network-layer address to the specified NIC.
	AddAddress(id NICID, protocol NetworkProtocolNumber, addr Address) error

	// Stats returns a snapshot of the current stats.
	Stats() Stats
}
