package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"strings"
)

func JoinEvent(userUid, eventUid models.UID, password string) error {
	if len(userUid) == 0 {
		return errors.New("user uid is empty")
	}
	if len(eventUid) == 0 {
		return errors.New("event uid is empty")
	}

	if _, err := GetEventUserByUid(eventUid, userUid); err != nil {
		if !strings.Contains(err.Error(), "sql: no rows in result set") {
			return errors.Wrap(err, "select error")
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

	if event.VisitorsNum+1 > event.VisitorsLimit {
		return errors.New("users limit is over")
	}

	event.VisitorsNum += 1

	if err := putEvent(event); err != nil {
		return errors.WithStack(err)
	}

	//if _, err := database.MysqlDB.Exec("INSERT INTO event_users(event_uid, user_uid) VALUES (?, ?)", eventUid, userUid); err != nil {
	//	return errors.New("cant join to event")
	//}

	role := EventUser{
		UserUID:         userUid,
		EventUID:        eventUid,
		Role:            "Посетитель",
		RoleDescription: "Стандартная роль",
	}

	return putEventUser(role)
}
