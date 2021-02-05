package place_model

import (
	"../../../database"
	"../../models"
	"github.com/pkg/errors"
)

func (place *Place) GetByUID(uid models.UID) error {
	if err := database.MysqlDB.Get(place, "select * from places where uid=?", uid.String()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
