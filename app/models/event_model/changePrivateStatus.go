package event_model

import (
	"errors"
	"sport_kld/app/models"
)

func ChangePrivateStatus(eventUid models.UID) error {
	event, err := GetEventByUid(eventUid)
	if err != nil {
		return err
	}
	if event.IsOver {
		return errors.New("event is over")
	}

	event.IsPrivate = !event.IsPrivate
	return putEvent(event)
}
