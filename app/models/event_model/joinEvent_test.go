package event_model

import (
	"sport_kld/app/models/user_model"
	"testing"
)

func TestJoinEvent(t *testing.T) {
	user := user_model.User{
		UID:      "12998",
		Name:     "user",
		Login:    "u",
		Password: "p",
	}

	if err := JoinEvent(user.UID, "123", ""); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := JoinEvent(user.UID, "1491b99d-7778-11eb-8979-0c9d9244", ""); err == nil {
		t.Log("must return error")
		t.FailNow()
	}

	user.UID = "laasdaf"

	//if err := JoinEvent(user, "1491b99d-7778-11eb-8979-0c9d92446328", "lo"); err != nil {
	//	t.Log(err)
	//	t.FailNow()
	//}
	t.Log("passed")
}
