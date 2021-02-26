package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func AddEventToPlace(eventUid, placeUid models.UID) error {
	_, err := database.MysqlDB.Exec("INSERT into place_events(place_uid, event_uid) VALUES (?, ?)", placeUid, eventUid)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
