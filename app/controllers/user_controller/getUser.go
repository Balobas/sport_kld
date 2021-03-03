package user_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/user_model"
)

func GetUserByUid(uid string) (user_model.User, error) {
	return user_model.GetUserByUid(models.UID(uid))
}
