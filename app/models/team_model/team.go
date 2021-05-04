package team_model

import "sport_kld/app/models"

type Team struct {
	UID         models.UID `json:"uid" db:"uid"`
	Name        string     `json:"name" db:"name"`
	CreatedDate int64      `json:"createdDate" db:"createdDate"`
	CreatorUid 	models.UID `json:"creatorUid" db:"creatorUid"`
}

type TeamMember struct {
	UserUid     models.UID `json:"userUid" db:"userUid"`
	TeamUid     models.UID `json:"teamUid" db:"teamUid"`
	ConnectDate int64      `json:"connectDate"`
}
