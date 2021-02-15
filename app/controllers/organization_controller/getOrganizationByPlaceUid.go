package organization_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/organization_model"
)

func GetOrganizationByPlaceUID(placeUid string) ([]organization_model.Organization, []error) {
	return organization_model.GetOrganizationByPlaceUID(models.UID(placeUid))
}
