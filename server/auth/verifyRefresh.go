package auth

import (
	"github.com/pkg/errors"
	"sport_kld/server/auth/token"
	"time"
)

func VerifyRefresh(refreshToken string) (string, string, error) {
	refresh, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}
	if refresh.Expires < time.Now().Unix() {
		return "", "", errors.New("token expires")
	}

	storedRefresh, err := getRefresh(refresh.RefreshUid)
	if err != nil {
		return "", "", errors.Wrap(err, "fake token")
	}

	if storedRefresh.UserUid != refresh.UserUid {
		return "", "", errors.New("fake token")
	}
	if storedRefresh.Expires != refresh.Expires {
		return "", "", errors.New("fake token")
	}

	authUid, err := GetAuthUidByRefreshUid(refresh.RefreshUid)
	return authUid, refresh.UserUid, err
}
