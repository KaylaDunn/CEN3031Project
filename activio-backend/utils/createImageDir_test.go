package utils

import (
	"os"
	"testing"
)

func TestCreateImageDir(t *testing.T) {

	CreateImageDir()
	imageDirExists := true

	if _, err := os.Stat("../images"); os.IsNotExist(err) {
		imageDirExists = false
	}

	got := imageDirExists
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
