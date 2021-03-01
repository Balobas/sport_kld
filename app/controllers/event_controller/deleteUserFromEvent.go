package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func DeleteEventUser(userUid, eventUid, executorUid string) error {
	return event_model.DeleteEventUser(models.UID(userUid), models.UID(eventUid), models.UID(executorUid))
}
