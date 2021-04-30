package user_model

import (
	"github.com/pkg/errors"
	"sport_kld/database"
)

func GetUserByLogin(login string) (User, error) {
	var user User
	if err := database.MysqlDB.Get(&user, "select * from users where login=?", login); err != nil {
		return User{}, errors.Wrap(err, "cant select user")
	}
	user.Password = ""
	return user, nil
}
