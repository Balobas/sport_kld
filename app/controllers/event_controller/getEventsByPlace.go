package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventsByPlace(placeUid string) ([]event_model.Event, []error) {
	return event_model.GetEventsByPlaceUid(models.UID(placeUid))
}
