package event_model

import (
	"errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func DeleteEvent(eventUid, executorUid models.UID) error {
	event, err := GetEventByUid(eventUid)
	if err != nil {
		return err
	}
	if event.CreatorUID != executorUid {
		return errors.New("permission denied")
	}
	if _, err := database.MysqlDB.Exec("delete from events where uid = ?", eventUid); err != nil {
		return err
	}
	if _, err := database.MysqlDB.Queryx("delete from event_user_roles where eventUid = ?", eventUid); err != nil {
		return err
	}
	return nil
}
