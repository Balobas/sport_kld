package core

import (
	"github.com/pkg/errors"
	"../database"
)

func ActionCreate(obj Object) (string, error) {
	if err := obj.Validate(); err != nil {
		return "", errors.WithStack(err)
	}
	objI, err := obj.GenerateAndSetKey()
	if err != nil {
		return "", errors.WithStack(err)
	}
	if err := database.DATABASE.Set(objI.GetKey(), objI); err != nil {
		return "", errors.WithStack(err)
	}
	return objI.GetKey(), nil
}