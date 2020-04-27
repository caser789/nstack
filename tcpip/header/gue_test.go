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
