package event_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/event_model"
	"sport_kld/app/models/user_model"
)

func JoinEvent(user user_model.User, eventUid, password string) error {
	return event_model.JoinEvent(user, models.UID(eventUid), password)
}
