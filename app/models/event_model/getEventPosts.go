package event_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetEventInfoPosts(eventUid models.UID) ([]EventInfoPost, []error) {
	rows, err := database.MysqlDB.Queryx("select * from event_posts where eventUid = ?", eventUid)
	if err != nil {
		return []EventInfoPost{}, []error{errors.WithStack(err)}
	}

	var posts []EventInfoPost
	var resultErrors []error

	for rows.Next() {
		var post EventInfoPost
		if err := rows.StructScan(&post); err != nil {
			resultErrors = append(resultErrors, err)
			continue
		}

		posts = append(posts, post)
	}

	return posts, resultErrors
}
