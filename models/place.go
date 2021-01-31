package models

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"../database"
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
