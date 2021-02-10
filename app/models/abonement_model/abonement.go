package abonement_model

import (
	"sport_kld/app/models"
)

type Abonement struct {
	UID          models.UID    `json:"uid" db:"uid"`
	Name         string `json:"name" db:"name"`
	Description  string `json:"description" db:"description"`
	Price        string `json:"price" db:"price"`
	ActivateTime string `json:"activateTime" db:"activateTime"`
}
