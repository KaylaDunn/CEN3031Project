package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	PostedBy uint `json:"postedBy"` // User ID
	Description string `json:"postDescription"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	LocationName string `json:"locationName"`
}

