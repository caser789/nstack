package header

import (
	"reflect"
	"testing"

	"github.com/caser789/nstack/tcpip"
)

func TestIPVersion(t *testing.T) {
	var tests = []struct {
		b       []byte
		version int
	}{
		{
			b:       []byte{},
			version: -1,
		},
		{
			b:       []byte{byte(4 << 4)},
			version: 4,
		},
	}

	for _, test := range tests {
		v := IPVersion(test.b)

		if want, got := test.version, v; want != got {
			t.Fatalf("TestIPVersion failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestHeaderLength(t *testing.T) {
	var tests = []struct {
		b      IPv4
		length uint8
	}{
		{
			b:      IPv4([]byte{byte(5)}),
			length: 20,
		},
		{
			b:      IPv4([]byte{byte(15)}),
			length: 60,
		},
	}

	for _, test := range tests {

		v := test.b.HeaderLength()

		if want, got := test.length, v; want != got {
			t.Fatalf("TestHeaderLength failed:\n- want: %v\n- got: %v", want, got)
		}
	}

}

func TestID(t *testing.T) {
	var tests = []struct {
		b  IPv4
		id int
	}{
		{
			b:  IPv4([]byte{byte(5), byte(0), byte(0), byte(0), byte(1), byte(2), byte(3)}),
			id: 1<<8 + 2,
		},
	}

	for _, test := range tests {

		v := test.b.ID()

		if want, got := test.id, v; int(want) != int(got) {
			t.Fatalf("TestID failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestProtocol(t *testing.T) {
	var tests = []struct {
		b        IPv4
		protocol uint8
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(0),
				byte(1), byte(2), byte(3), byte(0),
				byte(0), byte(11),
			}),
			protocol: uint8(11),
		},
	}

	for _, test := range tests {

		v := test.b.Protocol()

		if want, got := test.protocol, v; int(want) != int(got) {
			t.Fatalf("TestProtocol failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestFlags(t *testing.T) {
	var tests = []struct {
		b     IPv4
		flags uint8
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(0),
				byte(1), byte(2), byte(3 << 5), byte(0),
			}),
			flags: 3,
		},
	}

	for _, test := range tests {

		v := test.b.Flags()

		if want, got := test.flags, v; int(want) != int(got) {
			t.Fatalf("TestFlags failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestTTL(t *testing.T) {
	var tests = []struct {
		b   IPv4
		ttl uint8
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(0),
				byte(1), byte(2), byte(3 << 5), byte(0),
				byte(3), byte(2), byte(3 << 5), byte(0),
			}),
			ttl: 3,
		},
	}

	for _, test := range tests {

		v := test.b.Flags()

		if want, got := test.ttl, v; int(want) != int(got) {
			t.Fatalf("TestTTL failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestFragmentOffset(t *testing.T) {
	var tests = []struct {
		b      IPv4
		offset uint16
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(3 << 5), byte(0),
			}),
			offset: 0b00011111 << 11,
		},
	}

	for _, test := range tests {

		v := test.b.FragmentOffset()

		if want, got := test.offset, v; int(want) != int(got) {
			t.Fatalf("TestFragmentOffset failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestTotalLength(t *testing.T) {
	var tests = []struct {
		b      IPv4
		length uint16
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0b10101010), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(3 << 5), byte(0),
			}),
			length: uint16(0b1010101000000000),
		},
	}

	for _, test := range tests {

		v := test.b.TotalLength()

		if want, got := test.length, v; int(want) != int(got) {
			t.Fatalf("TestTotalLength failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestChecksum(t *testing.T) {
	var tests = []struct {
		b        IPv4
		checksum uint16
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0b10101010), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(0b10111011), byte(0),
			}),
			checksum: uint16(0b1011101100000000),
		},
	}

	for _, test := range tests {

		v := test.b.Checksum()

		if want, got := test.checksum, v; int(want) != int(got) {
			t.Fatalf("TestChecksum failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestSourceAddress(t *testing.T) {
	var tests = []struct {
		b    IPv4
		addr tcpip.Address
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0b10101010), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(0), byte(0),
				byte(1), byte(2), byte(3), byte(4),
			}),
			addr: tcpip.Address([]byte{byte(1), byte(2), byte(3), byte(4)}),
		},
	}

	for _, test := range tests {

		v := test.b.SourceAddress()

		if want, got := test.addr, v; want != got {
			t.Fatalf("TestSourceAddress failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestDestinationAddress(t *testing.T) {
	var tests = []struct {
		b    IPv4
		addr tcpip.Address
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0b10101010), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(0), byte(0),
				byte(1), byte(2), byte(3), byte(4),
				byte(4), byte(3), byte(2), byte(1),
			}),
			addr: tcpip.Address([]byte{byte(4), byte(3), byte(2), byte(1)}),
		},
	}

	for _, test := range tests {

		v := test.b.DestinationAddress()

		if want, got := test.addr, v; want != got {
			t.Fatalf("TestDestinationAddress failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestTransportProtocolNumber(t *testing.T) {
	var tests = []struct {
		b        IPv4
		protocol tcpip.TransportProtocolNumber
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(0),
				byte(1), byte(2), byte(3), byte(0),
				byte(0), byte(11),
			}),
			protocol: tcpip.TransportProtocolNumber(uint8(11)),
		},
	}

	for _, test := range tests {

		v := test.b.TransportProtocol()

		if want, got := test.protocol, v; want != got {
			t.Fatalf("TestTransportProtocolNumber failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestPayloadLength(t *testing.T) {
	var tests = []struct {
		b      IPv4
		length uint16
	}{
		{
			b: IPv4([]byte{
				byte(15), byte(0), byte(0b10101010), byte(0),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(3 << 5), byte(0),
			}),
			length: uint16(0b1010101000000000) - uint16(15*4),
		},
	}

	for _, test := range tests {

		v := test.b.PayloadLength()

		if want, got := test.length, v; want != got {
			t.Fatalf("TestPayloadLength failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestPayload(t *testing.T) {
	var tests = []struct {
		b       IPv4
		payload []byte
	}{
		{
			b: IPv4([]byte{
				byte(5), byte(0), byte(0), byte(24),
				byte(1), byte(2), byte(0b00011111), byte(0),
				byte(3), byte(2), byte(0), byte(0),
				byte(3), byte(2), byte(0), byte(0),
				byte(3), byte(2), byte(0), byte(0),
				byte(1), byte(2), byte(3), byte(4),
			}),
			payload: []byte{byte(1), byte(2), byte(3), byte(4)},
		},
	}

	for _, test := range tests {

		v := test.b.Payload()

		if want, got := test.payload, v; !reflect.DeepEqual(want, got) {
			t.Fatalf("TestPayload failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestTOS(t *testing.T) {
	b := IPv4(make([]byte, 10))

	v := uint8(11)
	b.SetTOS(v, 0)
	w, _ := b.TOS()

	if want, got := v, w; want != got {
		t.Fatalf("TestTOS failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestSetTotalLength(t *testing.T) {
	b := IPv4(make([]byte, 10))

	n := uint16(13)
	b.SetTotalLength(n)
	m := b.TotalLength()

	if want, got := n, m; want != got {
		t.Fatalf("TestSetTotalLength failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestSetChecksum(t *testing.T) {
	b := IPv4(make([]byte, 20))

	n := uint16(13)
	b.SetChecksum(n)
	m := b.Checksum()

	if want, got := n, m; want != got {
		t.Fatalf("TestSetChecksum failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestSetFlagsFragmentOffset(t *testing.T) {
	b := IPv4(make([]byte, 20))

	flags := uint8(3)
	offset := uint16(161 << 3)

	b.SetFlagsFragmentOffset(flags, offset)

	if want, got := flags, b.Flags(); want != got {
		t.Fatalf("TestSetFlagsFragmentOffset flags failed:\n- want: %v\n- got: %v", want, got)
	}

	if want, got := offset, b.FragmentOffset(); want != got {
		t.Fatalf("TestSetFlagsFragmentOffset offset failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestSetSourceAddress(t *testing.T) {
	b := IPv4(make([]byte, 20))

	addr := tcpip.Address([]byte{byte(1), byte(3), byte(5), byte(7)})

	b.SetSourceAddress(addr)

	if want, got := addr, b.SourceAddress(); want != got {
		t.Fatalf("TestSetSourceAddress failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestSetDestinationAddress(t *testing.T) {
	b := IPv4(make([]byte, 20))

	addr := tcpip.Address([]byte{byte(1), byte(3), byte(5), byte(7)})

	b.SetDestinationAddress(addr)

	if want, got := addr, b.DestinationAddress(); want != got {
		t.Fatalf("TestSetDestinationAddress failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestEncode(t *testing.T) {
	var tests = []struct {
		b     IPv4
		field *IPv4Fields
		desc  string
	}{
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
			}, make([]byte, 50-1)...)),
			field: &IPv4Fields{
				IHL: 20,
			},
			desc: "test versIHL",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
			}, make([]byte, 50-2)...)),
			field: &IPv4Fields{
				IHL: 20,
				TOS: 21,
			},
			desc: "test tos",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
			}, make([]byte, 50-4)...)),
			field: &IPv4Fields{
				IHL:         20,
				TOS:         21,
				TotalLength: 22,
			},
			desc: "test total length",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
			}, make([]byte, 50-6)...)),
			field: &IPv4Fields{
				IHL:         20,
				TOS:         21,
				TotalLength: 22,
				ID:          23,
			},
			desc: "test ID",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
			}, make([]byte, 50-8)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
			},
			desc: "test flags and fragment",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
				25,
			}, make([]byte, 50-9)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
				TTL:            25,
			},
			desc: "test TTL",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
				25,
				26,
			}, make([]byte, 50-10)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
				TTL:            25,
				Protocol:       26,
			},
			desc: "test Protocol",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
				25,
				26,
				0,
				27,
			}, make([]byte, 50-12)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
				TTL:            25,
				Protocol:       26,
				Checksum:       27,
			},
			desc: "test Checksum",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
				25,
				26,
				0,
				27,
				1, 2, 3, 4,
			}, make([]byte, 50-16)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
				TTL:            25,
				Protocol:       26,
				Checksum:       27,
				SrcAddr:        tcpip.Address([]byte{1, 2, 3, 4}),
			},
			desc: "test SrcAddr",
		},
		{
			b: IPv4(append([]byte{
				(4 << 4) | (5 & 0xf),
				21,
				0, 22,
				0, 23,
				3 << 5,
				24,
				25,
				26,
				0,
				27,
				1, 2, 3, 4,
				4, 3, 2, 1,
			}, make([]byte, 50-20)...)),
			field: &IPv4Fields{
				IHL:            20,
				TOS:            21,
				TotalLength:    22,
				ID:             23,
				Flags:          3,
				FragmentOffset: 24 << 3,
				TTL:            25,
				Protocol:       26,
				Checksum:       27,
				SrcAddr:        tcpip.Address([]byte{1, 2, 3, 4}),
				DstAddr:        tcpip.Address([]byte{4, 3, 2, 1}),
			},
			desc: "test DstAddr",
		},
	}

	for _, test := range tests {
		bb := IPv4(make([]byte, 50))

		bb.Encode(test.field)

		if want, got := test.b, bb; !reflect.DeepEqual(want, got) {
			t.Fatalf("TestEncode %s failed:\n- want: %v\n- got: %v", test.desc, want, got)
		}

	}
}
