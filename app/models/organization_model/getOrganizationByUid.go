package organization_model

import (
	"github.com/pkg/errors"
	"sport_kld/app/models"
	"sport_kld/database"
)

func GetOrganizationByUID(uid models.UID) (Organization, error) {
	var org Organization
	if err := database.MysqlDB.Get(&org, "select * from organizations where uid=?", uid.String()); err != nil {
		return Organization{}, errors.WithStack(err)
	}
	org.Preprocess()
	return org, nil
}

func GetOrganizationsByUIDs(uids []models.UID) (result []Organization, resultErrors []error) {
	for _, uid := range uids {
		if org, err := GetOrganizationByUID(uid); err != nil {
			resultErrors = append(resultErrors, err)
		} else {
			result = append(result, org)
		}
	}

	return
}