package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"sport_kld/app/models/user_model"
	"sport_kld/app/settings"
)

type UserParams struct {
	Login string
	Password string
}

func SignIn(userParams UserParams) (string, error) {

	pwd := sha1.New()
	pwd.Write([]byte(userParams.Password))
	pwd.Write([]byte("saaaaalt"))

	password := fmt.Sprintf("%x", pwd.Sum(nil))

	if _, err := user_model.GetUserByLoginAndPassword(userParams.Login, password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ,
			ID:        "",
			IssuedAt:  nil,
			Issuer:    "",
			NotBefore: nil,
			Subject:   "",
		},
		Username:       userParams.Login,
	})
	return token.SignedString(settings.SIGNING_KEY)
}
