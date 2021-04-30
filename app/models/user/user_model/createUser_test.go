package user_model

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := User{
		UID:      "",
		Name:     "vasya pertov",
		Login:    "vas",
		Email:    "adas@asd.au",
		Age:      20,
		Password: "lolkek",
	}

	uid, err := CreateUser(user)
	if err != nil {
		t.Log("failed", err)
		t.FailNow()
	}

	t.Log("successful created, uid: ", uid.String(), "  ")

	_, err = CreateUser(user)
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	_, err = CreateUser(User{
		UID:      "",
		Name:     "",
		Login:    "",
		Email:    "",
		Age:      0,
		Password: "",
	})
	if err == nil {
		t.Log("error expected")
		t.FailNow()
	}

	t.Log("passed")
}
