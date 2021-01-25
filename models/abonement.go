package models

import (
	"../app/utils"
	"../core"
	"encoding/json"
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
func (abonement Abonement) GenerateAndSetKey() (core.Object, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	abonement.UID = UID(uid.String())
	return abonement, nil
}

func (abonement Abonement) GetKey() string {
	return abonement.UID.String()
}

func (abonement Abonement) SetKey(key string) (core.Object, error) {
	if !(UID(key)).isCorrect() {
		return nil, errors.Errorf("uid is not correct %s", key)
	}
	abonement.UID = UID(key)
	return abonement, nil
}

//

//new
func (abonement Abonement) New() core.Object {
	return Abonement{
		UID:            "",
		Name:           "",
		Description:    "",
		Price:          "",
		ActivateTime:   "",
		PlacesUIDs:     nil,
		CategoriesUIDs: nil,
	}
}

//

//unmarshal
func (abonement Abonement) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&abonement, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return abonement, nil
}

//

//update
func (abonement Abonement) Update(updateAbonementInterface interface{}) (core.Object, error) {
	var updateAbonement Abonement
	b, err := json.Marshal(updateAbonementInterface)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := json.Unmarshal(b, &updateAbonement); err != nil {
		return nil, errors.WithStack(err)
	}

	//обновление полей, которые можно обновить, ошибки если нельзя обновить
	abonement = updateAbonement

	return abonement, nil
}

//

//validations
func (abonement Abonement) Validate() error {

	return nil
}

func (abonement Abonement) IsValidUID() bool {
	return len(abonement.UID) != 0
}

//
