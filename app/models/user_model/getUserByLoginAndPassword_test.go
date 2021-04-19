package user_model

import "testing"

func Test_getUserByLoginAndPassword(t *testing.T) {
	user, err := getUserByLoginAndPassword("biba", "boba")
	if err != nil {
		t.Log("failed", err)
		t.FailNow()
	}

	t.Log(user)

	user, err = getUserByLoginAndPassword("vaspetгaoggo", "lolkekaa")
	if err == nil {
		t.Log("error expected")
		t.Log(user)
		t.FailNow()
	}

	t.Log("passed")
}
