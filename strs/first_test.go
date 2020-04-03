package strs

import "testing"

func TestUpperFirst(t *testing.T) {
	got := UpperFirst("", true)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := UpperFirst("a", true)
	want2 := "A"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := UpperFirst("aB", false)
	want3 := "AB"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := UpperFirst("aB", true)
	want4 := "Ab"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := UpperFirst("aBC", true)
	want5 := "Abc"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}
}

func TestLowerFirst(t *testing.T) {
	got := LowerFirst("", true)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := LowerFirst("A", true)
	want2 := "a"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}

	got3 := LowerFirst("Ab", false)
	want3 := "ab"
	if got3 != want3 {
		t.Errorf("got %q; want %q", got3, want3)
	}

	got4 := LowerFirst("Ab", true)
	want4 := "aB"
	if got4 != want4 {
		t.Errorf("got %q; want %q", got4, want4)
	}

	got5 := LowerFirst("Abc", true)
	want5 := "aBC"
	if got5 != want5 {
		t.Errorf("got %q; want %q", got5, want5)
	}
}
