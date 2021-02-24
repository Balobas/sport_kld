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
	PlaceUID      models.UID `json:"placeUid" db:"placeUid"`
	CreatorUID    models.UID `json:"creatorUid" db:"creatorUid"`
	IsPrivate	  bool       `json:"isPrivate" db:"isPrivate"`
	EventPassword string     `json:"eventPassword" db:"eventPassword"`
	IsOver        bool       `json:"isOver" db:"isOver"`
}

type EventUserRole struct {
	UserUID         models.UID `json:"userUid" db:"userUid"`
	EventUID        models.UID `json:"eventUid" db:"eventUid"`
	Role            string `json:"role" db:"role"`
	RoleDescription string `json:"roleDescription" db:"roleDescription"`
}

type EventInfoPost struct {
	UID       models.UID `json:"uid" db:"uid"`
	EventUID  models.UID `json:"eventUid" db:"eventUid"`
	AuthorUID models.UID `json:"authorUid" db:"authorUid"`
	Text      string     `json:"text" db:"text"`
	CreatedTime string	 `json:"createdTime" db:"createdTime"`
	UpdatedTime string	 `json:"updatedTime" db:"updatedTime"`
}
