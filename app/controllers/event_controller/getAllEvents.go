package event_controller

import "sport_kld/app/models/event_model"

func GetAllEvents() ([]event_model.Event, error) {
	return event_model.GetAllEvents()
}
