package organization

import (
	"../../../data"
	"../../../database"
	"../../utils"
	"github.com/pkg/errors"
)

func GetOrganization(uid data.UID) (data.Organization, error) {
	var organization data.Organization
	orgMap, err := database.DATABASE.Get(string(uid))
	if err != nil {
		return data.Organization{}, errors.WithStack(err)
	}
	if err := utils.UnmarshalFromMap(&organization, orgMap); err != nil {
		return data.Organization{}, errors.WithStack(err)
	}
	return organization, nil
}