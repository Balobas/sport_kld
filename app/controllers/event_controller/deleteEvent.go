package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func DeleteEvent(eventUid string) error {
	return event_model.DeleteEvent(models.UID(eventUid))
}
