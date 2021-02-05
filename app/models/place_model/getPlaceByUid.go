package place_model

import (
	"../../../database"
	"../../models"
	"github.com/pkg/errors"
)

func GetByUID(uid models.UID) (Place, error) {
	var place Place
	if err := database.MysqlDB.Get(&place, "select * from places where uid=?", uid.String()); err != nil {
		return Place{}, errors.WithStack(err)
	}
	return place, nil
}


