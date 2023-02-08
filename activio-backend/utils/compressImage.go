package utils

import (
	"log"

	"github.com/h2non/bimg"
)

// file path to the image directory
const imageDir = "images/"

func CompressImage(buffer []byte, quality int, fileName string) (string, error) {
	// TODO: Implement this endpoint
	hashName := StringHash(fileName)
	log.Println(hashName)

	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		log.Println(err)
		return fileName, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})

	write := bimg.Write(imageDir+hashName+".webp", processed)
	if write != nil {
		log.Println("Error saving file: " + fileName)
		log.Println(write)
		return fileName, err
	}

	return hashName, nil

}
