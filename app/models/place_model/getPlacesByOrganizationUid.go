package place_model


import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetPlacesByOrganizationUID(orgUid models.UID) ([]Place, []error) {

	var placesUids []models.UID
	rows, err := database.MysqlDB.Queryx("select (place_uid) from organization_places where organization_uid = ?", orgUid)
	if err != nil {
		return []Place{}, []error{errors.Wrap(err, "cant select places uids")}
	}

	var resultErrors []error

	for rows.Next() {
		var uid models.UID
		if err := rows.Scan(&uid); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan place uid"))
			continue
		}
		placesUids = append(placesUids, uid)
	}

	return GetPlacesByUIDs(placesUids)
}
