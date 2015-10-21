package comms

import (
	"testing"
)

func TestSerial(t *testing.T) {
	p := InitComms()

	_, err := SendReceive(p, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x7F})
	if err != nil {
		t.Error("SendReceive failed")
	}
}
