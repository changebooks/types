package strs

import "testing"

func TestMultiByteLen(t *testing.T) {
	got := MultiByteLen("汉字a")
	want := 3
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func TestMultiByteCut(t *testing.T) {
	got := MultiByteCut("", 0)
	want := ""
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}

	got2 := MultiByteCut("汉字", 1)
	want2 := "字"
	if got2 != want2 {
		t.Errorf("got %q; want %q", got2, want2)
	}
}
