package auth

import (
	"github.com/pkg/errors"
	"sport_kld/database"
	"sport_kld/server/auth/token"
)

func getAuth(authUid string) (Auth, error) {
	var auth Auth
	if err := database.MysqlDB.Get(&auth, "select * from auth where uid = ?", authUid); err != nil {
		return Auth{}, errors.Wrap(err, "cant get auth")
	}
	return auth, nil
}

func getAccess(uid string) (token.Access, error) {
	var access token.Access
	if err := database.MysqlDB.Get(&access, "select * from access where accessUid = ?", uid); err != nil {
		return token.Access{}, errors.Wrap(err, "cant get access")
	}
	return access, nil
}

func getRefresh(uid string) (token.Refresh, error) {
	var refresh token.Refresh
	if err := database.MysqlDB.Get(&refresh, "select * from refresh where refreshUid = ?", uid); err != nil {
		return token.Refresh{}, errors.Wrap(err, "cant get refresh")
	}
	return refresh, nil
}

func getAuthByUserUid(uid string) (Auth, error) {
	var auth Auth
	if err := database.MysqlDB.Get(&auth, "select * from auth where userUid = ?", uid); err != nil {
		return Auth{}, errors.Wrap(err, "cant get auth")
	}
	return auth, nil
}

func AuthExist(userUid string) error {
	_, err := getAuthByUserUid(userUid)
	return err
}

func GetAuthUidByAccessUid(uid string) (string, error) {
	var authUid string
	if err := database.MysqlDB.Get(&authUid, "select uid from auth where accessUid = ?", uid); err != nil {
		return "", errors.Wrap(err, "cant get authUid")
	}
	return authUid, nil
}

func GetAuthUidByRefreshUid(uid string) (string, error) {
	var authUid string
	if err := database.MysqlDB.Get(&authUid, "select uid from auth where refreshUid = ?", uid); err != nil {
		return "", errors.Wrap(err, "cant get authUid")
	}
	return authUid, nil
}
