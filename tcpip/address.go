package tcpip

import (
	"fmt"
)

// Address is a byte slice cast as a string that represents the address of a
// network node. Or, in the case of unix endpoints, it may represent a path
type Address string

// String implements the fmt.Stringer interface.
func (a Address) String() string {
	switch len(a) {
	case 4:
		return fmt.Sprintf("%d.%d.%d.%d", int(a[0]), int(a[1]), int(a[2]), int(a[3]))
	default:
		return fmt.Sprintf("%x", []byte(a))
	}
}

// FullAddress represents a full transport node address, as required by the
// Connect() and Bind() methods
type FullAddress struct {
	// NIC is the ID of the NIC this address refers to.
	//
	// This may not be used by all endpoint types.
	NIC NICID

	// Addr is the network address.
	Addr Address

	// Port is the transport port.
	//
	// This may not be used by all endpoint types.
	Port uint16
}
