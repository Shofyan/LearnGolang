package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseURL = "http://localhost:8181"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", ID)
	payload := bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/user", payload)
	if err != nil {
		return data, err
	}

	//	request.Header.Set("Content-Type","application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil

}

func main() {

	var users, err = fetchUsers()
	if err != nil {
		fmt.Println("error :", err.Error())
	}

	for _, item := range users {
		fmt.Printf("id: %s, nama :%s, grade:%d \n", item.ID, item.Name, item.Grade)
	}

	var us, err1 = fetchUser("a01")
	if err1 != nil {
		fmt.Println("error :", err1.Error())
	}
	fmt.Printf("id: %s, nama :%s, grade:%d \n", us.ID, us.Name, us.Grade)
}
