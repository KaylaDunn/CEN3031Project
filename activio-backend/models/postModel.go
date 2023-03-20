package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID uint `json:"postedBy" gorm:"index"` // User ID
	Description string `json:"postDescription"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	LocationName string `json:"locationName"`
}

