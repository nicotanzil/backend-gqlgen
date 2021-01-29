package model

type Login struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type AccessDetails struct {
	AccessUUID string `json:"accessUuid"`
	UserID     int `json:"userId"`
}

type TokenDetail struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	AccessUUID     string `json:"accessUuid"`
	RefreshUUID    string `json:"refreshUuid"`
	AccessExpires  int64    `json:"accessExpires"`
	RefreshExpires int64    `json:"refreshExpires"`
}