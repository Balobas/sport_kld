package comments

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"time"
)

func DeletePlaceComment(commentUid, executorUid models.UID) error {
	comment, err := getPlaceComment(commentUid)
	if err != nil {
		return errors.Wrap(err, "cant delete comment")
	}

	if !comment.CheckAccess(executorUid) {
		return errors.New("permission denied")
	}

	comment.IsDeleted = true
	comment.UpdatedTime = time.Now().Unix()

	return putPlaceComment(comment)
}

func DeleteOrganizationComment(commentUid, executorUid models.UID) error {
	comment, err := getOrganizationComment(commentUid)
	if err != nil {
		return errors.Wrap(err, "cant delete comment")
	}

	if !comment.CheckAccess(executorUid) {
		return errors.New("permission denied")
	}

	comment.IsDeleted = true
	comment.UpdatedTime = time.Now().Unix()

	return putOrganizationComment(comment)
}
