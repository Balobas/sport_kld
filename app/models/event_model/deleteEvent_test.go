package event_model

import "testing"

func TestDeleteEvent(t *testing.T) {
	if err := DeleteEvent("18b1f399-7a72-11eb-a516-0c9d9244", ""); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}
