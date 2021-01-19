package actions

import (
	"../../core"
	"../../database"
	"github.com/pkg/errors"
)

func ActionCreate(obj core.Object) (string, error) {
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