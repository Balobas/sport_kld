package comments

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sport_kld/app/models"
	"time"
)

func PutPlaceComment(comment PlaceComment, executorUid models.UID) (models.UID, error) {
	if len(comment.UID) == 0 {
		comment.UID = models.UID(uuid.NewV1().String())
		comment.AuthorUid = executorUid
		comment.CreatedTime = time.Now().Unix()

		if err := putPlaceComment(comment); err != nil {
			return "", err
		}
		return comment.UID, nil
	}

	oldComment, err :=  getPlaceComment(comment.UID)
	if err != nil {
		return comment.UID, err
	}

	if !oldComment.CheckAccess(executorUid) {
		return "", errors.New("permission denied. only author can change comment")
	}

	oldComment.Text = comment.Text
	oldComment.UpdatedTime = time.Now().Unix()

	if err := putPlaceComment(oldComment); err != nil {
		return "", err
	}
	return oldComment.UID, nil
}
