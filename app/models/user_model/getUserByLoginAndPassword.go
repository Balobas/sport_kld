package user_model

import (
	"github.com/pkg/errors"
	"sport_kld/database"
)

func GetUserByLoginAndPassword(login, password string) (User, error) {
	var user User
	if err := database.MysqlDB.Get(&user, "select * from users where login=? and password=?", login, password); err != nil {
		return User{}, errors.Wrap(err, "cant get user")
	}

	return user, nil
}
