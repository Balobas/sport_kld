package event_model

import "testing"

func TestGetEventInfoPost(t *testing.T) {
	post, err := GetEventInfoPost("16154873-777e-11eb-8737-0c9d9244", )
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(post)

	t.Log("passed")
}