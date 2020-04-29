package stack

import (
	"sync/atomic"

	"github.com/caser789/nstack/ilist"
	"github.com/caser789/nstack/tcpip"
)

type referencedNetworkEndpoint struct {
	ilist.Entry

	refs     int32
	ep       NetworkEndpoint
	nic      *NIC
	protocol tcpip.NetworkProtocolNumber

	// holdsInsertRef is protected by the NIC's mutex. It indicates whether
	// the reference count is biased by 1 due to the insertion of the
	// endpoint. It is reset to false when RemoveAddress is called on the
	// NIC.
	holdsInsertRef bool
}

func newReferencedNetworkEndpoint(ep NetworkEndpoint, protocol tcpip.NetworkProtocolNumber, nic *NIC) *referencedNetworkEndpoint {
	return &referencedNetworkEndpoint{
		refs:           1,
		ep:             ep,
		nic:            nic,
		protocol:       protocol,
		holdsInsertRef: true,
	}
}

func (r *referencedNetworkEndpoint) decRef() {
	if atomic.AddInt32(&r.refs, -1) == 0 {
		r.nic.removeEndpoint(r)
	}
}

func (r *referencedNetworkEndpoint) incRef() {
	atomic.AddInt32(&r.refs, 1)
}

func (r *referencedNetworkEndpoint) tryIncRef() bool {
	for {
		v := atomic.LoadInt32(&r.refs)
		if v == 0 {
			return false
		}

		if atomic.CompareAndSwapInt32(&r.refs, v, v+1) {
			return true
		}
	}
}
