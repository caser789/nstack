package buffer

// VectorisedView is a vectorised version of View using non contigous memory.
// It supports all the convenience methods supported by View.
type VectorisedView struct {
    views []View
    size int
}

// NewVectorisedView creates a new vectorised view from an already-allocated slice
// of View and sets its size
func NewVectorisedView(size int, views []View) *VectorisedView {
    return &VectorisedView{views: views, size: size}
}

// Size returns the size in bytes of the entire content stored in the vectorised view.
func (vv *VectorisedView) Size() int {
    return vv.size
}

func (vv *VectorisedView) First() View {
    if len(vv.views) == 0 {
        panic("vview is empty")
    }

    return vv.views[0]
}

func (vv *VectorisedView) RemoveFirst() {
    if len(vv.views) == 0 {
        return
    }

    vv.size -= len(vv.views[0])
    vv.views = vv.views[1:]
}
