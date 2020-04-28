package buffer

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	var tests = []struct {
		desc        string
		prependable *Prependable
		size        int
		res         []byte
	}{
		{
			desc: "test smaller size",
			prependable: &Prependable{
				usedIdx: 4,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			size: 5,
			res:  nil,
		},
		{
			desc: "test ok",
			prependable: &Prependable{
				usedIdx: 4,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			size: 2,
			res:  []byte{3, 4},
		},
	}

	for _, tt := range tests {
		if want, got := tt.res, tt.prependable.Prepend(tt.size); !reflect.DeepEqual(want, got) {
			t.Fatalf("TestPrepend %s failed:\n- want: %v\n- got: %v", tt.desc, want, got)
		}
	}
}

func TestPrependableView(t *testing.T) {
	var tests = []struct {
		desc        string
		prependable *Prependable
		res         []byte
	}{
		{
			desc: "test used",
			prependable: &Prependable{
				usedIdx: 2,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			res: []byte{3, 4},
		},
		{
			desc: "test not used",
			prependable: &Prependable{
				usedIdx: 4,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			res: []byte{},
		},
	}

	for _, tt := range tests {
		if want, got := View(tt.res), tt.prependable.View(); !reflect.DeepEqual(want, got) {
			t.Fatalf("TestPrependableView %s failed:\n- want: %v\n- got: %v", tt.desc, want, got)
		}
	}
}

func TestPrependableUsedBytes(t *testing.T) {
	var tests = []struct {
		desc        string
		prependable *Prependable
		res         []byte
	}{
		{
			desc: "test used",
			prependable: &Prependable{
				usedIdx: 2,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			res: []byte{3, 4},
		},
		{
			desc: "test not used",
			prependable: &Prependable{
				usedIdx: 4,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			res: []byte{},
		},
	}

	for _, tt := range tests {
		if want, got := tt.res, tt.prependable.UsedBytes(); !reflect.DeepEqual(want, got) {
			t.Fatalf("TestPrependableUsedBytes %s failed:\n- want: %v\n- got: %v", tt.desc, want, got)
		}
	}
}

func TestPrependableUsedLength(t *testing.T) {
	var tests = []struct {
		desc        string
		prependable *Prependable
		length      int
	}{
		{
			desc: "test used",
			prependable: &Prependable{
				usedIdx: 2,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			length: 2,
		},
		{
			desc: "test not used",
			prependable: &Prependable{
				usedIdx: 4,
				buf: View([]byte{
					1, 2, 3, 4,
				}),
			},
			length: 0,
		},
	}

	for _, tt := range tests {
		if want, got := tt.length, tt.prependable.UsedLength(); want != got {
			t.Fatalf("TestPrependableUsedLength %s failed:\n- want: %v\n- got: %v", tt.desc, want, got)
		}
	}
}
