package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func ChangeEventPrivateStatus(eventUid, executorUid string) error {
	return event_model.ChangeEventPrivateStatus(models.UID(eventUid), models.UID(executorUid))
}
