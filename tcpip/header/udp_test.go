package header

import (
	"reflect"
	"testing"
)

func TestUDPSourcePort(t *testing.T) {
	var tests = []struct {
		b    UDP
		port uint16
	}{
		{
			b:    UDP([]byte{0b10101010, 0b01010101, 0b11111111}),
			port: uint16(0b1010101001010101),
		},
	}

	for _, test := range tests {

		if want, got := test.port, test.b.SourcePort(); want != got {
			t.Fatalf("TestUDPSourcePort failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestUDPDestinationPort(t *testing.T) {
	var tests = []struct {
		b    UDP
		port uint16
	}{
		{
			b: UDP([]byte{
				0b10101010, 0b01010101,
				0b01010101, 0b10101010,
				0b11111111,
			}),
			port: uint16(0b0101010110101010),
		},
	}

	for _, test := range tests {

		if want, got := test.port, test.b.DestinationPort(); want != got {
			t.Fatalf("TestUDPDestinationPort failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestUDPLength(t *testing.T) {
	var tests = []struct {
		b      UDP
		length uint16
	}{
		{
			b: UDP([]byte{
				0b10101010, 0b01010101,
				0b01010101, 0b10101010,
				0b10010010, 0b01101101,
				0b11111111,
			}),
			length: uint16(0b1001001001101101),
		},
	}

	for _, test := range tests {

		if want, got := test.length, test.b.Length(); want != got {
			t.Fatalf("TestUDPLength failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestUDPChecksum(t *testing.T) {
	var tests = []struct {
		b        UDP
		checksum uint16
	}{
		{
			b: UDP([]byte{
				0b10101010, 0b01010101,
				0b01010101, 0b10101010,
				0b10010010, 0b01101101,
				0b11011011, 0b00100100,
				0b11111111,
			}),
			checksum: uint16(0b1101101100100100),
		},
	}

	for _, test := range tests {

		if want, got := test.checksum, test.b.Checksum(); want != got {
			t.Fatalf("TestUDPChecksum failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestUDPPayload(t *testing.T) {
	var tests = []struct {
		b       UDP
		payload []byte
	}{
		{
			b: UDP([]byte{
				0b10101010, 0b01010101,
				0b01010101, 0b10101010,
				0b10010010, 0b01101101,
				0b11011011, 0b00100100,
				0b11111111, 0b00000001,
			}),
			payload: []byte{0b11111111, 0b00000001},
		},
	}

	for _, test := range tests {

		if want, got := test.payload, test.b.Payload(); !reflect.DeepEqual(want, got) {
			t.Fatalf("TestUDPPayload failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestUDPSetSourcePort(t *testing.T) {
	port := uint16(11)
	b := UDP(make([]byte, 8))
	b.SetSourcePort(port)

	if want, got := port, b.SourcePort(); want != got {
		t.Fatalf("TestSetSOurcePort failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestUDPSetDestinationPort(t *testing.T) {
	port := uint16(11)
	b := UDP(make([]byte, 8))
	b.SetDestinationPort(port)

	if want, got := port, b.DestinationPort(); want != got {
		t.Fatalf("TestSetDestinationPort failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestUDPSetChecksum(t *testing.T) {
	checksum := uint16(11)
	b := UDP(make([]byte, 8))
	b.SetChecksum(checksum)

	if want, got := checksum, b.Checksum(); want != got {
		t.Fatalf("TestSetChecksum failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestUDPEncode(t *testing.T) {
	var tests = []struct {
		fields *UDPFields
		b      UDP
		bb     UDP
		desc   string
	}{
		{
			fields: &UDPFields{
				SrcPort: 0b0101010110101010,
			},
			b: UDP(append(
				[]byte{
					0b01010101, 0b10101010,
				},
				make([]byte, 8-2)...,
			)),
			desc: "test SrcPort",
			bb:   UDP(make([]byte, 8)),
		},
		{
			fields: &UDPFields{
				SrcPort: 0b0101010110101010,
				DstPort: 0b1010101001010101,
			},
			b: UDP(append(
				[]byte{
					0b01010101, 0b10101010,
					0b10101010, 0b01010101,
				},
				make([]byte, 8-4)...,
			)),
			desc: "test DstPort",
			bb:   UDP(make([]byte, 8)),
		},
		{
			fields: &UDPFields{
				SrcPort: 0b0101010110101010,
				DstPort: 0b1010101001010101,
				Length:  0b1001001001101101,
			},
			b: UDP(append(
				[]byte{
					0b01010101, 0b10101010,
					0b10101010, 0b01010101,
					0b10010010, 0b01101101,
				},
				make([]byte, 8-6)...,
			)),
			desc: "test Length",
			bb:   UDP(make([]byte, 8)),
		},
		{
			fields: &UDPFields{
				SrcPort:  0b0101010110101010,
				DstPort:  0b1010101001010101,
				Length:   0b1001001001101101,
				Checksum: 0b0110110111011010,
			},
			b: UDP(append(
				[]byte{
					0b01010101, 0b10101010,
					0b10101010, 0b01010101,
					0b10010010, 0b01101101,
					0b01101101, 0b11011010,
				},
				make([]byte, 8-8)...,
			)),
			desc: "test Checksum",
			bb:   UDP(make([]byte, 8)),
		},
	}

	for _, test := range tests {

		test.bb.Encode(test.fields)

		if want, got := test.b, test.bb; !reflect.DeepEqual(want, got) {
			t.Fatalf("TestUDPEncode %s failed:\n- want: %v\n- got: %v", test.desc, want, got)
		}
	}
}
