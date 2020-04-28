package stack

import (
	"reflect"
	"testing"

	"github.com/caser789/nstack/tcpip"
	"github.com/caser789/nstack/tcpip/buffer"
)

type StubTransportEndpoint struct{}

func (s StubTransportEndpoint) HandlePacket(r *Route, id TransportEndpointID, v buffer.View) {}

func Test_registerEndpointUnknownProtocol(t *testing.T) {
	stack := &Stack{}
	mux := newTransportDemuxer(stack)

	protocol := tcpip.TransportProtocolNumber(1)
	id := TransportEndpointID{}
	ep := StubTransportEndpoint{}

	err := mux.registerEndpoint(protocol, id, ep)

	if want, got := tcpip.ErrUnknownProtocol, err; want != got {
		t.Fatalf("Test_registerEndpointUnknownProtocol failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_registerEndpointDuplicateAddress(t *testing.T) {
	stack := &Stack{}
	mux := newTransportDemuxer(stack)

	protocol := tcpip.TransportProtocolNumber(1)
	id := TransportEndpointID{}
	ep := StubTransportEndpoint{}

	eps := map[TransportEndpointID]TransportEndpoint{
		id: ep,
	}
	mux.protocol[protocol] = &transportEndpoints{endpoints: eps}

	err := mux.registerEndpoint(protocol, id, ep)

	if want, got := tcpip.ErrDuplicateAddress, err; want != got {
		t.Fatalf("Test_registerEndpointDuplicateAddress failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_registerEndpointOK(t *testing.T) {
	stack := &Stack{}
	mux := newTransportDemuxer(stack)

	protocol := tcpip.TransportProtocolNumber(1)
	id := TransportEndpointID{}
	ep := StubTransportEndpoint{}

	mux.protocol[protocol] = &transportEndpoints{endpoints: make(map[TransportEndpointID]TransportEndpoint)}

	err := mux.registerEndpoint(protocol, id, ep)

	if err != nil {
		t.Fatalf("Test_registerEndpointOK failed\n")
	}
}

func Test_unregisterEndpointNewProtocolNewProtocol(t *testing.T) {
	stack := &Stack{}
	mux := newTransportDemuxer(stack)

	protocol := tcpip.TransportProtocolNumber(1)
	id := TransportEndpointID{}

	mux.unregisterEndpoint(protocol, id)

	want := make(map[tcpip.TransportProtocolNumber]*transportEndpoints)
	got := mux.protocol

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Test_unregisterEndpointNewProtocolNewProtocol failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_unregisterEndpointNewProtocol(t *testing.T) {
	stack := &Stack{}
	mux := newTransportDemuxer(stack)

	protocol := tcpip.TransportProtocolNumber(1)
	id := TransportEndpointID{}
	ep := StubTransportEndpoint{}

	eps := map[TransportEndpointID]TransportEndpoint{
		id: ep,
	}
	mux.protocol[protocol] = &transportEndpoints{endpoints: eps}

	mux.unregisterEndpoint(protocol, id)

	want := map[tcpip.TransportProtocolNumber]*transportEndpoints{
		protocol: {endpoints: make(map[TransportEndpointID]TransportEndpoint)},
	}
	got := mux.protocol

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Test_unregisterEndpointNewProtocol failed:\n- want: %v\n- got: %v", want, got)
	}
}

func Test_deliverPacket(t *testing.T) {
	var tests = []struct {
		d     *transportDemuxer
		desc  string
		proto tcpip.TransportProtocolNumber
		id    TransportEndpointID
		res   bool
	}{
		{
			desc:  "unknown protocol",
			proto: tcpip.TransportProtocolNumber(1),
			id:    TransportEndpointID{},
			res:   false,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					2: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "match by id",
			proto: tcpip.TransportProtocolNumber(1),
			id:    TransportEndpointID{},
			res:   true,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "match by id",
			proto: tcpip.TransportProtocolNumber(1),
			id: TransportEndpointID{
				LocalPort:     uint16(1),
				LocalAddress:  tcpip.Address("a"),
				RemotePort:    uint16(2),
				RemoteAddress: tcpip.Address("b"),
			},
			res: true,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{
								LocalPort:     uint16(1),
								LocalAddress:  tcpip.Address("a"),
								RemotePort:    uint16(2),
								RemoteAddress: tcpip.Address("b"),
							}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "match by id minus local address",
			proto: tcpip.TransportProtocolNumber(1),
			id: TransportEndpointID{
				LocalPort:     uint16(1),
				LocalAddress:  tcpip.Address("a"),
				RemotePort:    uint16(2),
				RemoteAddress: tcpip.Address("b"),
			},
			res: true,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{
								LocalPort:     uint16(1),
								LocalAddress:  tcpip.Address(""),
								RemotePort:    uint16(2),
								RemoteAddress: tcpip.Address("b"),
							}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "match by id minus remote part",
			proto: tcpip.TransportProtocolNumber(1),
			id: TransportEndpointID{
				LocalPort:     uint16(1),
				LocalAddress:  tcpip.Address("a"),
				RemotePort:    uint16(2),
				RemoteAddress: tcpip.Address("b"),
			},
			res: true,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{
								LocalPort:     uint16(1),
								LocalAddress:  tcpip.Address("a"),
								RemotePort:    uint16(0),
								RemoteAddress: tcpip.Address(""),
							}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "match by local port",
			proto: tcpip.TransportProtocolNumber(1),
			id: TransportEndpointID{
				LocalPort:     uint16(1),
				LocalAddress:  tcpip.Address("a"),
				RemotePort:    uint16(2),
				RemoteAddress: tcpip.Address("b"),
			},
			res: true,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{
								LocalPort:     uint16(1),
								LocalAddress:  tcpip.Address(""),
								RemotePort:    uint16(0),
								RemoteAddress: tcpip.Address(""),
							}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
		{
			desc:  "not match",
			proto: tcpip.TransportProtocolNumber(1),
			id: TransportEndpointID{
				LocalPort:     uint16(1),
				LocalAddress:  tcpip.Address("a"),
				RemotePort:    uint16(2),
				RemoteAddress: tcpip.Address("b"),
			},
			res: false,
			d: &transportDemuxer{
				protocol: map[tcpip.TransportProtocolNumber]*transportEndpoints{
					1: {
						endpoints: map[TransportEndpointID]TransportEndpoint{
							{
								LocalPort:     uint16(2),
								LocalAddress:  tcpip.Address(""),
								RemotePort:    uint16(0),
								RemoteAddress: tcpip.Address(""),
							}: StubTransportEndpoint{},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		r := &Route{}
		v := make(buffer.View, 0)

		want := tt.res
		got := tt.d.deliverPacket(r, tt.proto, v, tt.id)

		if want != got {
			t.Fatalf("Test_deliverPacket %s failed:\n- want: %v\n- got: %v", tt.desc, want, got)
		}
	}
}
