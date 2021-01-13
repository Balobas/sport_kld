package data

import (
	"../core"
	"encoding/json"
	"github.com/pkg/errors"
)

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