package car

import "testing"

func TestToByte(t *testing.T) {
	c := Car{Power: 42, LaneChange: false, Brake: false}
	b := ToByte(c)
	if b != 0xD5 {
		t.Error("Expected 213, got ", b)
	}

	c = Car{Power: 0, LaneChange: true, Brake: true}
	b = ToByte(c)
	if b != 0x3F {
		t.Error("Expected 63, got", b)
	}
}

func TestFromByte(t *testing.T) {
	b := uint8(0x00)
	c := FromByte(b)

	if !c.Brake {
		t.Error("Expected brake true, got", c.Brake)
	}
	if !c.LaneChange {
		t.Error("Expected lane change true, got", c.LaneChange)
	}

	if c.Power != 63 {
		t.Error("Expected power level of 63, got", c.Power)
	}

	b = uint8(0xD5)
	c = FromByte(b)

	if c.Brake {
		t.Error("Expected brake false, got", c.Brake)
	}
	if c.LaneChange {
		t.Error("Expected lane change false, got", c.LaneChange)
	}

	if c.Power != 42 {
		t.Error("Expected power level of 42, got", c.Power)
	}
}
