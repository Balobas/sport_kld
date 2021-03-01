package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventUserByUid(eventUid, userUid models.UID) (EventUser, error) {
	var role EventUser
	if err := database.MysqlDB.Get(&role, "select * from event_user_roles where userUid=? and eventUid=?", userUid, eventUid); err != nil {
		return EventUser{}, errors.Wrap(err, "cant get user role")
	}
	return role, nil
}
