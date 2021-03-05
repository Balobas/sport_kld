package auth

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/server/auth/token"
	"time"
)

func VerifyAuth(accessToken string) (authUid, userUid string, err error) {
	access, err := token.ParseAccessToken(accessToken)
	if err != nil {
		return "", "", err
	}
	if access.Expires < time.Now().Unix() {
		return "", "", errors.New("access expires")
	}

	storedAccess, err := getAccess(access.AccessUid)
	if err != nil {
		return "", "", errors.Wrap(err, "fake token")
	}

	if storedAccess.Expires != access.Expires {
		fmt.Println("secret key was spizhen")
		return "", "", errors.New("fake token")
	}
	if storedAccess.UserUid != access.UserUid {
		fmt.Println("secret key was spizhen")
		return "", "", errors.New("fake token")
	}

	userUid = access.UserUid
	authUid, err = GetAuthUidByAccessUid(access.AccessUid)
	return
}
