package event_model

import "testing"

func TestDeleteUserFromEvent(t *testing.T) {
	if err := DeleteEventUser("12", "123", "12"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}
