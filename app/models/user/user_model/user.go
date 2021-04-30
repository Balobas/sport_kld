package user_model

import (
	"errors"
	"sport_kld/app/models"
)

type User struct {
	UID      models.UID `json:"uid" db:"uid"`
	Name     string     `json:"name" db:"name"`
	Login    string     `json:"login" db:"login"`
	Email    string     `json:"email" db:"email"`
	Age      int64      `json:"age" db:"age"`
	Password string     `json:"password" db:"password"`
}

func (user *User) validate() error {
	if len(user.Name) == 0 {
		return errors.New("empty name")
	}
	if len(user.Login) == 0 {
		return errors.New("empty login")
	}
	if len(user.Password) == 0 {
		return errors.New("empty password")
	}
	if user.Age < 0 {
		return errors.New("negative age")
	}
	return nil
}
