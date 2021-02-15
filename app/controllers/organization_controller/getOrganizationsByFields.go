package organization_controller

import "sport_kld/app/models/organization_model"

func GetOrganizationsByFields(fieldsNames []string, searchString string) ([]organization_model.Organization, []error) {
	return organization_model.GetOrganizationsByFields(fieldsNames, searchString)
}
