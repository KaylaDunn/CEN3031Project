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

type CommentApiResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"createdAt"`
	Comment string `json:"comment"`
	UserID uint `json:"commentedBy"`
	PostID uint `json:"postID"`
}