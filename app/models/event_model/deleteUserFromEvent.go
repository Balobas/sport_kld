package event_model

import (
	"errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func DeleteUserFromEvent(userUid, eventUid, executorUid models.UID) error {
	event, err := GetEventByUid(eventUid)
	if err != nil {
		return err
	}

	if event.CreatorUID != executorUid && userUid != executorUid {
		return errors.New("access denied")
	}

	if  _,err := database.MysqlDB.Exec("Delete from event_users where user_uid = ? and event_uid = ?", userUid, eventUid); err != nil {
		return err
	}
	if  _,err := database.MysqlDB.Exec("Delete from event_user_roles where userUid = ? and eventUid = ?", userUid, eventUid); err != nil {
		return err
	}

	return nil
}
