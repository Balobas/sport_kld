package auth

import (
	"github.com/pkg/errors"
	"sport_kld/database"
)

func DeleteAuth(authUid string) error {
	auth, err := getAuth(authUid)
	if err != nil {
		return errors.WithStack(err)
	}
	if _, err := database.MysqlDB.Exec("delete from auth where uid = ?", authUid); err != nil {
		return errors.Wrap(err, "cant delete auth")
	}
	if _, err := database.MysqlDB.Exec("delete from access where accessUid = ?", auth.AccessUid); err != nil {
		return errors.Wrap(err, "cant delete auth")
	}
	if _, err := database.MysqlDB.Exec("delete from refresh where refreshUid = ?", auth.RefreshUid); err != nil {
		return errors.Wrap(err, "cant delete auth")
	}
	return nil
}
