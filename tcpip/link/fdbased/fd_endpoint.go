package fdbased

import (
	"github.com/caser789/nstack/tcpip"
)

type endpoint struct {
	// fd is the file descriptor used to send and receive packets.
	fd int

	// mtu (maximum transmission unit) is the maximum size of a packet.
	mtu int

	// closed is a function to be called when the FD's peer (if any) closes
	// its end of the communication pipe.
	closed func(error)
}
