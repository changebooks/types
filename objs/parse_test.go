package objs

import "testing"

func TestParseBool(t *testing.T) {
	got, _ := ParseBool(nil)
	if got == true {
		t.Errorf("got %v; want false", got)
	}

	got2, _ := ParseBool([]byte(""))
	if got2 == true {
		t.Errorf("got %v; want false", got2)
	}

	data3 := []byte("")
	data3 = append(data3, 0x00)
	got3, _ := ParseBool(data3)
	if got3 == true {
		t.Errorf("got %v; want false", got3)
	}

	data4 := []byte("")
	data4 = append(data4, 0x01)
	got4, _ := ParseBool(data4)
	if got4 == false {
		t.Errorf("got %v; want true", got4)
	}
}

func TestParseFloat(t *testing.T) {
	got, _ := ParseFloat(3.14)
	want := 3.14
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}

	got2, _ := ParseFloat("3.14")
	want2 := 3.14
	if got2 != want2 {
		t.Errorf("got %v; want %v", got2, want2)
	}

	got3, _ := ParseFloat([]byte("3.14"))
	want3 := 3.14
	if got3 != want3 {
		t.Errorf("got %v; want %v", got3, want3)
	}

	got4, _ := ParseFloat(3)
	var want4 float64 = 3
	if got4 != want4 {
		t.Errorf("got %v; want %v", got4, want4)
	}
}
