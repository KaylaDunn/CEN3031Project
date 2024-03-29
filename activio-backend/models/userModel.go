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
	Birthday string `json:"birthday"`
	Phonenumber string `json:"phoneNumber"`
	Verified bool   `json:"verified"`
	ProfilePicture string `json:"profilePicture"`
}

type UserApiResponse struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	ProfilePicture string `json:"profilePicture"`
}

type UserUpdateInfo struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Username string `json:"username"`
	Phonenumber string `json:"phonenumber"`
	Birthday string `json:"birthday"`
}