package user_controller

import "sport_kld/app/models/user_model"

func UpdateUser(user user_model.User) error {
	return user_model.UpdateUser(user)
}
