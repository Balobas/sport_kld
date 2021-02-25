package event_model

import (
	"sport_kld/app/models/user_model"
	"testing"
)

func TestJoinEvent(t *testing.T) {
	user := user_model.User{
		UID:      "1299",
		Name:     "user",
		Login:    "u",
		Password: "p",
	}

	if err := JoinEvent(user, "123", ""); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if err := JoinEvent(user, "19191919", ""); err == nil {
		t.Log("must return error")
		t.FailNow()
	}

	user.UID = "laasdaf"

	if err := JoinEvent(user, "19191919", "lo"); err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("passed")
}
