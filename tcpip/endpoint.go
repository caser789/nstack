package tcpip

import (
	"github.com/caser789/nstack/tcpip/buffer"
	"github.com/caser789/nstack/waiter"
)

// Endpoint is the interface implemented by transport protocols (e.g., tcp, udp)
// that exposes functionality like read, write, connect, etc. to users of the
// networking stack.
type Endpoint interface {
	// Close puts the endpoint in a closed state and frees all resources
	// associated with it.
	Close()

	// Read reads data from the endpoint and optionally returns the sender.
	// This method does not block if there is no data pending.
	// It will also either return an error or data, never both.
	Read(*FullAddress) (buffer.View, error)

	// Write writes data to the endpoint's peer, or the provided address if
	// one is specified. This method does not block if the data cannot be
	// written.
	Write(buffer.View, *FullAddress) (uintptr, error)

	// RecvMsg reads data and a control message from the endpoint. This method
	// does not block if there is no data pending.
	RecvMsg(*FullAddress) (buffer.View, ControlMessages, error)

	// SendMsg writes data and a control message to the endpoint's peer.
	// This method does not block if the data cannot be written.
	//
	// SendMsg does not take ownership of any of its arguments on error.
	SendMsg(buffer.View, ControlMessages, *FullAddress) (uintptr, error)

	// Peek reads data without consuming it from the endpoint.
	//
	// This method does not block if there is no data pending.
	Peek(io.Writer) (uintptr, error)

	// Connect connects the endpoint to its peer. Specifying a NIC is
	// optional.
	//
	// There are three classes of return values:
	//	nil -- the attempt to connect succeeded.
	//	ErrConnectStarted -- the connect attempt started but hasn't
	//		completed yet. In this case, the actual result will
	//		become available via GetSockOpt(ErrorOption) when
	//		the endpoint becomes writable. (This mimics the
	//		connect(2) syscall behavior.)
	//	Anything else -- the attempt to connect failed.
	Connect(address FullAddress) error

	// ConnectEndpoint connects this endpoint directly to another.
	//
	// This should be called on the client endpoint, and the (bound)
	// endpoint passed in as a parameter.
	//
	// The error codes are the same as Connect.
	ConnectEndpoint(server Endpoint) error

	// Shutdown closes the read and/or write end of the endpoint connection
	// to its peer.
	Shutdown(flags ShutdownFlags) error

	// Listen puts the endpoint in "listen" mode, which allows it to accept
	// new connections.
	Listen(backlog int) error

	// Accept returns a new endpoint if a peer has established a connection
	// to an endpoint previously set to listen mode. This method does not
	// block if no new connections are available.
	//
	// The returned Queue is the wait queue for the newly created endpoint.
	Accept() (Endpoint, *waiter.Queue, error)

	// Bind binds the endpoint to a specific local address and port.
	// Specifying a NIC is optional.
	//
	// An optional commit function will be executed atomically with respect
	// to binding the endpoint. If this returns an error, the bind will not
	// occur and the error will be propagated back to the caller.
	Bind(address FullAddress, commit func() error) error

	// GetLocalAddress returns the address to which the endpoint is bound.
	GetLocalAddress() (FullAddress, error)

	// GetRemoteAddress returns the address to which the endpoint is
	// connected.
	GetRemoteAddress() (FullAddress, error)

	// Readiness returns the current readiness of the endpoint. For example,
	// if waiter.EventIn is set, the endpoint is immediately readable.
	Readiness(mask waiter.EventMask) waiter.EventMask

	// SetSockOpt sets a socket option.
	SetSockOpt(interface{}) error

	// GetSockOpt gets a socket option.
	GetSockOpt(interface{}) error
}
