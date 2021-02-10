package place_model

import (
	"sport_kld/database"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func GetPlacesByFields(fieldsNames []string, searchString string) ([]Place, []error)  {

	//ищем места где содержимое полей похоже на строку поиска
	var dbConditionPart []string

	numFields := reflect.TypeOf(&Place{}).Elem().NumField()

	if len(fieldsNames) == 0 {
		for i := 0; i < numFields; i++ {
			dbConditionPart = append(dbConditionPart, " " + reflect.TypeOf(&Place{}).Elem().Field(i).Tag.Get("db") + " like ? ")
		}
	} else {
		for i := 0; i < numFields; i++ {
			fieldTag := reflect.TypeOf(&Place{}).Elem().Field(i).Tag.Get("json")

			for _, field := range fieldsNames {
				if field == fieldTag {
					dbConditionPart = append(dbConditionPart, " " + reflect.TypeOf(&Place{}).Elem().Field(i).Tag.Get("db") + " like ? ")
				}
			}
		}
	}

	query := "select * from places where " + strings.Join(dbConditionPart, " or ")

	var params []interface{}
	for i := 0; i < len(dbConditionPart); i++ {
		params = append(params, "%" + searchString + "%")
	}

	rows, err := database.MysqlDB.Queryx(query, params...)
	if err != nil {
		return []Place{}, []error{errors.Wrap(err, "cant select places")}
	}

	var resultErrors []error
	var result []Place

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
	return result, nil
}
