package models

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Abonement struct {
	UID          UID    `json:"uid" db:"uid"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Price        string `json:"price" db:"price"`
	ActivateTime string `json:"activateTime" db:"activateTime"`
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
