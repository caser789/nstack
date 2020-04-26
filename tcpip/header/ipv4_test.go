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
