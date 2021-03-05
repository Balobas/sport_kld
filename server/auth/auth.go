package auth

type Auth struct {
	Uid string `json:"uid"`
	UserUid string `json:"userUid"`
	AccessUid string `json:"accessUid"`
	RefreshUid string `json:"refreshUid"`
}
