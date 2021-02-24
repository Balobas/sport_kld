package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventUserRoleByUid(eventUid, userUid models.UID) (EventUserRole, error) {
	var role EventUserRole
	if err := database.MysqlDB.Get(&role, "select * from event_user_roles where userUid=? and eventUid=?", userUid, eventUid); err != nil {
		return EventUserRole{}, errors.Wrap(err, "cant get user role")
	}
	return role, nil
}
