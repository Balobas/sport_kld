package event_model

import "testing"

func TestChangeEventPrivateStatus(t *testing.T) {
	if err := ChangeEventPrivateStatus("19191919"); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}