package organization_controller

import "sport_kld/app/models/organization_model"

func GetOrganizationsByTags(searchString string) ([]organization_model.Organization, []error) {
	return organization_model.GetOrganizationsByTags(searchString)
}
