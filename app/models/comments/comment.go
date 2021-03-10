package comments

import "sport_kld/app/models"

type Comment struct {
	UID       models.UID `json:"uid" db:"uid"`
	Text      string     `json:"text" db:"text"`
	AuthorUid models.UID `json:"authorUid" db:"authorUid"`
	Time      string     `json:"time" db:"time"`
}
