package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestCreatePost(t *testing.T) {

	//login first

	data := map[string]interface{}{
		"password": "bobspassword",
		"email":    "bob@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/login"

	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	token := response.Header.Get("Set-Cookie")
	token = token[14:strings.IndexByte(token, ';')]

	data = map[string]interface{}{
		"postDescription": "Created using api",
		"longitude":       -5.234,
		"latitude":        0.234,
		"locationName":    "paris",
	}

	jsonData, err = json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl = "http://0.0.0.0:3000/api/auth/createpost"

	request, _ = http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	cookie := &http.Cookie{
		Name:   "Authorization",
		Value:  token,
		MaxAge: 86400,
	}
	request.AddCookie(cookie)

	client = &http.Client{}
	response, error = client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	got := response.Status == "200 OK"
	want := true

	body, _ := io.ReadAll(response.Body)

	println("status: " + response.Status)
	println("body: " + string(body))

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestGetPosts(t *testing.T) {

	data := map[string]interface{}{
		"password": "bobspassword",
		"email":    "bob@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/posts"

	request, _ := http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	got := response.Status == "200 OK"
	want := true

	println(response.Status)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
