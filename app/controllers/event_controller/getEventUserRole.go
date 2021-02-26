package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventUserRole(eventUid, userUid string) (event_model.EventUserRole, error) {
	return event_model.GetEventUserRoleByUid(models.UID(eventUid), models.UID(userUid))
}
