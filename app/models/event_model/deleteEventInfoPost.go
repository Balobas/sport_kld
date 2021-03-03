package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func DeleteEventInfoPost(postUid models.UID) error {
	if _, err := GetEventInfoPost(postUid); err != nil {
		return errors.Wrap(err, "cant get event info post")
	}

	if _, err := database.MysqlDB.Exec("delete from event_posts where uid = ?", postUid); err != nil {
		return err
	}
	return nil
}
