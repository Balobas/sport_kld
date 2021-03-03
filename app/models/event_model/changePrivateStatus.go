package event_model

import (
	"errors"
	"sport_kld/app/models"
)

func ChangeEventPrivateStatus(eventUid, executorUid models.UID) error {
	event, err := GetEventByUid(eventUid)
	if err != nil {
		return err
	}

	if event.CreatorUID != executorUid {
		return errors.New("permission denied")
	}

	if event.IsOver {
		return errors.New("event is over")
	}

	event.IsPrivate = !event.IsPrivate
	return putEvent(event)
}
