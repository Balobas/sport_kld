package event_model

import "testing"

func TestChangeEventPrivateStatus(t *testing.T) {
	if err := ChangeEventPrivateStatus("1491b99d-7778-11eb-8979-0c9d9244", ""); err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("passed")
}
