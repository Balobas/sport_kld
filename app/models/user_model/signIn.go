package user_model

import (
	"github.com/pkg/errors"
)

func SignIn(user User) (string, error) {
	oldUser, err := getUserByLoginAndPassword(user.Login, user.Password)
	if err != nil {
		return "", errors.Wrap(err, "cant sign in")
	}
	return oldUser.GenerateNewToken()
}
