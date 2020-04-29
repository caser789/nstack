package stack

import (
	"github.com/caser789/nstack/tcpip"
	"github.com/caser789/nstack/tcpip/buffer"

	"sync"
	"testing"
)

func Test_referencedNetworkEndpointincRef(t *testing.T) {
	ep := stubNetworkEndpoint{}
	protocol := tcpip.NetworkProtocolNumber(1)
	nic := &NIC{}

	rne := newReferencedNetworkEndpoint(ep, protocol, nic)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			rne.incRef()
		}(&wg)
	}
	wg.Wait()

	if want, got := 101, int(rne.refs); want != got {
		t.Fatalf("Test_referencedNetworkEndpointincRef failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_referencedNetworkEndpointdecRef(t *testing.T) {
	ep := stubNetworkEndpoint{}
	protocol := tcpip.NetworkProtocolNumber(1)
	nic := &NIC{}

	rne := &referencedNetworkEndpoint{
		refs:           200,
		ep:             ep,
		nic:            nic,
		protocol:       protocol,
		holdsInsertRef: true,
	}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			rne.decRef()
		}(&wg)
	}
	wg.Wait()

	if want, got := 100, int(rne.refs); want != got {
		t.Fatalf("Test_referencedNetworkEndpointdecRef failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_referencedNetworkEndpointtryIncRefFalse(t *testing.T) {
	ep := stubNetworkEndpoint{}
	protocol := tcpip.NetworkProtocolNumber(1)
	nic := &NIC{}

	rne := &referencedNetworkEndpoint{
		refs:           0,
		ep:             ep,
		nic:            nic,
		protocol:       protocol,
		holdsInsertRef: true,
	}

	if want, got := false, rne.tryIncRef(); want != got {
		t.Fatalf("Test_referencedNetworkEndpointtryIncRefFalse failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_referencedNetworkEndpointtryIncRefTrue(t *testing.T) {
	ep := stubNetworkEndpoint{}
	protocol := tcpip.NetworkProtocolNumber(1)
	nic := &NIC{}

	rne := &referencedNetworkEndpoint{
		refs:           1,
		ep:             ep,
		nic:            nic,
		protocol:       protocol,
		holdsInsertRef: true,
	}

	if want, got := true, rne.tryIncRef(); want != got {
		t.Fatalf("Test_referencedNetworkEndpointtryIncRefTrue failed:\n- want: %v\n- got: %v", want, got)
	}
}

type stubNetworkEndpoint struct{}

func (stubNetworkEndpoint) MTU() uint32             { return uint32(0) }
func (stubNetworkEndpoint) MaxHeaderLength() uint16 { return uint16(0) }
func (stubNetworkEndpoint) ID() *NetworkEndpointID  { return nil }
func (stubNetworkEndpoint) NICID() tcpip.NICID      { return tcpip.NICID(0) }
func (stubNetworkEndpoint) WritePacket(r *Route, hdr *buffer.Prependable, payload buffer.View, protocol tcpip.TransportProtocolNumber) error {
	return nil
}
func (stubNetworkEndpoint) HandlePacket(r *Route, v buffer.View) {}
