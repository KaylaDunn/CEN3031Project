package utils

import "testing"

func TestHashPassword(t *testing.T) {

	got := ComparePassword("password", HashPassword("password"))
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
