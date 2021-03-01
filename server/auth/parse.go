package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
)

type Claims struct {
	jwt.StandardClaims
	UserUid string
}

func ParseToken(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserUid, nil
	}

	return "", errors.New("invalid auth token")
}