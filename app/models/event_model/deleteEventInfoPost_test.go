package event_model

import "testing"

func TestDeleteEventInfoPost(t *testing.T) {
	err := DeleteEventInfoPost("16154873-777e-11eb-8737-0c9d9244")
	if err != nil {
		t.Log("cant delete")
		t.FailNow()
	}

	t.Log("passed")
}
