package event_model

import (
	"sport_kld/app/models"
	"testing"
)

func TestGetEventByUid(t *testing.T) {
	uid := models.UID("123")

	event, err := GetEventByUid(uid)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(event)

	event, err = GetEventByUid("789")
	if err != nil {
		t.Log(err)
	}

	t.Log("passed")
}
