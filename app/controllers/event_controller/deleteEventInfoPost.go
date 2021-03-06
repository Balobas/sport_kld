package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func DeleteEventInfoPost(postUid, executorUid string) error {
	return event_model.DeleteEventInfoPost(models.UID(postUid), models.UID(executorUid))
}
