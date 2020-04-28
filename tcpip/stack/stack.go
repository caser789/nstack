package stack

import (
	"github.com/caser789/nstack/tcpip"
)

type Stack struct {
	transportProtocols map[tcpip.TransportProtocolNumber]int
}
