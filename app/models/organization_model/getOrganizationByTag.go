package organization_model

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
	"strings"
)

func GetOrganizationsByTags(searchString string) ([]Organization, []error) {
	if len(strings.TrimLeft(strings.TrimRight(searchString, " "), " ")) == 0 {
		return nil, []error{errors.New("empty search string")}
	}

	//ищем похожие теги
	// TODO: задуматься о переносе этого функционала в модель тега
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
		return []Organization{}, nil
	}

	//по тегам ищем места находим все uid подходящих организаций
	var wherePart []string
	var uids []interface{}

	for _, tag := range tags {
		wherePart = append(wherePart, " tag_uid=? ")
		uids = append(uids, tag.UID.String())
	}

	rows, err = database.MysqlDB.Queryx("select * from organization_tags where "+strings.Join(wherePart, " or "), uids...)
	if err != nil {
		return nil, append(resultErrors, errors.Wrap(err, "cant select organizations uids"))
	}

	var orgsUids []interface{}
	for rows.Next() {
		var uid, t string

		if err := rows.Scan(&uid, &t); err != nil {
			resultErrors = append(resultErrors, errors.Wrap(err, "cant scan organization uid"))
			continue
		}
		orgsUids = append(orgsUids, uid)
	}

	// забираем все подходящие организации из базы
	wherePart = []string{}
	for i := 0; i < len(orgsUids); i++ {
		wherePart = append(wherePart, " uid=? ")
	}

	rows, err = database.MysqlDB.Queryx("select * from organizations where "+strings.Join(wherePart, " or "), orgsUids...)
	if err != nil {
		return nil, append(resultErrors, errors.Wrap(err, "cant select organizations"))
	}

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

	return result, resultErrors
}
