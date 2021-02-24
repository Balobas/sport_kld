package user_model

import "sport_kld/app/models"

type User struct {
	UID      models.UID `json:"uid" db:"uid"`
	Name     string     `json:"name" db:"name"`
	Login    string     `json:"login" db:"login"`
	Password string     `json:"password" db:"password"`
}
