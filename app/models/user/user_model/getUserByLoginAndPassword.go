package user_model

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func getUserByLoginAndPassword(login, password string) (User, error) {
	h := md5.New()
	h.Write([]byte(password))
	h.Write([]byte("hashsalt"))
	password = fmt.Sprintf("%x", h.Sum(nil))

	var user User
	if err := database.MysqlDB.Get(&user, "select * from users where login=? and password=?", login, password); err != nil {
		return User{}, errors.Wrap(err, "cant select user")
	}

	return user, nil
}
