package token

type TokenDetails struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUid string `json:"accessUid"`
	RefreshUid string `json:"refreshUid"`
	AtExpires int64 `json:"atExpires"`
	RtExpires int64 `json:"rtExpires"`
}

type Access struct {
	AccessUid string `json:"accessUid"`
	UserUid string `json:"userUid"`
	Expires int64 `json:"expires"`
}

type Refresh struct {
	RefreshUid string `json:"refreshUid"`
	UserUid string `json:"userUid"`
	Expires int64 `json:"expires"`
}
