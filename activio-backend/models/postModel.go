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

type PostApiResponse struct {
	ID uint `json:"id"`
	CreatedAt string `json:"createdAt"`
	UserID uint		`json:"postedBy"`
	Description string `json:"postDescription"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	LocationName string `json:"locationName"`
}

