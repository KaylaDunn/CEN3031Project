package models

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	CommentedBy string `json:"commentedBy"`
	Comment string `json:"comment"`
	PostId uint `json:"postId"`
	NumberOfLikes uint `json:"numberOfLikes"`
}