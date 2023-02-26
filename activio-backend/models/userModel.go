package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email" binding:"required,email"`
	Verified bool   `json:"verified"`
	ProfilePicture string `json:"profilePicture"`
}
