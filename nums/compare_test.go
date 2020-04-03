package nums

import "testing"

func TestMax(t *testing.T) {
	got := Max(0)
	if got != 0 {
		t.Errorf("got %f; want 0", got)
	}

	got2 := Max(-1)
	if got2 != -1 {
		t.Errorf("got %f; want -1", got2)
	}

	got3 := Max(-1, -2)
	if got3 != -1 {
		t.Errorf("got %f; want -1", got3)
	}

	got4 := Max(-3, -1, -2)
	if got4 != -1 {
		t.Errorf("got %f; want -1", got4)
	}

	got5 := Max(3, 1, 2)
	if got5 != 3 {
		t.Errorf("got %f; want 3", got5)
	}

	got6 := Max(0, 1, -2)
	if got6 != 1 {
		t.Errorf("got %f; want 1", got6)
	}
}

func TestMin(t *testing.T) {
	got := Min(0)
	if got != 0 {
		t.Errorf("got %f; want 0", got)
	}

	got2 := Min(-1)
	if got2 != -1 {
		t.Errorf("got %f; want -1", got2)
	}

	got3 := Min(-1, -2)
	if got3 != -2 {
		t.Errorf("got %f; want -2", got3)
	}

	got4 := Min(-3, -1, -2)
	if got4 != -3 {
		t.Errorf("got %f; want -3", got4)
	}

	got5 := Min(3, 1, 2)
	if got5 != 1 {
		t.Errorf("got %f; want 1", got5)
	}

	got6 := Min(0, 1, -2)
	if got6 != -2 {
		t.Errorf("got %f; want -2", got6)
	}
}
