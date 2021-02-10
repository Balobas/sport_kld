package place_controller

import (
	"sport_kld/app/models/place_model"
)

func GetPlacesByFields(fieldsNames []string, searchString string) ([]place_model.Place, []error)  {
	return place_model.GetPlacesByFields(fieldsNames, searchString)
}
