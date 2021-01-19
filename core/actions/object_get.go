package actions

import (
	"../../core"
	"../../data"
	"../../database"
	"fmt"
	"github.com/pkg/errors"
)

func ActionGet(uid data.UID, obj core.Object) error {
	objMap, err := database.DATABASE.Get(uid.String())
	if err != nil {
		return errors.WithStack(err)
	}
	objI, err := (obj).UnmarshalFromMap(objMap)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Println(objI)
	obj = objI
	return nil
}