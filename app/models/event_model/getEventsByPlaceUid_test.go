package event_model

import "testing"

func TestGetEventsByPlaceUid(t *testing.T) {
	events, errors := GetEventsByPlaceUid("asjdkasdjkasdjkasd")
	t.Log(errors)
	t.Log(events)
}
