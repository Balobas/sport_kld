package place_model

import (
	"../../../database"
	"../../models"
	"github.com/pkg/errors"
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

	rows, err := database.MysqlDB.Queryx("select * from places where " + strings.Join(conditionPart, " or "), params...)
	if err != nil {
		return nil, []error{errors.Wrap(err, "cant select places by uids")}
	}

	var result []Place
	var resultErrors []error

	for rows.Next() {
		var place Place
		if err := rows.Scan(&place.UID, &place.Name, &place.BuildingName,
			&place.BuildingType, &place.Description,
			&place.Address, &place.City,
			&place.OpeningHours, &place.PostIndex,
			&place.WebSite, &place.Phones, &place.Email,
			&place.Facebook, &place.Instagram, &place.Twitter, &place.VK);
		err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan place"))
			continue
		}
		place.Preprocess()
		result = append(result, place)
	}

	return result, resultErrors
}


