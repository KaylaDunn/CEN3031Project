package utils

import (
	"os"
	"testing"
)

func TestCreateImageDir(t *testing.T) {

	CreateImageDir()
	imageDirExists := false

	_, err := os.Stat("./images")
	if err == nil {
		imageDirExists = true
	}

	got := imageDirExists
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
