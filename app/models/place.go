package models

import (
	"../../database"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Place struct {
	UID          UID    `json:"uid" bd:"uid"`
	Name         string `json:"name" db:"name"`
	BuildingName string `json:"buildingName" db:"buildingName"`
	BuildingType string `json:"buildingType" db:"buildingType"`
	Description  string `json:"description" db:"description"`
	Address      string `json:"address" db:"address"`
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
func (place *Place) GenerateAndSetKey() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}
	place.UID = UID(uid.String())
	return nil
}
//

//validations
func (place *Place) Validate() error {

	return nil
}

func (place *Place) IsValidUID() bool {
	return len(place.UID) != 0
}

//

func (place *Place) GetByUID(uid UID) error {
	if err := database.MysqlDB.Get(place, "select * from places where uid=?", uid.String()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
