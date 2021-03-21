package place_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
	"strings"
)

func GetByUID(uid models.UID) (Place, error) {
	var place Place
	if err := database.MysqlDB.Get(&place, "select * from places where uid=?", uid.String()); err != nil {
		return Place{}, errors.WithStack(err)
	}
	place.Preprocess()
	return place, nil
}

func GetPlacesByUIDs(uids []models.UID) ([]Place, []error) {
	var params []interface{}
	var conditionPart []string

	for _, uid := range uids {
		params = append(params, uid)
		conditionPart = append(conditionPart, " uid = ? ")
	}

	rows, err := database.MysqlDB.Queryx("select * from places where "+strings.Join(conditionPart, " or "), params...)
	if err != nil {
		return nil, []error{errors.Wrap(err, "cant select places by uids")}
	}

	var result []Place
	var resultErrors []error

	for rows.Next() {
		var place Place
		if err := rows.StructScan(&place); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan place"))
			continue
		}
		place.Preprocess()
		result = append(result, place)
	}

	return result, resultErrors
}
