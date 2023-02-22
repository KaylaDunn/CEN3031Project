package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required,min=2,max=20"`
	LastName string `json:"lastName" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Verified bool   `json:"verified"`
	ProfilePicture string `json:"profilePicture"`
}
