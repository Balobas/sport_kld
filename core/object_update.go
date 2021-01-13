package core

import (
	"../database"
	"github.com/pkg/errors"
)

func ActionUpdate(obj Object) (string, error) {
	if !obj.IsValidUID() {
		return "", errors.Errorf("uid '%s' is not valid", obj.GetKey())
	}
	if err := obj.Validate(); err != nil {
		return "", errors.WithStack(err)
	}
	oldObj, err := obj.SetKey(obj.GetKey())
	if err := ActionGet(obj.GetKey(), &oldObj); err != nil {
		return "", errors.WithStack(err)
	}
	oldObjI, err := oldObj.Update(obj)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if err := database.DATABASE.Set(obj.GetKey(), oldObjI); err != nil {
		return "", errors.WithStack(err)
	}
	return obj.GetKey(), nil
}