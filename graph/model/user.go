package model

import "time"

type User struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	AccountName string  `json:"accountName"`
	ProfileName string  `json:"profileName"`
	RealName    string  `json:"realName"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Balance     float64 `json:"balance"`
	CustomURL   string  `json:"customURL"`
	Summary     string  `json:"summary"`

	Avatar                string `json:"avatar"`
	AvatarFrame           string `json:"avatarFrames"`
	ProfileBackground     string `json:"profileBackground"`
	MiniProfileBackground string `json:"miniProfileBackground"`
	Theme                 string `json:"theme"`

	CountryID  int      `json:"countryId"`
	Country    *Country `json:"country" gorm:"foreignKey:CountryID"`
	Games      []*Game  `json:"games" gorm:"many2many:game_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Friends    []*User  `json:"friends" gorm:"many2many:user_friends;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Experience int      `json:"experience"`
	IsSuspend  bool     `json:"isSuspend"`

	FeaturedBadgeID int      `json:"featuredBadgeId"`
	FeaturedBadge   *Badge   `json:"featuredBadge" gorm:"foreignKey:FeaturedBadgeID"`
	Badges          []*Badge `json:"badges" gorm:"many2many:user_badges;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type NewUser struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type UpdateUser struct {
	AccountName           string `json:"accountName"`
	ProfileName           string `json:"profileName"`
	RealName              string `json:"realName"`
	CustomURL             string `json:"customURL"`
	Summary               string `json:"summary"`
	Avatar                string `json:"avatar"`
	AvatarFrame           string `json:"avatarFrames"`
	ProfileBackground     string `json:"profileBackground"`
	MiniProfileBackground string `json:"miniProfileBackground"`
	Theme                 string `json:"theme"`
	CountryID             int    `json:"CountryId"`
}

type InputUser struct {
	ID int `json:"id"`
}
