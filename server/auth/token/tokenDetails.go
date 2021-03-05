package auth

type TokenDetails struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUid string `json:"accessUid"`
	RefreshUid string `json:"refreshUid"`
	AtExpires int64 `json:"atExpires"`
	RtExpires int64 `json:"rtExpires"`
}
