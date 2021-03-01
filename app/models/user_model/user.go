package user_model

import "sport_kld/app/models"

type User struct {
	UID      models.UID `json:"uid" db:"uid"`
	Name     string     `json:"name" db:"name"`
	Login    string     `json:"login" db:"login"`
	Email    string     `json:"email" db:"email"`
	Age      int64      `json:"age" db:"age"`
	Password string     `json:"password" db:"password"`
}
