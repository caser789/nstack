package ports

import (
	"testing"
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
