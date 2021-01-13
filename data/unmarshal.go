package data

import (
	"../app/utils"
	"../core"
	"github.com/pkg/errors"
)

func (org Organization) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&org, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return org, nil
}

func (place Place) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&place, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return place, nil
}

func (abonement Abonement) UnmarshalFromMap(objMap map[string]interface{}) (core.Object, error) {
	if err := utils.UnmarshalFromMap(&abonement, objMap); err != nil {
		return nil, errors.WithStack(err)
	}
	return abonement, nil
}
