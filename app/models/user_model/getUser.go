package user_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetUserByUid(uid models.UID) (User, error) {
	var user User
	if err := database.MysqlDB.Get(&user, "select * from users where uid=?", uid); err != nil {
		return User{}, errors.Wrap(err, "cant select user")
	}
	user.Password = ""
	return user, nil
}
