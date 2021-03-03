package user_controller

import "sport_kld/app/models/user_model"

func GetUserByLogin(login string) (user_model.User, error) {
	return user_model.GetUserByLogin(login)
}
