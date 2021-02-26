package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func DeleteUserFromEvent(userUid, eventUid, executorUid string) error {
	return event_model.DeleteUserFromEvent(models.UID(userUid), models.UID(eventUid), models.UID(executorUid))
}
