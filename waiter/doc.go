// Package waiter provides the implementation of a wallet queue, where waiters can
// be enqueued to be notified when an event of interest happends.
//
// Becoming readable and/or writable are examples of events. Waiters are expected
// to use a pattern similar to this to make a blocking function out of
// a non-blocking one:
//
// func (o *object) blockingRead(...) error {
//     err := o.nonblockingRead(...)
//     if err != ErrAgain {
//         // Completed with no need to wait!
//         return err
//     }
//
//     e := createOrGetWaiterEntry(...)
//     o.EventRegister(&e, waiter.EventIn)
//     defer o.EventUnregister(&e)
//
//     // We need to try to read again after registration because the
//     // object may have become readable between the last attempt to
//     // read and read registration.
//     err := o.nonblockingRead(...)
//     if err == ErrAgain {
//         wait()
//         err = o.nonblockingRead(...)
//     }
//
//     return err
//
// }
// Another goroutine needs to notify waiters when events happen. For example:
//
// func(o *object) Write(...) ... {
//     // Do write work.
//     [...]
//
//     if oldDataAvailableSize == 0 && dataAvailableSize > 0 {
//         // If no data was available and now some data is
//         // available, the object became readable, so notify
//         // potential waiters about this
//         o.Notify(waiter.EventIn)
//     }
//
// }
package waiter
