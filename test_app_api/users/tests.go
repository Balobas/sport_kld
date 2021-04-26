package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sport_kld/app/models/user_model"
	"sport_kld/test_app_api/config"
)

func TestCreateUser() {
	user := user_model.User{
		UID:      "",
		Name:     "balobasr",
		Login:    "biba",
		Email:    "sadsd",
		Age:      20,
		Password: "boba",
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(config.BaseUrl + "/user", "application/json", r)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("DONE: CreateUser test")
}

func TestLogin() {
	user := user_model.User{
		Login:    "biba",
		Password: "boba",
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	r := bytes.NewReader(b)
	resp, err := http.Post(config.BaseUrl + "/login", "application/json", r)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("DONE: Login test")
}

func TestLogout() {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdWlkIjoiOTg2Mjk4YmEtMzE4Yi00MjdmLTgyNjEtNmMyMTIzM2YwZjYyIiwiZXhwIjoxNjE5MzQ5ODExLCJ1c2VyX3VpZCI6ImE1MTY4OGUwLWEwZmItMTFlYi05ODJjLWVhNGMyZGE3NTk3OCJ9.T-t8QhiPOk523VeCfuRsLkPu7DFcWxIl1uLW8ykG1Mo"
	client := &http.Client{}
	URL := config.BaseUrl + "/logout"
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("FAIL: ", err)
		return
	}
	fmt.Println(string(b))

	fmt.Println("DONE: Logout test")
}
