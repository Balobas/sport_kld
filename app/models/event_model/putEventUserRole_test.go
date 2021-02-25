package event_model

import "testing"

func TestPutEventUserRole(t *testing.T) {
	if err := putEventUserRole(EventUserRole{
		UserUID:         "user",
		EventUID:        "event",
		Role:            "asd",
		RoleDescription: "asda",
	}); err != nil {
			t.Log(err)
			t.FailNow()
	}

	t.Log("passed")
}
