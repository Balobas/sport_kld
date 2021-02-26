package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func ChangeUserEventRole(role event_model.EventUserRole, executorUid string) error {
	return event_model.ChangeUserRole(role, models.UID(executorUid))
}
