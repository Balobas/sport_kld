package place_model

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
	"strings"
)

func GetPlacesByTags(searchString string) ([]Place, []error) {
	if len(strings.TrimLeft(strings.TrimRight(searchString, " "), " ")) == 0 {
		return nil, []error{errors.New("empty search string")}
	}

	//ищем похожие теги
	var tags []models.Tag
	var rows *sqlx.Rows

	var err error

	if len(searchString) > 3 {
		rows, err = database.MysqlDB.Queryx("select * from tags where name like ? or name like ? ", "%"+searchString+"%", "%"+searchString[:len(searchString)-2]+"%")
		//err = database.MysqlDB.Get(&tag, "select * from tags where name like \"%" + searchString + "%\" or name like \"% " + searchString[:len(searchString) - 2] + " %\"")
	} else {
		rows, err = database.MysqlDB.Queryx("select * from tags where name like ? ", "%"+searchString+"%")
	}

	if err != nil {
		return nil, []error{errors.WithStack(err)}
	}

	var resultErrors []error

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.UID, &tag.Name); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan tag"))
			continue
		}
		tags = append(tags, tag)
	}

	if len(tags) == 0 {
		return []Place{}, nil
	}

	//по тегам ищем места находим все uid подходящих мест
	var wherePart []string
	var uids []interface{}

	for _, tag := range tags {
		wherePart = append(wherePart, " tag_uid=? ")
		uids = append(uids, tag.UID.String())
	}

	rows, err = database.MysqlDB.Queryx("select * from places_tags where "+strings.Join(wherePart, " or "), uids...)
	if err != nil {
		return nil, append(resultErrors, errors.Wrap(err, "cant select places uids"))
	}

	var placesUids []interface{}
	for rows.Next() {
		var uid, t string

		if err := rows.Scan(&uid, &t); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan place uid"))
			continue
		}
		placesUids = append(placesUids, uid)
	}

	// забираем все подходящие места из базы
	wherePart = []string{}
	for i := 0; i < len(placesUids); i++ {
		wherePart = append(wherePart, " uid=? ")
	}

	rows, err = database.MysqlDB.Queryx("select * from places where "+strings.Join(wherePart, " or "), placesUids...)
	if err != nil {
		return nil, append(resultErrors, errors.Wrap(err, "cant select places"))
	}

	var result []Place
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
