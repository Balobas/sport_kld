package event_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putEventUser(role EventUser) error {
	_, err := GetEventUserByUid(role.EventUID, role.UserUID)
	if err == nil {
		if result, err := database.MysqlDB.NamedExec("update event_user_roles set role = :role, roleDescription = :roleDescription where userUid= :userUid and eventUid = :eventUid", &role); err != nil {
			return errors.Wrap(err, "cant update role")
		} else {
			// Заменить на логер
			a, _ := result.RowsAffected()
			fmt.Println("inserted ", a, " rows")
		}
		return nil
	}

	if result, err := database.MysqlDB.NamedExec("insert into event_user_roles(userUid, eventUid, role, roleDescription) values (:userUid, :eventUid, :role, :roleDescription)", &role); err != nil {
		return errors.Wrap(err, "cant create role")
	} else {
		// Заменить на логер
		a, _ := result.RowsAffected()
		fmt.Println("inserted ", a, " rows")
	}
	return nil
}
