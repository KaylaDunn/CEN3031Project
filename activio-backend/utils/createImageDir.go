package utils

import "os"

// CreateImageDir creates the image directory if it does not exist
func CreateImageDir() {
	// Create the image directory if it does not exist
	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		os.Mkdir("./images", 0755)
	} else if err != nil {
		panic(err)
	}
}