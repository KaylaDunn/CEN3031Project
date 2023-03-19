package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model	
	Comment string `json:"comment"`
	CommentedBy uint `json:"commentedBy"` // User ID
	CommentedOn uint `json:"commentedOn" gorm:"index"` // Post ID
}