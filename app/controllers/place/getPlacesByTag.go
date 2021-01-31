package place

import (
	"../../../models"
)

func GetPlacesByTags(searchString string) ([]models.Place, error) {

	var params []string
	//ищем похожие теги
	q := "select * from tags where name like '%?%' "
	params = append(params, searchString)
	if len(searchString) > 3 {
		q += " or name like '%?%' "
		params = append(params, searchString[:len(searchString) - 2])
	}

	//database.MysqlDB.Get()


	//по тегам ищем места


	return nil, nil
}
