package event_model

import (
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventInfoPost(postUid models.UID) (EventInfoPost, error) {
	var post EventInfoPost
	if err := database.MysqlDB.Get(&post, "select * from event_posts where uid = ?", postUid); err != nil {
		return EventInfoPost{}, err
	}
	return post, nil
}
