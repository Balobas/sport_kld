package organization_model

import (
	"../../models"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Organization struct {
	UID         models.UID    `json:"uid" db:"uid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

//keys
func (org *Organization) GenerateAndSetKey() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}
	org.UID = models.UID(uid.String())
	return nil
}
//

//validations
func (org Organization) Validate() error {

	return nil
}

func (org Organization) IsValidUID() bool {
	return len(org.UID) != 0
}

//
