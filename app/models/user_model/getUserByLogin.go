package user_model

import "sport_kld/database"

func GetUserByLogin(login string) (User, error) {
	var user User
	if err := database.MysqlDB.Get(&user, "select * from users where login=?", login); err != nil {
		return User{}, err
	}
	return user, nil
}
