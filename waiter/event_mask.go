package waiter

import (
    "syscall"
)

// EventMask represents io events as used in the poll() syscall.
type EventMask uint16

// Events that waiters can wait on. The meaning is the same as those in the
// poll() syscall.
const (
    EventIn   EventMask = syscall.EPOLLIN
    EventPri  EventMask = syscall.EPOLLPRI
    EventOut  EventMask = syscall.EPOLLOUT
    EventErr  EventMask = syscall.EPOLLERR
    EventHup  EventMask = syscall.EPOLLHUP
    EventNVal EventMask = 0x20 // Not defined in syscall.
)
