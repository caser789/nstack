package waiter

import (
    "github.com/caser789/nstack/ilist"
)

// Entry represents a waiter that can be added to a waiter queue. It can
// only be in at one queue at a time, and is added "intrusively" to the queue with
// no extra memory allocations.
type Entry struct {
    // Context stores any state the waiter may wish to store in the entry
    // itself, which can be used at wake up time.
    Context interface{}

    // Callback is the function to be called when the waiter entry is
    // notified. It is responsible for doing whatever is needed to wake up
    // the waiter.
    //
    // The callback is supposed to perform minimum work, and cannot call
    // any method on the queue itself because it will be locked while the callback
    // is running
    Callback func(e *Entry)

    // The following fields are protected by the queue lock.
    mask EventMask
    ilist.Entry
}

// NewChannelEntry initializes a new Entry that does a non-blocking write of nil
// to an interface{} channel when the callback is called. It returns the new
// Entry instance and the channeld being used.
//
// If a channel isn't specified (i.e., if "c" is nil), then NewChannelEntry
// allocates a new channel.
func NewChannelEntry(c chan interface{}) (Entry, chan interface{}) {
    if c == nil {
        // TODO: consider a pool.
        c = make(chan interface{}, 1)
    }

    return Entry{
        Context: c,
        Callback: func(e *Entry) {
            ch := e.Context.(chan interface{})
            select {
            case ch <- nil:
            default:
            }
        },
    }, c
}
