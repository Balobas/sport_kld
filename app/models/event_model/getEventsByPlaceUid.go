package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventsByPlaceUid(placeUid models.UID) ([]Event, []error) {
	rows, err := database.MysqlDB.Queryx("select * from events where placeUid = ?", placeUid)
	if err != nil {
		return []Event{}, []error{errors.WithStack(err)}
	}

	var events []Event
	var resultErrors []error

	for rows.Next() {
		var event Event
		if err := rows.StructScan(&event); err != nil {
			resultErrors = append(resultErrors, err)
			continue
		}

		events = append(events, event)
	}

	return events, resultErrors
}