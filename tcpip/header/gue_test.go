package header

import (
	"testing"
)

func TestTypeAndControl(t *testing.T) {
	var tests = []struct {
		b              GUE
		typeAndControl uint8
	}{
		{
			b:              GUE([]byte{0b11100000}),
			typeAndControl: uint8(0b00000111),
		},
	}

	for _, test := range tests {

		if want, got := test.typeAndControl, test.b.TypeAndControl(); want != got {
			t.Fatalf("TestTypeAndControl failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestGUEHeaderLength(t *testing.T) {
	var tests = []struct {
		b      GUE
		length uint8
	}{
		{
			b:      GUE([]byte{0b11101000}),
			length: uint8(0b00001000*4 + 4),
		},
	}

	for _, test := range tests {

		if want, got := test.length, test.b.HeaderLength(); want != got {
			t.Fatalf("TestHeaderLength failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}
