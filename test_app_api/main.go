package main

import (
	"crypto/md5"
	"fmt"
	"sport_kld/test_app_api/users"
)

func main() {
	//places.TestGetPlaceByUID()
	//places.TestGetPlacesByUIDs()
	//
	//events.TestCreateEvent()
	//users.TestCreateUser()
	users.TestLogin()

	h := md5.New()
	h.Write([]byte("boba"))
	h.Write([]byte("hashsalt"))
	pass := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println(pass)
}
