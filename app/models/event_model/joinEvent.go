package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func JoinEvent(userUid , eventUid models.UID, password string) error {
	if len(userUid) == 0 {
		return errors.New("user uid is empty")
	}
	if len(eventUid) == 0 {
		return errors.New("event uid is empty")
	}

	var t struct {
		EventUid string `db:"event_uid"`
		UserUid string `db:"user_uid"`
	}

	if err := database.MysqlDB.Get(&t, "SELECT * FROM event_users WHERE event_uid=? AND user_uid=?", eventUid, userUid); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return errors.New("select error")
		}
	} else {
		return errors.New("user has already join to event")
	}

	event, err := GetEventByUid(eventUid)
	if err != nil {
		return errors.WithStack(err)
	}

	if event.IsOver {
		return errors.New("Event is over")
	}

	if event.IsPrivate {
		if password != event.EventPassword {
			return errors.New("wrong password")
		}
	}

	if event.VisitorsNum + 1 > event.VisitorsLimit {
		return errors.New("users limit is over")
	}

	event.VisitorsNum += 1

	if err := putEvent(event); err != nil {
		return errors.WithStack(err)
	}

	if _, err := database.MysqlDB.Exec("INSERT INTO event_users(event_uid, user_uid) VALUES (?, ?)", eventUid, userUid); err != nil {
		return errors.New("cant join to event")
	}

	role := EventUserRole{
		UserUID:         userUid,
		EventUID:        eventUid,
		Role:            "Посетитель",
		RoleDescription: "Стандартная роль",
	}

	return putEventUserRole(role)
}
