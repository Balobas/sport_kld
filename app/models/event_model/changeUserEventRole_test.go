package event_model

import "testing"

func TestChangeUserRole(t *testing.T) {
	if err := ChangeUserRole(EventUser{
		UserUID:         "12345",
		EventUID:        "123",
		Role:            "Чорт",
		RoleDescription: "",
	}, "123123123123"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}