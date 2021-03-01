package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventUserRole(eventUid, userUid string) (event_model.EventUser, error) {
	return event_model.GetEventUserByUid(models.UID(eventUid), models.UID(userUid))
}
