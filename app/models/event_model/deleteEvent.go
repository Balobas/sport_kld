package event_model

import (
	"sport_kld/app/models"
	"sport_kld/database"
)

func DeleteEvent(eventUid models.UID) error {
	if _, err := GetEventByUid(eventUid); err != nil {
		return err
	}
	if _, err := database.MysqlDB.Exec("delete from events where uid = ?", eventUid); err != nil {
		return err
	}
	//if _, err := database.MysqlDB.Queryx("delete from event_users where event_uid = ?", eventUid); err != nil {
	//	return err
	//}
	if _, err := database.MysqlDB.Queryx("delete from event_user_roles where eventUid = ?", eventUid); err != nil {
		return err
	}
	return nil
}
