package db

import (
	"log"
)

func GetNumberOfLikes(postID uint) (int64, error) {

	var count int64

	err := GetDB().Table("likes").Where("post_id = ?", postID).Count(&count).Error

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return count, nil
}