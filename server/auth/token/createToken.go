package token

import (
	"github.com/dgrijalva/jwt-go/v4"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	UserUid string
	AccessUid string
}

func CreateToken(userUid string) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute*30).Unix()
	td.AccessUid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 2).Unix()
	td.RefreshUid = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["access_uid"] = td.AccessUid
	atClaims["user_uid"] = userUid
	atClaims["exp"] = td.AtExpires

	var err error

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte("asdabjaghsda"))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uid"] = td.RefreshUid
	rtClaims["user_uid"] = userUid
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte("ianfijnewrnesef"))
	if err != nil {
		return nil, err
	}
	return td, nil
}
