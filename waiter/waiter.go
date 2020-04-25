package waiter

import (
	"sync"

	"github.com/caser789/nstack/ilist"
)

// Queue represents the wait queue where waiters can be added and
// notifiers can notify them when evnets happen.
//
// The zero value for waiter.Queue is an empty queue reaqdy for use.
type Queue struct {
	list ilist.List
	mu   sync.RWMutex
}

// EventRegister adds a waiter to the waiter queue; the waiter will be notified
// when at least one of the events specified in mask happens.
func (q *Queue) EventRegister(e *Entry, mask EventMask) {
	q.mu.Lock()
	e.mask = mask
	q.list.PushBack(e)
	q.mu.Unlock()
}

// EventUnregister removes the given waiter entry from the waiter queue.
func (q *Queue) EventUnregister(e *Entry) {
	q.mu.Lock()
	q.list.Remove(e)
	q.mu.Unlock()
}

// Notify notifies all waiters in the queue whose masks have at least one bit
// in common with the notification mask.
func (q *Queue) Notify(mask EventMask) {
	q.mu.RLock()
	for it := q.list.Front(); it != nil; it = it.Next() {
		e := it.(*Entry)
		if (mask & e.mask) != 0 {
			e.Callback(e)
		}
	}
	q.mu.RUnlock()
}

// Events returns the set of events being waited on. It is the union of the
// masks of all registered entries.
func (q *Queue) Events() EventMask {
	ret := EventMask(0)

	q.mu.RLock()
	for it := q.list.Front(); it != nil; it = it.Next() {
		e := it.(*Entry)
		ret |= e.mask
	}
	q.mu.RUnlock()

	return ret
}

// IsEmpty returns if the waiter queue is empty or not.
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.list.Front() == nil
}
