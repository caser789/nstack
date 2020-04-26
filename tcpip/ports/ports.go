// Package ports provides PortManager that manages allocating, reserving and releasing ports
package ports

import (
	"math"
	"math/rand"
	"sync"

	"github.com/caser789/nstack/tcpip"
)

const (
	// firstEphemeral is the first ephemeral port.
	firstEphemeral uint16 = 16000
)

type portDescriptor struct {
	network   tcpip.NetworkProtocolNumber
	transport tcpip.TransportProtocolNumber
	port      uint16
}

// PortManager manages allocating, reserving and releasing ports
type PortManager struct {
	mu             sync.RWMutex
	allocatedPorts map[portDescriptor]struct{}
}

// NewPortManager creates new PortManager
func NewPortManager() *PortManager {
	return &PortManager{allocatedPorts: make(map[portDescriptor]struct{})}
}

// PickEphemeralPort randomly chooses a starting point and iterates over all
// possible ephemeral ports, allowing the caller to decide whether a given port
// is suitable for its needs, and stopping when a port is found or an error
// occurs
func (s *PortManager) PickEphemeralPort(testPort func(p uint16) (bool, error)) (port uint16, err error) {
	count := uint16(math.MaxUint16 - firstEphemeral + 1)
	offset := uint16(rand.Int31n(int32(count)))

	for i := uint16(0); i < count; i++ {
		port = firstEphemeral + (offset+i)%count
		ok, err := testPort(port)
		if err != nil {
			return 0, err
		}

		if ok {
			return port, nil
		}
	}

	return 0, tcpip.ErrNoPortAvailable
}
