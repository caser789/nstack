package header

import (
    "testing"
)

func TestIPVersion(t *testing.T) {
    var tests = []struct{
        b []byte
        expected int
    }{
        {
            b: []byte{},
            expected: -1,
        },
        {
            b: []byte{byte(4<<4)},
            expected: 4,
        },
    }

    for _, test := range tests {
        v := IPVersion(test.b)

        if want, got := test.expected, v; want != got {
            t.Fatalf("TestIPVersion failed:\n- want: %v\n- got: %v", want, got)
        }
    }
}

