package ilist

import (
	"testing"
)

func TestResetEmpty(t *testing.T) {
	lst := &List{}
	a := &Entry{}
	lst.PushFront(a)

	if want, got := false, lst.Empty(); want != got {
		t.Fatalf("TestReset lst not empty:\n- want: %v\n- got: %v", want, got)
	}

	lst.Reset()
	if want, got := true, lst.Empty(); want != got {
		t.Fatalf("TestReset lst empty:\n- want: %v\n- got: %v", want, got)
	}
}

func TestFrontPushFront(t *testing.T) {
	lst := &List{}
	e := &Entry{}
	lst.PushFront(e)

	if want, got := lst.Front(), e; want != got {
		t.Fatalf("front error:\n- want: %v\n- got: %v", want, got)
	}
}

func TestBackPushBack(t *testing.T) {
	lst := &List{}
	b := &Entry{}
	lst.PushBack(b)

	if want, got := lst.Back(), b; want != got {
		t.Fatalf("back error:\n- want: %v\n- got: %v", want, got)
	}
}

func TestPushBackList(t *testing.T) {
	lst := &List{}

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

func TestInsertAfter(t *testing.T) {
	lst := &List{}

	a := &Entry{}
	b := &Entry{}
	c := &Entry{}
	lst.PushBack(a)
	lst.PushBack(c)

	lst.InsertAfter(a, b)
	lst.Remove(c)

	if want, got := lst.Back(), b; want != got {
		t.Fatalf("TestInsertAfter error:\n- want: %v\n- got: %v", want, got)
	}
}

func TestInsertBefore(t *testing.T) {
	lst := &List{}

	a := &Entry{}
	b := &Entry{}
	c := &Entry{}
	lst.PushBack(a)
	lst.PushBack(c)

	lst.InsertBefore(c, b)
	lst.Remove(a)

	if want, got := lst.Front(), b; want != got {
		t.Fatalf("TestInsertBefore error:\n- want: %v\n- got: %v", want, got)
	}
}
