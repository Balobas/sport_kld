package models

import (
	"../app/utils"
	"../core"
	"encoding/json"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Place struct {
	UID          UID    `json:"uid" bd:"uid"`
	Name         string `json:"name" db:"name"`
	BuildingName string `json:"buildingName" db:"buildingName"`
	BuildingType string `json:"buildingType" db:"buildingType"`
	Description  string `json:"description" db:"description"`
	Address      string `json:"adress" db:"adress"`
	City         string `json:"city" db:"city"`
	OpeningHours string `json:"openingHours" db:"openingHours"`
	PostIndex    string `json:"postIndex" db:"postIndex"`
	WebSite      string `json:"webSite" db:"webSite"`
	Phones       string `json:"phones" db:"phones"`
	Email        string `json:"email" db:"email"`
	Facebook     string `json:"facebook" db:"facebook"`
	Instagram    string `json:"instagram" db:"instagram"`
	Twitter      string `json:"twitter" db:"twitter"`
	VK           string `json:"vk" db:"vk"`
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
		UID:          "",
		Name:         "",
		BuildingName: "",
		BuildingType: "",
		Description:  "",
		Address:      "",
		City:         "",
		OpeningHours: "",
		PostIndex:    "",
		WebSite:      "",
		Phones:       "",
		Email:        "",
		Facebook:     "",
		Instagram:    "",
		Twitter:      "",
		VK:           "",
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
