package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Verified bool   `json:"verified"`
	ProfilePicture string `json:"profilePicture"`
}
