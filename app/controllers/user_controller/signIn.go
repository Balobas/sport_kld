package user_controller

import "sport_kld/app/models/user_model"

func SignIn(user user_model.User) (string, error) {
	return user_model.SignIn(user)
}
