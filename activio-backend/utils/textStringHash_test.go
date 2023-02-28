package utils

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestStringHash(t *testing.T) {

	// Hash an input string using SHA256
	h := sha256.New()
	h.Write([]byte("password"))

	// Convert the hashed bytes to a string
	hashedString := fmt.Sprintf("%x", h.Sum(nil))

	got := hashedString == StringHash("password")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	// Hash an input string using SHA256

	h.Write([]byte("incorrectpassword"))

	// Convert the hashed bytes to a string
	hashedString = fmt.Sprintf("%x", h.Sum(nil))

	got = hashedString == StringHash("password")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
