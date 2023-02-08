package utils

import (
	"crypto/sha256"
	"fmt"
)

func StringHash(inputString string) string {
	// Hash an input string using SHA256
	h := sha256.New()
	h.Write([]byte(inputString))

	// Convert the hashed bytes to a string
	hashedString := fmt.Sprintf("%x", h.Sum(nil))

	return hashedString
}
