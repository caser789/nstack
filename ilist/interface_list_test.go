package ilist

import (
    "testing"
)

func TestIList(t *testing.T) {
    lst := &List{}

    if want, got := true, lst.Empty(); want != got {
        t.Fatalf("lst not empty:\n- want: %v\n- got: %v", want, got)
    }

    e := &Entry{}
    lst.PushFront(e)

    if want, got := false, lst.Empty(); want != got {
        t.Fatalf("lst empty:\n- want: %v\n- got: %v", want, got)
    }

    if want, got := lst.Front(), e; want != got {
        t.Fatalf("front error:\n- want: %v\n- got: %v", want, got)
    }

    b := &Entry{}
    lst.PushBack(b)

    if want, got := lst.Back(), b; want != got {
        t.Fatalf("back error:\n- want: %v\n- got: %v", want, got)
    }

    lst2 := &List{}
    c := &Entry{}
    d := &Entry{}
    lst2.PushBack(c)
    lst2.PushBack(d)
    lst.PushBackList(lst2)

    if want, got := lst.Back(), d; want != got {
        t.Fatalf("PushBackList error:\n- want: %v\n- got: %v", want, got)
    }
}
