package user_model

import "testing"

func TestUpdateUser(t *testing.T) {
	user := User{
		UID:      "8bf8914b-818b-11eb-9e00-0c9d92446328",
		Name:     "vasya petrov",
		Login:    "VASSSui",
		Email:    "a@a.a",
		Age:      20,
		Password: "lolkek",
	}

	err := UpdateUser(user, "8bf8914b-818b-11eb-9e00-0c9d92446328")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}
