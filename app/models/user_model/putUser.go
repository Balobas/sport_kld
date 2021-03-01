package user_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putUser(user User) error {
	res, err := database.MysqlDB.NamedExec("INSERT INTO users(uid, name, login, email, age, password) VALUES (:uid, :name, :login, :email, :age, :password)", &user)
	if err != nil {
		return errors.Wrap(err, "cant insert user")
	}
	a, _ := res.RowsAffected()
	fmt.Println("inserted ", a, " rows")
	return nil
}
