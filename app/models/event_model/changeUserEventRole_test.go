package event_model

import "testing"

func TestChangeUserRole(t *testing.T) {
	if err := ChangeUserRole(EventUserRole{
		UserUID:         "123123123123",
		EventUID:        "5f460829-76b7-11eb-8b56-0c9d9244",
		Role:            "Чорт",
		RoleDescription: "",
	}, "123123123123"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}