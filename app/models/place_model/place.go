package place_model

import (
	"../../../database"
	"../../models"
	"github.com/pkg/errors"
)

type Place struct {
	UID          models.UID    `json:"uid" bd:"uid"`
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
}

func (place *Place) GetByUID(uid models.UID) error {
	if err := database.MysqlDB.Get(place, "select * from places where uid=?", uid.String()); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
