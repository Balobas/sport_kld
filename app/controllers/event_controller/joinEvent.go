package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func JoinEvent(userUid, eventUid, password string) error {
	return event_model.JoinEvent(models.UID(userUid), models.UID(eventUid), password)
}
