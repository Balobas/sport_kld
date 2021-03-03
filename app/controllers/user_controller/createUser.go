package user_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/user_model"
)

func CreateUser(user user_model.User) (models.UID, string, error) {
	return user_model.CreateUser(user)
}
