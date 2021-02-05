package place

import (
	"../../models"
	"fmt"
	"github.com/pkg/errors"
	"../../../database"
)

func GetPlaceByUID(uid models.UID) (models.Place, error) {
	var place models.Place
	fmt.Println(place)
	if err := database.MysqlDB.Get(&place, "select * from places where uid=?", uid.String()); err != nil {
		return models.Place{}, errors.WithStack(err)
	}
	return place, nil
}

func GetPlacesByUIDs(uids []models.UID) ([]models.Place, []error) {
	var places []models.Place
	var errs []error
	for _, uid := range uids {
		place, err := GetPlaceByUID(uid)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "place with uid: " + uid.String() + "  "))
		} else {
			places = append(places, place)
		}
	}
	return places, errs
}
