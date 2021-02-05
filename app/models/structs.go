package models

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type UID string

func (uid UID) String() string {
	return string(uid)
}

func (uid UID) isCorrect() bool {
	return true
}

type Tag struct {
	UID  UID    `json:"uid" db:"uid"`
	Name string `json:"name" db:"name"`
}

type SportCategory struct {
	UID  UID    `json:"uid"`
	Name string `json:"name"`
}
