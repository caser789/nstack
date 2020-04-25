package buffer

// View is a slice of a buffer, with convenience methods
type View []byte

// NewView allocates a new buffer and returns an initialized view that covers
// the whole buffer
func NewView(size int) View {
    return make(View, size)
}

// NewView allocates a new buffer and returns an initialized view that covers
// the whole buffer
func (v *View) TrimFront(count int) {
    *v = (*v)[count:]
}
