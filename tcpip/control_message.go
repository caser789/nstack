package tcpip

// A ControlMessages represents a collection of socket control messages.
type ControlMessages interface {
	// Release releases any resources owned by the control message.
	Release()

	// CloneCreds returns a copy of any credentials (if any) contained in the
	// ControlMessages.
	CloneCreds() ControlMessages
}
