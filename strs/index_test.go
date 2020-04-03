package strs

import "testing"

func TestCaseIndex(t *testing.T) {
	got := CaseIndex("ChangeBooks", "O")
	want := 7
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestLastCaseIndex(t *testing.T) {
	got := LastCaseIndex("ChangeBooks", "O")
	want := 8
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
