package buffer

import (
    "testing"
)

func TestTrimView(t *testing.T) {
    size := 10
    v := NewView(size)
    v[0] = 'a'
    v[1] = 'b'
    v[2] = 'c'

    v.TrimFront(1)

	if want, got := byte('b'), v[0]; want != got {
		t.Fatalf("TestTrimView failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestCapLength(t *testing.T) {
    size := 10
    v := NewView(size)
    v[0] = 'a'
    v[1] = 'b'
    v[2] = 'c'

	if want, got := 10, len(v); want != got {
		t.Fatalf("TestCapLength before cap failed:\n- want: %v\n- got: %v", want, got)
	}

    v.CapLength(3)

	if want, got := 3, len(v); want != got {
		t.Fatalf("TestCapLength after cap failed:\n- want: %v\n- got: %v", want, got)
	}
}
