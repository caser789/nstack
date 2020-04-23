nstack [![Build Status](https://travis-ci.org/caser789/nstack.svg?branch=master)](https://travis-ci.org/caser789/nstack)
[![GoDoc](https://godoc.org/github.com/caser789/nstack?status.svg)](https://godoc.org/github.com/caser789/nstack)
[![Go Report Card](https://goreportcard.com/badge/github.com/caser789/nstack)](https://goreportcard.com/report/github.com/caser789/nstack)
[![Coverage Status](https://coveralls.io/repos/caser789/nstack/badge.svg?branch=master)](https://coveralls.io/r/caser789/nstack?branch=master)
=====

## Waiter

![waiter](./waiter.png)

```
@startuml

title waiter

interface Waitable {
    +Readiness(EventMask) EventMask
    +EventRegister(*Entry, EventMask)
    +EventUnregister(*Entry)
}

class AlwaysReady {}

Waitable <|-- AlwaysReady

class ilist.Entry {}
class ilist.List {}
ilist.List o-- ilist.Entry

class Entry {
    +Context interface{}
    +Callback func(*Entry)
    -mask EvnetMask
    ilist.Entry
}
class Queue {
    -list ilist.List
    -mu sync.RWMutext
    +EventRegister(*Entry, EventMask)
    +EventUnregister(*Entry)
    +Notify(EventMask)
    +Events() EventMask
    +IsEmpty() bool
}

ilist.Entry <|-- Entry
ilist.List *-- Queue
Queue o-- Entry

@enduml
```
