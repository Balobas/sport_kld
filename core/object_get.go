package core

import (
	"../database"
	"github.com/pkg/errors"
)

func ActionGet(uid string, obj *Object) error {
	objMap, err := database.DATABASE.Get(uid)
	if err != nil {
		return errors.WithStack(err)
	}
	objI, err := (*obj).UnmarshalFromMap(objMap)
	if err != nil {
		return errors.WithStack(err)
	}
	*obj = objI
	return nil
}