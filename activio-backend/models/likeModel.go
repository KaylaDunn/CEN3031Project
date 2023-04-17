package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	PostID uint `json:"postID"`
	UserID uint `json:"likedBy"`
}