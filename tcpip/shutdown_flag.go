package tcpip

// ShutdownFlags represents flags that can be passed to the Shutdown() method
// of the Endpoint interface
type ShutdownFlags int

// Values of the flags that can be passed to the Shutdown() method. They can
// be OR'ed together
const (
	ShutdownRead ShutdownFlags = 1 << iota
	ShutdownWrite
)
