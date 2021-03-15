package place_model

import (
	"sport_kld/app/models"
)

type Place struct {
	UID          models.UID `json:"uid" db:"uid"`
	Name         string     `json:"name" db:"name"`
	BuildingName string     `json:"buildingName" db:"buildingName"`
	BuildingType string     `json:"buildingType" db:"buildingType"`
	Description  string     `json:"description" db:"description"`
	Address      string     `json:"address" db:"address"`
	City         string     `json:"city" db:"city"`
	OpeningHours string     `json:"openingHours" db:"openingHours"`
	PostIndex    string     `json:"postIndex" db:"postIndex"`
	Lat          float64    `json:"lat"`
	Lon          float64  	`json:"lon"`
	WebSite      string     `json:"webSite" db:"webSite"`
	Phones       string     `json:"phones" db:"phones"`
	Email        string     `json:"email" db:"email"`
	Facebook     string     `json:"facebook" db:"facebook"`
	Instagram    string     `json:"instagram" db:"instagram"`
	Twitter      string     `json:"twitter" db:"twitter"`
	VK           string     `json:"vk" db:"vk"`
}

func (place *Place) Preprocess() {
	models.ReplaceNoneOrNanValueByEmptyString(&place.Name)
	models.ReplaceNoneOrNanValueByEmptyString(&place.BuildingName)
	models.ReplaceNoneOrNanValueByEmptyString(&place.BuildingType)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Description)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Address)
	models.ReplaceNoneOrNanValueByEmptyString(&place.City)
	models.ReplaceNoneOrNanValueByEmptyString(&place.OpeningHours)
	models.ReplaceNoneOrNanValueByEmptyString(&place.PostIndex)
	models.ReplaceNoneOrNanValueByEmptyString(&place.WebSite)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Phones)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Email)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Facebook)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Instagram)
	models.ReplaceNoneOrNanValueByEmptyString(&place.Twitter)
	models.ReplaceNoneOrNanValueByEmptyString(&place.VK)
}
