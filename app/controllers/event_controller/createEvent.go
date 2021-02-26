package event_controller

import "sport_kld/app/models/event_model"

func CreateEvent(event event_model.Event) (string, error) {
	uid, err := event_model.CreateEvent(event)
	return string(uid), err
}
