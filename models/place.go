package models

import (
	"../app/utils"
	"../core"
	"encoding/json"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Place struct {
	UID          UID    `json:"uid"`
	Name         string `json:"name"`
	BuildingName string `json:"buildingName"`
	BuildingType string `json:"buildingType"`
	Description  string `json:"description"`
	Address      string `json:"adress"`
	City         string `json:"city"`
	OpeningHours string `json:"openingHours"`
	PostIndex    string `json:"postIndex"`
	WebSite      string `json:"webSite"`
	Phones       string `json:"phones"`
	Email        string `json:"email"`
	Facebook     string `json:"facebook"`
	Instagram    string `json:"instagram"`
	Twitter      string `json:"twitter"`
	VK           string `json:"vk"`

	TagsUIDs []UID `json:"tagsUIDs"`
	//под вопросом//
	HolderOrganization Organization   `json:"organization"`
	BasedOrganizations []Organization `json:"basedOrganizations"`
	FreeVisit          bool           `json:"freeVisit"`
	//
}

//keys
func (place Place) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	place.UID = UID(uid.String())
	return place, nil
}

func (place Place) GetKey() string {
	return place.UID.String()
}

func (place Place) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	place.UID = UID(key)
	return place, nil
}

//

//new
func (place Place) New() core.Object {
	return Place{
		UID:                "",
		Name:               "",
		BuildingName:       "",
		BuildingType:       "",
		Description:        "",
		Address:            "",
		City:               "",
		OpeningHours:       "",
		PostIndex:          "",
		WebSite:            "",
		Phones:             "",
		Email:              "",
		Facebook:           "",
		Instagram:          "",
		Twitter:            "",
		VK:                 "",
		TagsUIDs:           nil,
		HolderOrganization: Organization{},
		BasedOrganizations: nil,
		FreeVisit:          false,
	}
}

//

//unmarshal
func (place Place) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&place, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return place, nil
}

//

//update
func (place Place) Update(updatePlaceInterface interface{}) (core.Object, error) {
	var updatePlace Place
	b, err := json.Marshal(updatePlaceInterface)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := json.Unmarshal(b, &updatePlace); err != nil {
		return nil, errors.WithStack(err)
	}

	//обновление полей, которые можно обновить, ошибки если нельзя обновить
	place = updatePlace

	return place, nil
}

//

//validations
func (place Place) Validate() error {

	return nil
}

func (place Place) IsValidUID() bool {
	return len(place.UID) != 0
}

//
