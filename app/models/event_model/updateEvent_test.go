package event_model

import "testing"

func TestUpdateEvent(t *testing.T) {
	if err := UpdateEvent(Event{
		UID:           "1491b99d-7778-11eb-8979-0c9d9244",
		Name:          "test event update",
		Description:   "test event update description",
		Dates:         "12",
		Time:          "12",
		VisitorsNum:   0,
		VisitorsLimit: 10,
		PlaceUID:      "asjdkasdjkasdjkasd",
		CreatorUID:    "ttt",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}, "ttt"); err != nil {
			t.Log(err)
			t.FailNow()
	}

	t.Log("passed")
}
