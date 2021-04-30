package user_model

import "testing"

func TestGetUserByLogin(t *testing.T) {
	user, err := GetUserByLogin("vaspetгa")
	if err != nil {
		t.Log("failed", err)
		t.FailNow()
	}

	t.Log("passed", user)
}
