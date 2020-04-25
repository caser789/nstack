package tcpip

// ErrorOption is used in GetSockOpt to specify that the last error reported by
// the endpoint should be cleared and returned
type ErrorOption struct{}

// SendBufferSizeOption is used by SetSockOpt/GetSockOpt to specify the send
// buffer size option
type SendBufferSizeOption int

// ReceiveBufferSizeOption is used by SetSockOpt/GetSockOpt to specify the
// receive buffer size option.
type ReceiveBufferSizeOption int

// ReceiveQueueSizeOption is used in GetSockOpt to specify that the number of
// unread bytes in the input buffer should be returned.
type ReceiveQueueSizeOption int

// NoDelayOption is used by SetSockOpt/GetSockOpt to specify if data should be
// sent out immediately by the transport protocol. For TCP, it determines if the
// Nagle algorithm is on or off.
type NoDelayOption int

// ReuseAddressOption is used by SetSockOpt/GetSockOpt to specify whether Bind()
// should allow reuse of local address.
type ReuseAddressOption int

// PasscredOption is used by SetSockOpt/GetSockOpt to specify whether
// SCM_CREDENTIALS socket control messages are enabled.
//
// Only supported on Unix sockets.
type PasscredOption int3
