package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventByUid(uid string) (event_model.Event, error) {
	return event_model.GetEventByUid(models.UID(uid))
}
