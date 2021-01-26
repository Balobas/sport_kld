package models

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Abonement struct {
	UID          UID    `json:"uid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	ActivateTime string `json:"activateTime"`
	//Места, где действует абонемент
	PlacesUIDs []UID `json:"placesUIDs"`
	//Виды спорта относящиеся к абонементу
	CategoriesUIDs []UID `json:"categoriesUIDs"`
}

//keys
func (abonement Abonement) GenerateAndSetKey() error {
	uid, err := uuid.NewV4()
	if err != nil {
		return errors.WithStack(err)
	}
	abonement.UID = UID(uid.String())
	return nil
}


//validations
func (abonement Abonement) Validate() error {

	return nil
}

func (abonement Abonement) IsValidUID() bool {
	return len(abonement.UID) != 0
}

//
