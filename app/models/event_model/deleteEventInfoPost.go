package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func DeleteEventInfoPost(postUid, executorUid models.UID) error {
	post, err := GetEventInfoPost(postUid)
	if err != nil {
		return errors.Wrap(err, "cant get event info post")
	}
	if post.AuthorUID != executorUid {
		event, err := GetEventByUid(post.EventUID)
		if err != nil {
			return errors.New("cant get event")
		}
		if event.CreatorUID != executorUid {
			return errors.New("permission denied")
		}
	}

	if _, err := database.MysqlDB.Exec("delete from event_posts where uid = ?", postUid); err != nil {
		return err
	}
	return nil
}
