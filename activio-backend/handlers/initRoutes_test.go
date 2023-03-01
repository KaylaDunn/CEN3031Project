package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestAliveEndpoint(t *testing.T) {

	data := map[string]interface{}{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/alive"

	request, _ := http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	body, _ := io.ReadAll(response.Body)

	got := string(body) == "{\"message\":\"I'm alive!\"}"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
