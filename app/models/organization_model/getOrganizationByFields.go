package organization_model

import (
	"github.com/pkg/errors"
	"reflect"
	"sport_kld/database"
	"strings"
)

func GetOrganizationsByFields(fieldsNames []string, searchString string) ([]Organization, []error)  {

	//ищем организации где содержимое полей похоже на строку поиска
	var dbConditionPart []string

	numFields := reflect.TypeOf(&Organization{}).Elem().NumField()

	if len(fieldsNames) == 0 {
		for i := 0; i < numFields; i++ {
			dbConditionPart = append(dbConditionPart, " " + reflect.TypeOf(&Organization{}).Elem().Field(i).Tag.Get("db") + " like ? ")
		}
	} else {
		for i := 0; i < numFields; i++ {
			fieldTag := reflect.TypeOf(&Organization{}).Elem().Field(i).Tag.Get("json")

			for _, field := range fieldsNames {
				if field == fieldTag {
					dbConditionPart = append(dbConditionPart, " " + reflect.TypeOf(&Organization{}).Elem().Field(i).Tag.Get("db") + " like ? ")
				}
			}
		}
	}

	query := "select * from organization where " + strings.Join(dbConditionPart, " or ")

	var params []interface{}
	for i := 0; i < len(dbConditionPart); i++ {
		params = append(params, "%" + searchString + "%")
	}

	rows, err := database.MysqlDB.Queryx(query, params...)
	if err != nil {
		return []Organization{}, []error{errors.Wrap(err, "cant select organizations")}
	}

	var resultErrors []error
	var result []Organization

	for rows.Next() {
		var org Organization
		if err := rows.Scan(&org.UID, &org.Name, &org.Description); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan organization"))
			continue
		}
		org.Preprocess()
		result = append(result, org)
	}

	return result, nil
}