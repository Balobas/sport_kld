package user_model

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"sport_kld/app/models"
	"strings"
)

func CreateUser(user User) (models.UID, string, error) {
	//валидация
	if err := user.validate(); err != nil {
		return "", "", errors.Wrap(err, "cant create user")
	}

	// проверка занят ли логин
	_, err := GetUserByLogin(user.Login)
	if err == nil {
		return "", "", errors.New("login exist")
	} else if !strings.Contains(err.Error(), "sql: no rows in result set") {
		return "", "", err
	}

	// генерация uid
	user.UID = models.UID(uuid.NewV1().String())

	if err := putUser(user); err != nil {
		return "", "", errors.Wrap(err, "cant create user")
	}

	//генерация jwt токена
	token, err := user.GenerateNewToken()
	if err != nil {
		return "", "", err
	}

	//добавление
	return user.UID, token, putUser(user)
}
