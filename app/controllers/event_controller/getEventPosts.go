package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
)

func GetEventInfoPosts(eventUid string) ([]event_model.EventInfoPost, []error) {
	return event_model.GetEventInfoPosts(models.UID(eventUid))
}
