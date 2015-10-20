package car

type Car struct {
	Power      uint8
	Brake      bool
	LaneChange bool
}

func laneChangeByte(s bool) byte {
	if s {
		return 0x80
	}
	return 0x00
}

func brakeByte(s bool) byte {
	if s {
		return 0x40
	}
	return 0x00
}

func ToByte(c Car) byte {
	return (laneChangeByte(c.LaneChange) | brakeByte(c.Brake) | (0x3F & c.Power)) ^ 0xFF
}

func FromByte(b byte) (c Car) {
	c.Brake = (b & 0x80) == 0
	c.LaneChange = (b & 0x40) == 0
	c.Power = (b ^ 0xFF) & 0x3F

	return
}
