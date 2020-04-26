package tcpip

import (
	"errors"
)

// Errors that can be returned by the network stack.
var (
	ErrUnknownProtocol      = errors.New("unknown protocol")
	ErrUnknownNICID         = errors.New("unknown nic id")
	ErrDuplicateNICID       = errors.New("duplicate nic id")
	ErrDuplicateAddress     = errors.New("duplicate address")
	ErrNoRoute              = errors.New("no route")
	ErrBadLinkEndpoint      = errors.New("bad link layer endpoint")
	ErrAlreadyBound         = errors.New("endpoint already bound")
	ErrInvalidEndpointState = errors.New("endpoint is in invalid state")
	ErrAlreadyConnecting    = errors.New("endpoint is already connecting")
	ErrAlreadyConnected     = errors.New("endpoint is already connected")
	ErrNoPortAvailable      = errors.New("no ports are available")
	ErrPortInUse            = errors.New("port is in use")
	ErrBadLocalAddress      = errors.New("bad local address")
	ErrClosedForSend        = errors.New("endpoint is closed for send")
	ErrClosedForReceive     = errors.New("endpoint is closed for receive")
	ErrWouldBlock           = errors.New("operation would block")
	ErrConnectionRefused    = errors.New("connection was refused")
	ErrTimeout              = errors.New("operation timed out")
	ErrAborted              = errors.New("operation aborted")
	ErrConnectStarted       = errors.New("connection attempt started")
	ErrDestinationRequired  = errors.New("destination address is required")
	ErrNotSupported         = errors.New("operation not supported")
	ErrNotConnected         = errors.New("endpoint not connected")
	ErrConnectionReset      = errors.New("connection reset by peer")
	ErrConnectionAborted    = errors.New("connection aborted")
)
