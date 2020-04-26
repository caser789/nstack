package header

import (
	"testing"
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
