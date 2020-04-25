package tcpip

// Stats holds statistics about the networking stack.
type Stats struct {
	// UnkownProtocolRcvdPackets is the number of packets received by the
	// stack that were for an unknown or unsupported protocol.
	UnknownProtocolRcvdPackets uint64

	// UnknownNetworkEndpointRcvdPackets is the number of packets received
	// by the stack that were for a supported network protocol, but whose
	// destination address didn't having a matching endpoint.
	UnknownNetworkEndpointRcvdPackets uint64

	// MalformedRcvPackets is the number of packets received by the stack
	// that were deemed malformed.
	MalformedRcvdPackets uint64
}
