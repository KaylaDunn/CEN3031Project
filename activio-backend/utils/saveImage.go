package utils

import (
	"errors"
	"io"
	"mime/multipart"
)


func SaveImage(fileHeader *multipart.FileHeader, quality int) (string, error) {
	// This function is used to save an image to the server

	// Check if the file is an image
	fileType := fileHeader.Header["Content-Type"][0]
		if fileType != "image/jpeg" && fileType != "image/png" {
			return "", errors.New("file is not an image")
		}

	// Open the file
	file, err := fileHeader.Open()

	if err != nil {
		return "", err
	}

	// Read the file
	buffer, err := io.ReadAll(file)

	if err != nil {
		return "", err
	}

	// Close the file
	defer file.Close()

	// Compress the image
	hashedFilename, err := CompressImage(buffer, quality, fileHeader.Filename)

	if err != nil {
		return "", err
	}

	return hashedFilename, nil
}