package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID uint `json:"postedBy" gorm:"index"` // User ID
	Description string `json:"postDescription" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Latitude float64 `json:"latitude" binding:"required"`
	LocationName string `json:"locationName" binding:"required"`
}

