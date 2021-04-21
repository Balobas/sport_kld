package user_model

import (
	"sport_kld/app/models"
	"testing"
)

func TestGetUserByUid(t *testing.T) {

	uid := models.UID("8bf8914b-818b-11eb-9e00-0c9d92446328")
	user, err := GetUserByUid(uid)
	if err != nil {
		t.Log("failed", err)
		t.FailNow()
	}
	t.Log(user)
	t.Log("passed")
}
