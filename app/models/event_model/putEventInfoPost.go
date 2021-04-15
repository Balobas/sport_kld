package event_model

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"sport_kld/app/models"
	"sport_kld/database"
)

func PutEventInfoPost(post EventInfoPost) (models.UID, error) {
	event, err := GetEventByUid(post.EventUID)
	if err != nil {
		return "", err
	}
	if event.CreatorUID != post.AuthorUID {
		return "", errors.New("access denied")
	}

	if len(post.Text) == 0 {
		return "", errors.New("post text is empty")
	}

	if len(post.UID) == 0 {
		post.UID = models.UID(uuid.NewV1().String())
		/*
			post.CreatedTime =
			post.UpdatedTime =
		*/
		_, err := database.MysqlDB.NamedExec("INSERT into event_posts(uid, eventUid, authorUid, text, createdTime, updatedTime) VALUES (:uid, :eventUid, :authorUid, :text, :createdTime, :updatedTime)", &post)
		if err != nil {
			return "", errors.Wrap(err, "cant create event post")
		}

		return post.UID, nil
	}

	/*
		post.UpdatedTime =
	*/

	if _, err := database.MysqlDB.NamedExec("update event_posts set text = :text where uid = :uid", &post); err != nil {
		return "", errors.Wrap(err, "cant update event post")
	}

	return post.UID, nil
}
