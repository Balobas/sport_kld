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