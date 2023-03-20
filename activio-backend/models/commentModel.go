package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model	
	Comment string `json:"comment"`
	UserID uint `json:"commentedBy"` // User ID
	PostID uint `json:"postID" gorm:"index"` // Post ID
}