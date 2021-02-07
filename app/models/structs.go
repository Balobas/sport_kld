package models

type UID string

func (uid UID) String() string {
	return string(uid)
}

type Tag struct {
	UID  UID    `json:"uid" db:"uid"`
	Name string `json:"name" db:"name"`
}
