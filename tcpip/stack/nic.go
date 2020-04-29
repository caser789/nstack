package stack

import (
	"sync"

	"github.com/caser789/nstack/ilist"
	"github.com/caser789/nstack/tcpip"
)

type NIC struct {
	mu        sync.RWMutex
	primary   map[tcpip.NetworkProtocolNumber]*ilist.List
	endpoints map[NetworkEndpointID]*referencedNetworkEndpoint
}

func (n *NIC) removeEndpointLocked(r *referencedNetworkEndpoint) {
	id := *r.ep.ID()

	// Nothing to do if the reference has already been replaced with a
	// different one.
	if n.endpoints[id] != r {
		return
	}

	if r.holdsInsertRef {
		panic("Reference count dropped to zero before being removed")
	}

	delete(n.endpoints, id)
	n.primary[r.protocol].Remove(r)
}

func (n *NIC) removeEndpoint(r *referencedNetworkEndpoint) {
	n.mu.Lock()
	n.removeEndpointLocked(r)
	n.mu.Unlock()
}
