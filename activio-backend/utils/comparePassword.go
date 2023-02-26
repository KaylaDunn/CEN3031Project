package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(test string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(test))
	return err == nil
}