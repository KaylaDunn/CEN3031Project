package utils

import "testing"

func TestComparePassword(t *testing.T) {

	got := ComparePassword("password", HashPassword("password"))
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
func TestComparePasswordFalse(t *testing.T) {

	got := ComparePassword("wrong", HashPassword("password"))
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
