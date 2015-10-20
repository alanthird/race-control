package comms

import (
	"testing"
)

func TestSerial(t *testing.T) {
	p := initComms()

	read(p)
}
