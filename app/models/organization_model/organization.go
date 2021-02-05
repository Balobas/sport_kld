package organization_model

import (
	"../../models"
)

type Organization struct {
	UID         models.UID    `json:"uid" db:"uid"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
