package event_model

import "testing"

func TestCreateEvent(t *testing.T) {
	event := Event{
		UID:           "123",
		Name:          "test event",
		Description:   "test event description",
		Dates:         "25.04.2021",
		Time:          "13:00",
		VisitorsNum:   0,
		VisitorsLimit: 25,
		PlaceUID:      "8912829312312312",
		CreatorUID:    "123123123123",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}

	uid, err := CreateEvent(event)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(uid)

	event = Event{
		UID:           "123",
		Name:          "",
		Description:   "test event description",
		Dates:         "25.04.2021",
		Time:          "13:00",
		VisitorsNum:   0,
		VisitorsLimit: 25,
		PlaceUID:      "8912829312312312",
		CreatorUID:    "123123123123",
		IsPrivate:     false,
		EventPassword: "",
		IsOver:        false,
	}
	uid, err = CreateEvent(event)
	if err == nil {
		t.Log("must been failed")
		t.FailNow()
	}

	t.Log("passed")
}
