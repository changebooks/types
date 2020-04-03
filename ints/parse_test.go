package ints

import "testing"

func TestParse(t *testing.T) {
	got, _ := Parse(nil)
	if got != 0 {
		t.Errorf("got %d; want 0", got)
	}

	got2, _ := Parse(2)
	if got2 != 2 {
		t.Errorf("got %d; want 2", got2)
	}

	got3, _ := Parse(-3)
	if got3 != -3 {
		t.Errorf("got %d; want -3", got3)
	}

	var num4 int8 = 4
	got4, _ := Parse(num4)
	if got4 != 4 {
		t.Errorf("got %d; want 4", got4)
	}

	var num5 int16 = 5
	got5, _ := Parse(num5)
	if got5 != 5 {
		t.Errorf("got %d; want 5", got5)
	}

	var num6 int32 = 6
	got6, _ := Parse(num6)
	if got6 != 6 {
		t.Errorf("got %d; want 6", got6)
	}

	var num7 int32 = 7
	got7, _ := Parse(num7)
	if got7 != 7 {
		t.Errorf("got %d; want 7", got7)
	}

	var num8 int64 = 8
	got8, _ := Parse(num8)
	if got8 != 8 {
		t.Errorf("got %d; want 8", got8)
	}

	num9 := "9"
	got9, _ := Parse(num9)
	if got9 != 9 {
		t.Errorf("got %d; want 9", got9)
	}

	num10 := []byte("10")
	got10, _ := Parse(num10)
	if got10 != 10 {
		t.Errorf("got %d; want 10", got10)
	}

	num11 := true
	got11, _ := Parse(num11)
	if got11 != 1 {
		t.Errorf("got %d; want 1", got11)
	}

	num12 := false
	got12, _ := Parse(num12)
	if got12 != 0 {
		t.Errorf("got %d; want 0", got12)
	}

	var num13 float32 = 12.34
	got13, _ := Parse(num13)
	if got13 != 12 {
		t.Errorf("got %d; want 12", got13)
	}

	var num14 float64 = 56.78
	got14, _ := Parse(num14)
	if got14 != 56 {
		t.Errorf("got %d; want 56", got14)
	}

	var num15 uint8 = 15
	got15, _ := Parse(num15)
	if got15 != 15 {
		t.Errorf("got %d; want 15", got15)
	}

	var num16 uint16 = 16
	got16, _ := Parse(num16)
	if got16 != 16 {
		t.Errorf("got %d; want 16", got16)
	}

	var num17 uint32 = 17
	got17, _ := Parse(num17)
	if got17 != 17 {
		t.Errorf("got %d; want 17", got17)
	}

	var num18 uint = 18
	got18, _ := Parse(num18)
	if got18 != 18 {
		t.Errorf("got %d; want 18", got18)
	}

	var num19 uint64 = 19
	got19, _ := Parse(num19)
	if got19 != 19 {
		t.Errorf("got %d; want 19", got19)
	}
}
