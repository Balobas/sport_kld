package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventInfoPost(postUid string) (event_model.EventInfoPost, error) {
	return event_model.GetEventInfoPost(models.UID(postUid))
}
