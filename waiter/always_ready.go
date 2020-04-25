package waiter

// AlwaysReady implements the Waitable interface but is always ready. Embedding
// this struct into another struct makes it implement the boilerplate empty
// functions automatically.
type AlwaysReady struct{}

// Readiness always returns the input mask because this project is always ready.
func (*AlwaysReady) Readiness(mask EventMask) EventMask {
	return mask
}

// EventRegister doesn't do anything because this object doesn't need to issue
// notifications because its readiness never changes.
func (*AlwaysReady) EventRegister(*Entry, EventMask) {}

// EventUnregister doesn't do anything because this object doesn't need to issue
// notifications because its  Readiness never changes.
func (*AlwaysReady) EventUnregister(*Entry) {}
