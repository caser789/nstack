package tcpip

// NICID is a number that uniquely identifies a NIC
type NICID int32

// ShutdownFlags represents flags that can be passed to the Shutdown() method
// of the Endpoint interface
type ShutdownFlags int

// A ControlMessages represents a collection of socket control messages.
type ControlMessages interface {
	// Release releases any resources owned by the control message.
	Release()

	// CloneCreds returns a copy of any credentials (if any) contained in the
	// ControlMessages.
	CloneCreds() ControlMessages
}

// LinkEndpointID represents a data link layer endpoint.
type LinkEndpointID uint64

// TransportProtocolNumber is the number of a transport protocol.
type TransportProtocolNumber uint32

// NetworkProtocolNumber is the number of a network protocol.
type NetworkProtocolNumber uint32
