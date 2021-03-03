package user_model

import (
	"github.com/dgrijalva/jwt-go/v4"
	"sport_kld/app/settings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	UserUid string
}

func (user *User) SignIn() (string, error) {
	oldUser, err := getUserByLoginAndPassword(user.Login, user.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(10000)),
			ID:        "",
			IssuedAt:  nil,
			Issuer:    "",
			NotBefore: nil,
			Subject:   "",
		},
		UserUid:       oldUser.UID.String(),
	})
	return token.SignedString(settings.SIGNING_KEY)
}
