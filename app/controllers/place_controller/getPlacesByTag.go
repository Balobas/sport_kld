package place_controller

import (
	"sport_kld/app/models/place_model"
)

func GetPlacesByTags(searchString string) ([]place_model.Place, []error) {
	return place_model.GetPlacesByTags(searchString)
}
