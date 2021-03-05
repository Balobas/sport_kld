package user_controller

import "sport_kld/app/models/user_model"

func VerifyUser(user user_model.User) error {
	return user_model.VerifyUser(user)
}
