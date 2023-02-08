package utils

import (
	"crypto/sha256"
	"fmt"
)

func StringHash(inputString string) string {

	h := sha256.New()
	h.Write([]byte(inputString))
	hashedString := fmt.Sprintf("%x", h.Sum(nil))
	return hashedString
}
