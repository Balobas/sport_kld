package event_model

import "testing"

func TestDeleteEvent(t *testing.T) {
	if err := DeleteEvent("5f460829-76b7-11eb-8b56-0c9d9244"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}
