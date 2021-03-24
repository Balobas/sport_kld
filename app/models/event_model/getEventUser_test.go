package event_model

import "testing"

func TestGetEventUserRoleByUid(t *testing.T) {
	role, err := GetEventUserByUid("123", "1234")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(role)
	t.Log("passed")
}
