package organization_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetOrganizationByPlaceUID(placeUid models.UID) ([]Organization, []error) {

	var orgsUids []models.UID
	rows, err := database.MysqlDB.Queryx("select (organization_uid) from organization_places where place_uid = ?", placeUid)
	if err != nil {
		return []Organization{}, []error{errors.Wrap(err, "cant select organizations uids")}
	}

	var resultErrors []error

	for rows.Next() {
		var uid models.UID
		if err := rows.Scan(&uid); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan organization uid"))
			continue
		}
		orgsUids = append(orgsUids, uid)
	}

	return GetOrganizationsByUIDs(orgsUids)
}
