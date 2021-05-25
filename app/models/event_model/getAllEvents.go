package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/database"
)

func GetAllEvents() ([]Event, error){
	var events []Event
	if err := database.MysqlDB.Select(&events, "select * from events"); err != nil {
		return []Event{}, errors.Wrap(err, "cant get events")
	}
	return events, nil
}
