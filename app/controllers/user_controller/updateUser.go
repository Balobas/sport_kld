package user_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/user_model"
)

func UpdateUser(user user_model.User, executorUid string) error {
	return user_model.UpdateUser(user, models.UID(executorUid))
}
