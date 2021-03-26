package comment_model

import "sport_kld/app/models"

type Comment struct {
	UID       models.UID `json:"uid" db:"uid"`
	Text      string     `json:"text" db:"text"`
	AuthorUid models.UID `json:"authorUid" db:"authorUid"`
	CreatedTime      int64     `json:"time" db:"createdTime"`
	UpdatedTime int64 `json:"updatedTime" db:"updatedTime"`
	IsDeleted bool `json:"isDeleted" db:"isDeleted"`
}

func (c *Comment) CheckAccess(executorUid models.UID) bool {
	return c.AuthorUid == executorUid
}

type PlaceComment struct {
	Comment
	PlaceUID models.UID `json:"placeUid" db:"placeUid"`
}

type OrganizationComment struct {
	Comment
	OrganizationUID models.UID `json:"organizationUid" db:"organizationUid"`
}

