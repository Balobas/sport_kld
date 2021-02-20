package place_controller

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/app/models/place_model"
)

func GetPlaceByUID(uid string) (place_model.Place, error) {
	return place_model.GetByUID(models.UID(uid))
}

func GetPlacesByUIDs(uids []string) ([]place_model.Place, []error) {
	var places []place_model.Place
	var errs []error
	for _, uid := range uids {
		place, err := GetPlaceByUID(uid)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "place with uid: "+uid+"  "))
		} else {
			places = append(places, place)
		}
	}
	return places, errs
}
