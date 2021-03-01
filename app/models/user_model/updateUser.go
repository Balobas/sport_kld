package user_model

import (
	"github.com/pkg/errors"
	"strings"
)

func UpdateUser(user User) error {
	if err := user.validate(); err != nil {
		return errors.Wrap(err, "invalid user")
	}

	oldUser, err := GetUserByUid(user.UID)
	if err != nil {
		return errors.Wrap(err, "cant update user")
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
