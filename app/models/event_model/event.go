package event_model

import "sport_kld/app/models"

type Event struct {
	UID           models.UID `json:"uid" db:"uid"`
	Name          string     `json:"name" db:"name"`
	Description   string     `json:"description" db:"description"`
	Dates         string     `json:"dates" db:"dates"`
	Time          string     `json:"time" db:"time"`
	VisitorsNum   int64      `json:"visitorsNum" db:"visitorsNum"`
	VisitorsLimit int64      `json:"visitorsLimit" db:"visitorsLimit"`
	PlaceUID      models.UID `json:"placeUid"`
}
