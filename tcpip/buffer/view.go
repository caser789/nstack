package buffer

// View is a slice of a buffer, with convenience methods
type View []byte

// NewView allocates a new buffer and returns an initialized view that covers
// the whole buffer
func NewView(size int) View {
	return make(View, size)
}

// TrimFront trims the view fron front
func (v *View) TrimFront(count int) {
	*v = (*v)[count:]
}

// CapLength irreversibly reduces the length of the visible section of the
// buffer to the value specified
func (v *View) CapLength(length int) {
	// We also set the slice cap because if we don't, one would be able to
	// expand the view back to include the region just excluded. We want to
	// prevent that to avoid potential data leak if we have uninitialized
	// data in excluded region
	*v = (*v)[:length:length]
}
