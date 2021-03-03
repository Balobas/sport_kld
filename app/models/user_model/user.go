package user_model

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"sport_kld/app/models"
	"sport_kld/app/settings"
	"time"
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

type Claims struct {
	jwt.StandardClaims
	UserUid string
}

func (user *User) GenerateNewToken() (string, error) {
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