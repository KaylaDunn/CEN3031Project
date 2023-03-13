package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	PostedBy string `json:"postedBy"`
	PostDescription string `json:"postDescription"`
	PostImages []string `json:"postImages"`
	NumberOfLikes uint `json:"numberOfLikes"`
	Comments []Comments `json:"comments"`
}

