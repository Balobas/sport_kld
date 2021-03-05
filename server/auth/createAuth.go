package auth

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"sport_kld/app/controllers/user_controller"
	"sport_kld/database"
	"sport_kld/server/auth/token"
)

func CreateAuth(userUid string, td *token.TokenDetails) error {
	if _, err := user_controller.GetUserByUid(userUid); err != nil {
		return errors.Wrap(err, "user dont exist")
	}
	authUid := uuid.NewV4().String()
	auth := Auth{
		Uid:        authUid,
		UserUid:    userUid,
		AccessUid:  td.AccessUid,
		RefreshUid: td.RefreshUid,
	}

	access := token.Access{
		AccessUid: td.AccessUid,
		UserUid:   userUid,
		Expires:   td.AtExpires,
	}

	refresh := token.Refresh{
		RefreshUid: td.RefreshUid,
		UserUid:    userUid,
		Expires:    td.RtExpires,
	}

	if _, err := database.MysqlDB.NamedExec("insert into auth(uid, userUid, accessUid, refreshUid) values (:uid, :userUid, :accessUid, :refreshUid)", &auth); err != nil {
		return errors.Wrap(err, "cant create auth")
	}
	if _, err := database.MysqlDB.NamedExec("insert into access(accessUid, userUid, expires) values (:uid, :userUid, :expires)", &access); err != nil {
		return errors.Wrap(err, "cant create auth")
	}
	if _, err := database.MysqlDB.NamedExec("insert into refresh(refreshUid, userUid, expires) values (:uid, :userUid, :expires)", &refresh); err != nil {
		return errors.Wrap(err, "cant create auth")
	}
	return nil
}
