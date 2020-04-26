package ports

import (
	"testing"

	"github.com/caser789/nstack/tcpip"
)

func testPort(p uint16) (bool, error) {
	if p < firstEphemeral+100 {
		return false, nil
	}

	return true, nil
}

func TestPickEphemeralPort(t *testing.T) {
	p := NewPortManager()

	port, _ := p.PickEphemeralPort(testPort)

	if want, got := 16000, port; int(want) >= int(got) {
		t.Fatalf("TestPickEphemeralPort failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestReservePortErrPortInUse(t *testing.T) {
	p := NewPortManager()

	network := tcpip.NetworkProtocolNumber(1)
	transport := tcpip.TransportProtocolNumber(2)
	port := uint16(16700)

	desc := portDescriptor{network, transport, port}
	p.allocatedPorts[desc] = struct{}{}

	pt, e := p.ReservePort(network, transport, port)

	if want, got := 0, pt; int(want) != int(got) {
		t.Fatalf("TestReservePortErrPortInUse port failed:\n- want: %v\n- got: %v", want, got)
	}

	if want, got := tcpip.ErrPortInUse, e; want != got {
		t.Fatalf("TestReservePortErrPortInUse error failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestReservePortUseSpecifiedPort(t *testing.T) {
	p := NewPortManager()

	network := tcpip.NetworkProtocolNumber(1)
	transport := tcpip.TransportProtocolNumber(2)
	port := uint16(16700)

	pt, _ := p.ReservePort(network, transport, port)

	if want, got := port, pt; int(want) != int(got) {
		t.Fatalf("TestReservePortUseSpecifiedPort port failed:\n- want: %v\n- got: %v", want, got)
	}
}

func TestReservePort(t *testing.T) {
	p := NewPortManager()

	network := tcpip.NetworkProtocolNumber(1)
	transport := tcpip.TransportProtocolNumber(2)
	port := uint16(0)

	pt, _ := p.ReservePort(network, transport, port)

	if want, got := 16000, pt; int(want) >= int(got) {
		t.Fatalf("TestReservePort port failed:\n- want: %v\n- got: %v", want, got)
	}
}
