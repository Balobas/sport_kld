package user_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
	"strings"
)

func UpdateUser(user User, executorUid models.UID) error {
	if user.UID != executorUid {
		return errors.New("Permission denied")
	}
	if err := user.validate(); err != nil {
		return errors.Wrap(err, "invalid user")
	}

	var oldUser User
	if err := database.MysqlDB.Get(&oldUser, "select * from users where uid=?", user.UID); err != nil {
		return errors.Wrap(err, "cant get user")
	}

	if user.Login != oldUser.Login {
		if _, err := GetUserByLogin(user.Login); err == nil {
			return errors.New("login exist")
		} else if !strings.Contains(err.Error(), "sql: no rows in result set") {
			return err
		}
		oldUser.Login = user.Login
	}

	oldUser.Name = user.Name
	oldUser.Age = user.Age
	oldUser.Email = user.Email

	return putUser(oldUser)
}
