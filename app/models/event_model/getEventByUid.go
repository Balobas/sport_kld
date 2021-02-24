package event_model

import (
	"errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventByUid(uid models.UID) (Event, error) {
	var event Event
	if err := database.MysqlDB.Get(&event, "select * from events where uid=?", uid); err != nil {
		return Event{}, errors.New("cant get event")
	}
	return event, nil
}
