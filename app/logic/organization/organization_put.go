package organization

import (
	"../../../data"
	"../../../database"
	"crypto/md5"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

func PutOrganization(organization data.Organization) (data.UID, error) {

	//validation
	if err := organization.Validate(); err != nil {
		return "", errors.WithStack(err)
	}

	//uid generation...
	ts := time.Now().Unix()

	if len(organization.UID) != 0 {
		prevOrg, err := GetOrganization(organization.UID)
		if err != nil {
			return "", errors.WithStack(err)
		}

		//удалить это говно потом
		prevOrg.Name = prevOrg.Name


		//обновляем поля...


	} else {
		uidMd5 := md5.Sum([]byte(organization.Name + ":" + organization.MainAddress + ":" + strconv.Itoa(int(ts))))
		UUID, _ := uuid.FromBytes(uidMd5[:])
		uid := data.UID(UUID.String())

		organization.UID = uid
	}

	if err := database.DATABASE.Set(string(organization.UID), organization); err != nil {
		return "", err
	}
	return organization.UID, nil
}
