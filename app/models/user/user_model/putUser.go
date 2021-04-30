package user_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putUser(user User) error {

	if _, err := GetUserByUid(user.UID); err == nil {
		res, err := database.MysqlDB.NamedExec("UPDATE users SET name = :name, login = :login, email = :email, age = :age Where uid = :uid", &user)
		if err != nil {
			return errors.Wrap(err, "cant insert user")
		}
		a, _ := res.RowsAffected()
		fmt.Println("updated ", a, " rows")
		return nil
	}

	res, err := database.MysqlDB.NamedExec("INSERT INTO users(uid, name, login, email, age, password) VALUES (:uid, :name, :login, :email, :age, :password)", &user)
	if err != nil {
		return errors.Wrap(err, "cant insert user")
	}
	a, _ := res.RowsAffected()
	fmt.Println("inserted ", a, " rows")
	return nil
}
