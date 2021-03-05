package auth

import (
	"errors"
	"sport_kld/app/controllers/user_controller"
	"sport_kld/app/models/user_model"
	"sport_kld/server/auth/token"
)

func Login(user user_model.User) (string, string, error) {
	// check user in db
	if err := user_controller.VerifyUser(user); err != nil {
		return "", "", err
	}

	// Пока допускается только один активный клиент для пользователя
	_, err := getAuthByUserUid(user.UID.String())
	if err == nil {
		return "", "", errors.New("user already login")
	}

	td, err := token.CreateToken(user.UID.String())
	if err != nil {
		return "", "", err
	}

	if err := CreateAuth(user.UID.String(), td); err != nil {
		return "", "", err
	}

	return td.AccessToken, td.RefreshToken, nil
}
