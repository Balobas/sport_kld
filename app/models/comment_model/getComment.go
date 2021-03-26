package comments

import (
	"sport_kld/app/models"
	"sport_kld/database"
)

func getPlaceComment(uid models.UID) (PlaceComment, error) {
	var comment PlaceComment
	if err := database.MysqlDB.Get(&comment, "select * from place_comments where uid = ?", uid); err != nil {
		return PlaceComment{}, err
	}
	return comment, nil
}

func getOrganizationComment(uid models.UID) (OrganizationComment, error) {
	var comment OrganizationComment
	if err := database.MysqlDB.Get(&comment, "select * from organization_comments where uid = ?", uid); err != nil {
		return OrganizationComment{}, err
	}
	return comment, nil
}
