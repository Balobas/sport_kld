package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
)

func UpdateEvent(event Event, executorUid models.UID) error {
	if err := event.validate(); err != nil {
		return errors.Wrap(err, " cant update event")
	}

	oldEvent, err := GetEventByUid(event.UID)
	if err != nil {
		return err
	}

	if oldEvent.CreatorUID != executorUid {
		return errors.New("cant update event: access denied")
	}

	oldEvent.Name = event.Name
	oldEvent.Description = event.Description
	oldEvent.VisitorsLimit = event.VisitorsLimit
	oldEvent.Dates = event.Dates
	oldEvent.Time = event.Time
	oldEvent.EventPassword = event.EventPassword
	oldEvent.PlaceUID = event.PlaceUID

	return putEvent(oldEvent)
}
