package strs

import "testing"

func TestParse(t *testing.T) {
	got := Parse(nil)
	if got != "" {
		t.Errorf("got %s; want \"\"", got)
	}

	got2 := Parse("")
	if got2 != "" {
		t.Errorf("got %s; want \"\"", got2)
	}

	got3 := Parse([]byte(""))
	if got3 != "" {
		t.Errorf("got %s; want \"\"", got3)
	}

	got4 := Parse([]byte("a"))
	if got4 != "a" {
		t.Errorf("got %s; want a", got4)
	}

	got5 := Parse("a")
	if got5 != "a" {
		t.Errorf("got %s; want a", got5)
	}

	got6 := Parse(0)
	if got6 != "0" {
		t.Errorf("got %s; want 0", got6)
	}

	got7 := Parse(12)
	if got7 != "12" {
		t.Errorf("got %s; want 12", got7)
	}

	var num8 uint = 34
	got8 := Parse(num8)
	if got8 != "34" {
		t.Errorf("got %s; want 34", got8)
	}

	var num9 float32 = 12.34
	got9 := Parse(num9)
	if got9 != "12.34" {
		t.Errorf("got %s; want 12.34", got9)
	}

	var num10 float64 = 0.78
	got10 := Parse(num10)
	if got10 != "0.78" {
		t.Errorf("got %s; want 0.78", got10)
	}

	got11 := Parse(false)
	if got11 != "false" {
		t.Errorf("got %s; want false", got11)
	}

	got12 := Parse(true)
	if got12 != "true" {
		t.Errorf("got %s; want true", got12)
	}
}
