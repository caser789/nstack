package buffer

import (
    "testing"
    "reflect"
)

func TestSize(t *testing.T) {
    size := 11
    v := NewVectorisedView(size, nil)

	if want, got := size, v.Size(); want != got {
		t.Fatalf("TestSize failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestFirst(t *testing.T) {
    size := 11
    views := make([]View, 2)
    views[0] = NewView(size)

    v := NewVectorisedView(size, views)

	if want, got := views[0], v.First(); !reflect.DeepEqual(want, got) {
		t.Fatalf("TestFirst failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestRemoveFirstEmpty(t *testing.T) {
    size := 11
    v := NewVectorisedView(size, nil)

    v.RemoveFirst()

	if want, got := size, v.Size(); want != got {
		t.Fatalf("TestRemoveFirstEmpty failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestRemoveFirstNotEmpty(t *testing.T) {
    size := 11
    views := make([]View, 2)
    views[0] = NewView(size)
    v := NewVectorisedView(size, views)

    v.RemoveFirst()

	if want, got := 0, v.Size(); want != got {
		t.Fatalf("TestRemoveFirstNotEmpty failed:\n- want: %v\n- got: %v", want, got)
	}
}
