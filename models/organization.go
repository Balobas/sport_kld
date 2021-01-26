package models

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Organization struct {
	UID         UID    `json:"uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PlacesUIDs  []UID  `json:"placesUIDs"`
	TagsUIDs    []UID  `json:"tagsUIDs"`
}

//keys
func (org *Organization) GenerateAndSetKey() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}
	org.UID = UID(uid.String())
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
