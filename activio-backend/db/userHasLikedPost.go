package db

import (
	"log"

	"gorm.io/gorm"

	"activio-backend/models"
)


func UserHasLikedPost(userID uint, postID uint) (bool, error) {
	var like models.Like

	err := GetDB().Table("likes").Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}