package comment_model

import (
	"fmt"
	"github.com/pkg/errors"
	"sport_kld/database"
)

func putPlaceComment(comment PlaceComment) error {
	_, err := getPlaceComment(comment.UID)
	if err == nil {
		result, err := database.MysqlDB.NamedExec("update place_comments set text = :text, updatedTime = :updatedTime, isDeleted = :isDeleted, where uid = :uid", &comment)
		if err != nil {
			return errors.Wrap(err, "cant update place ")
		}

		a, _ := result.RowsAffected()
		fmt.Println("inserted ", a, " rows")

		return nil
	}

	result, err := database.MysqlDB.NamedExec("insert into place_comments(uid, text, authorUid, createdTime, updatedTime, isDeleted, placeUid) values (:uid, :text, :authorUid, :createdTime, :updatedTime, :isDeleted, :placeUid)", &comment)
	if err != nil {
		return errors.Wrap(err, "cant create place comment")
	}

	// Заменить на логер
	a, _ := result.RowsAffected()
	fmt.Println("inserted ", a, " rows")

	return err
}

func putOrganizationComment(comment OrganizationComment) error {
	_, err := getOrganizationComment(comment.UID)
	if err == nil {
		result, err := database.MysqlDB.NamedExec("update organization_comments set text = :text, updatedTime = :updatedTime, isDeleted = :isDeleted where uid = :uid", &comment)
		if err != nil {
			return errors.Wrap(err, "cant update place ")
		}

		a, _ := result.RowsAffected()
		fmt.Println("inserted ", a, " rows")

		return nil
	}

	result, err := database.MysqlDB.NamedExec("insert into organization_comments(uid, text, authorUid, createdTime, updatedTime, isDeleted, organizationUid) values (:uid, :text, :authorUid, :createdTime, :updatedTime, :isDeleted, :organizationUid)", &comment)
	if err != nil {
		return errors.Wrap(err, "cant create place comment")
	}

	// Заменить на логер
	a, _ := result.RowsAffected()
	fmt.Println("inserted ", a, " rows")

	return err
}

