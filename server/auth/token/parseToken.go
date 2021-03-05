package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
)

func ParseAccessToken(accessToken string) (Access, error) {
	tk, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("asdabjaghsda"), nil
	})
	if err != nil {
		return Access{}, err
	}

	claims, ok := tk.Claims.(*jwt.MapClaims);
	if !ok {
		return Access{}, errors.New("cant parse claims, invalid token")
	}
	if !tk.Valid {
		return Access{}, errors.New("invalid token")
	}

	accessUid, ok := (*claims)["access_uid"].(string)
	if !ok {
		return Access{}, errors.New("invalid token")
	}

	userUid, ok := (*claims)["user_uid"].(string)
	if !ok {
		return Access{}, errors.New("invalid token")
	}

	expires, ok := (*claims)["expires"].(int64)
	if !ok {
		return Access{}, errors.New("invalid token")
	}

	return Access{
		AccessUid: accessUid,
		UserUid:   userUid,
		Expires:   expires,
	}, nil
}

func ParseRefreshToken(refreshToken string) (Refresh, error) {
	tk, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ianfijnewrnesef"), nil
	})
	if err != nil {
		return Refresh{}, err
	}

	claims, ok := tk.Claims.(*jwt.MapClaims);
	if !ok {
		return Refresh{}, errors.New("cant parse claims, invalid token")
	}
	if !tk.Valid {
		return Refresh{}, errors.New("invalid token")
	}

	refreshUid, ok := (*claims)["refresh_uid"].(string)
	if !ok {
		return Refresh{}, errors.New("invalid token")
	}

	userUid, ok := (*claims)["user_uid"].(string)
	if !ok {
		return Refresh{}, errors.New("invalid token")
	}

	expires, ok := (*claims)["expires"].(int64)
	if !ok {
		return Refresh{}, errors.New("invalid token")
	}

	return Refresh{
		RefreshUid: refreshUid,
		UserUid:   userUid,
		Expires:   expires,
	}, nil
}
