package user_model

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"sport_kld/app/models"
	"strings"
)

func CreateUser(user User) (models.UID, error) {
	//валидация
	if err := user.validate(); err != nil {
		return "", errors.Wrap(err, "cant create user")
	}

	// проверка занят ли логин
	_, err := GetUserByLogin(user.Login)
	if err == nil {
		return "", errors.New("login exist")
	} else if !strings.Contains(err.Error(), "sql: no rows in result set") {
		return "", err
	}

	// генерация uid
	user.UID = models.UID(uuid.NewV1().String())

	h := md5.New()
	h.Write([]byte(user.Password))
	h.Write([]byte("hashsalt"))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	// Добавление
	if err := putUser(user); err != nil {
		return "", errors.Wrap(err, "cant create user")
	}

	return user.UID, nil
}
