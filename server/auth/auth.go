package auth

type Auth struct {
	Uid string `json:"uid" db:"uid"`
	UserUid string `json:"userUid" db:"userUid"`
	AccessUid string `json:"accessUid" db:"accessUid"`
	RefreshUid string `json:"refreshUid" db:"refreshUid"`
}
