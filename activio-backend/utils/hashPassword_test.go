package utils

import "testing"

func TestHashPassword(t *testing.T) {

	got := ComparePassword("password", HashPassword("password"))
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
func TestHashPasswordFalse(t *testing.T) {
	got := ComparePassword("incorrect", HashPassword("password"))
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
