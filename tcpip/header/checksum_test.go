package header

import (
	"testing"
)

func TestChecksumCombine(t *testing.T) {
	var tests = []struct {
		a        uint16
		b        uint16
		checksum uint16
	}{
		{
			a:        uint16(0b1111111111111111),
			b:        uint16(0b1111111111111111),
			checksum: uint16(0b1111111111111111),
		},
	}

	for _, test := range tests {

		v := ChecksumCombine(test.a, test.b)

		if want, got := test.checksum, v; want != got {
			t.Fatalf("TestChecksumCombine failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}
