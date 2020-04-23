package waiter

// Waitable contains the methods that need to be implemented by waitable
// objects.
type Waitable interface {
    // Readiness returns what the object is currently ready for. If it's
    // not ready for a desired purpose, the caller may use EventRegister and
    // EventUnregister to get notifications once the object becomes ready.
    Readiness(mask EventMask) EventMask

    // EventRegister registers the given waiter entry to receive
    // notifications when an event occurs that makes the object ready for
    // at least one of the events in mask.
    EventRegister(e *Entry, mask EventMask)

    // EventUnregister unregisters a waiter entry previously registered with
    // EventRegister
    EventUnregister(e *Entry)
}
