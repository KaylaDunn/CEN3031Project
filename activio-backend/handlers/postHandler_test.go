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

var postID = ""

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

	body, _ := io.ReadAll(response.Body)
	println("status: " + response.Status)
	println("body: " + string(body))

	id := string(body)
	postID = id[14:16]
	//println("id: " + id[14:16])

	got := response.Status == "200 OK"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestCreatePostFail(t *testing.T) {

	data := map[string]interface{}{
		"postDescription": "Created using api",
		"longitude":       -5.234,
		"latitude":        0.234,
		"locationName":    "paris",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/auth/createpost"

	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	got := response.Status == "401 Unauthorized"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	body, _ := io.ReadAll(response.Body)

	got = string(body) == "{\"error\":\"Unauthorized\"}"
	want = true

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

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetPost(t *testing.T) {

	data := map[string]interface{}{
		"password": "bobspassword",
		"email":    "bob@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/post/1"

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

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	body, _ := io.ReadAll(response.Body)

	sillyString := "{\"post\":{\"id\":1,\"createdAt\":\"2023-03-20T21:20:52.848Z\",\"postedBy\":2,\"postDescription\":\"Created using api\",\"longitude\":-5.234,\"latitude\":0.234,\"locationName\":\"paris\",\"images\":[],\"comments\":[]}}"

	got = string(body) == sillyString
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetPostFail(t *testing.T) {

	data := map[string]interface{}{
		"password": "bobspassword",
		"email":    "bob@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/post/999"

	request, _ := http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	got := response.Status == "500 Internal Server Error"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	body, _ := io.ReadAll(response.Body)

	got = string(body) == "{\"message\":\"Post not found\"}"
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestDeletePost(t *testing.T) {

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

	data = map[string]interface{}{}

	jsonData, err = json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl = "http://0.0.0.0:3000/api/auth/deletepost/" + postID

	request, _ = http.NewRequest("DELETE", httpposturl, bytes.NewBuffer(jsonData))

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

	body, _ := io.ReadAll(response.Body)
	println("status: " + response.Status)
	println("body: " + string(body))

	got := response.Status == "200 OK"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = string(body) == "{\"message\":\"Post deleted successfully\"}"
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestDeletePostUnauthorized(t *testing.T) {

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

	data = map[string]interface{}{}

	jsonData, err = json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl = "http://0.0.0.0:3000/api/auth/deletepost/2"

	request, _ = http.NewRequest("DELETE", httpposturl, bytes.NewBuffer(jsonData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client = &http.Client{}
	response, error = client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	println("status: " + response.Status)
	println("body: " + string(body))

	got := response.Status == "401 Unauthorized"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = string(body) == "{\"error\":\"Unauthorized\"}"
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}
