package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UploadedBy uint `json:"uploadedBy"` // User ID
	PostID uint `json:"postID"` // Post ID
	OriginalName string `json:"originalName"`
	HashedFileName string `json:"fileName"`
	Order int `json:"order"`
	IsProfilePicture bool `json:"isProfilePicture"`
}

type ImageApiResponse struct {
	HashedFileName string `json:"url"`
}
