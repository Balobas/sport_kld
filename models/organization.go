package models

import (
	"../app/utils"
	"../core"
	"encoding/json"
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
func (org Organization) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	org.UID = UID(uid.String())
	return org, nil
}

func (org Organization) GetKey() string {
	return org.UID.String()
}

func (org Organization) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	org.UID = UID(key)
	return org, nil
}

//

//new
func (org Organization) New() core.Object {
	return Organization{
		UID:         "",
		Name:        "",
		Description: "",
	}
}

//

//unmarshal
func (org Organization) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&org, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return org, nil
}

//

//update
func (org Organization) Update(updateOrgInterface interface{}) (core.Object, error) {
	var updateOrg Organization
	b, err := json.Marshal(updateOrgInterface)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := json.Unmarshal(b, &updateOrg); err != nil {
		return nil, errors.WithStack(err)
	}

	//обновление полей, которые можно обновить, ошибки если нельзя обновить
	org = updateOrg

	return org, nil
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
