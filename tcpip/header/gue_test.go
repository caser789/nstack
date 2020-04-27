package header

import (
	"reflect"
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

func TestGUEProtocol(t *testing.T) {
	var tests = []struct {
		b     GUE
		proto uint8
	}{
		{
			b:     GUE([]byte{0b11101000, 11}),
			proto: uint8(11),
		},
	}

	for _, test := range tests {
		if want, got := test.proto, test.b.Protocol(); want != got {
			t.Fatalf("TestGUEProtocol failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}

func TestGUEEncode(t *testing.T) {
	var tests = []struct {
		b      GUE
		fields *GUEFields
	}{
		{
			b: GUE([]byte{
				0b11100100,
				77,
			}),
			fields: &GUEFields{
				Type:         uint8(3),
				Control:      true,
				HeaderLength: uint8(20),
				Protocol:     uint8(77),
			},
		},
	}

	for _, test := range tests {
		bb := GUE(make([]byte, 2))

		bb.Encode(test.fields)

		if want, got := test.b, bb; !reflect.DeepEqual(want, got) {
			t.Fatalf("TestGUEEncode failed:\n- want: %v\n- got: %v", want, got)
		}
	}
}
