package organization_controller

import (
	"sport_kld/app/models"
	"sport_kld/app/models/organization_model"
)

func GetOrganizationByUID(uid string) (organization_model.Organization, error) {
	return organization_model.GetOrganizationByUID(models.UID(uid))
}

func GetOrganizationsByUIDs(uids []string) ([]organization_model.Organization, []error) {
	var params []models.UID
	for _, uid := range uids {
		params = append(params, models.UID(uid))
	}
	return organization_model.GetOrganizationsByUIDs(params)
}
