package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	// ProfilePicture string `json:"profilePicture"` TODO:IMPLEMENT LATER
}
