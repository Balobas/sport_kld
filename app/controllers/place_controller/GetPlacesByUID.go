package place_controller

import (
	"../../models"
	"../../models/place_model"
	"github.com/pkg/errors"
)

func GetPlaceByUID(uid models.UID) (place_model.Place, error) {
	place, err := place_model.GetByUID(uid)
	if err != nil {
		return place_model.Place{}, errors.WithStack(err)
	}
	return place, nil
}

func GetPlacesByUIDs(uids []models.UID) ([]place_model.Place, []error) {
	var places []place_model.Place
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
