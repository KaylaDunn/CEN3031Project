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

func TestLogin(t *testing.T) {

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

	got := response.Status == "200 OK"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestLoginFail(t *testing.T) {

	data := map[string]interface{}{
		"password": "notbobspassword",
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

	got := response.Status == "400 Bad Request"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSignup(t *testing.T) {

	data := map[string]interface{}{
		"username":  "bigchungus",
		"password":  "bigchunguspassword",
		"firstName": "Bugs",
		"lastName":  "Bunny",
		"email":     "bigchungus@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/signup"

	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	got := response.Status == "200 OK"
	want := true

	body, _ := io.ReadAll(response.Body)

	if string(body) == "{\"Error\":\"User already exists\"}" {
		got = true
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestDeleteUserAndUserData(t *testing.T) {

	//login first

	data := map[string]interface{}{
		"password": "bigchunguspassword",
		"email":    "bigchungus@ufl.edu",
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

	httpposturl = "http://0.0.0.0:3000/api/auth/deleteaccount"

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

	got := response.Status == "200 OK"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetUserDetails(t *testing.T) {

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

	httpposturl = "http://0.0.0.0:3000/api/auth/me"

	request, _ = http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))

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

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestGetUserDetailsFail(t *testing.T) {

	data := map[string]interface{}{
		"password": "notbobspassword",
		"email":    "bob@ufl.edu",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	httpposturl := "http://0.0.0.0:3000/api/auth/me"

	request, _ := http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))

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

func TestRefreshJWT(t *testing.T) {

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

	httpposturl = "http://0.0.0.0:3000/api/auth/refreshtoken"

	request, _ = http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))

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

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

}

func TestRefreshJWTFail(t *testing.T) {

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

	httpposturl = "http://0.0.0.0:3000/api/auth/refreshtoken"

	request, _ = http.NewRequest("GET", httpposturl, bytes.NewBuffer(jsonData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client = &http.Client{}
	response, error = client.Do(request)
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
