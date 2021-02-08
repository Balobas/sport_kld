package place_controller

import (
	"../../models"
	"../../models/place_model"
)


func GetPlacesByOrganizationUID(organizationUid string) ([]place_model.Place, []error) {
	return place_model.GetPlacesByOrganizationUID(models.UID(organizationUid))
}
