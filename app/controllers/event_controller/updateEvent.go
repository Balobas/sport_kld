package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func UpdateEvent(event event_model.Event, executorUid string) error {
	return event_model.UpdateEvent(event, models.UID(executorUid))
}
