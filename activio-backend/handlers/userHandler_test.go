package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {

	/*content, err := ioutil.ReadFile("./testSignup.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}*/

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

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)

	//save this business for now for reference later
	//body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("response Body:", string(body))

	//fmt.Printf("json data: %s\n", jsonData)

	got := response.Status == "200 OK"
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
