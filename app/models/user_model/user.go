package user_model

import "sport_kld/app/models"

type User struct {
	UID models.UID
	Name string
	Login string

}