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
	Points      int     `json:"points"`
	CustomURL   string  `json:"customURL"`
	Summary     string  `json:"summary"`

	Avatar                  string                   `json:"avatar"`
	AvatarFrameID           int                      `json:"avatarFrameId"`
	AvatarFrame             *AvatarFrame             `json:"avatarFrame" gorm:"foreignKey:AvatarFrameID"`
	ProfileBackgroundID     int                      `json:"profileBackgroundId"`
	ProfileBackground       *ProfileBackground       `json:"profileBackground" gorm:"foreignKey:ProfileBackgroundID"`
	MiniProfileBackgroundID int                      `json:"miniProfileBackgroundId"`
	MiniProfileBackground   *MiniProfileBackground   `json:"miniProfileBackground" gorm:"foreignKey:MiniProfileBackgroundID"`
	ThemeID                 int                      `json:"themeId"`
	Theme                   *Theme                   `json:"theme" gorm:"foreignKey:ThemeID"`
	AvatarFrames            []*AvatarFrame           `json:"avatarFrames" gorm:"many2many:user_avatarFrames;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProfileBackgrounds      []*ProfileBackground     `json:"profileBackgrounds" gorm:"many2many:user_profileBackgrounds;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MiniProfileBackgrounds  []*MiniProfileBackground `json:"miniProfileBackgrounds" gorm:"many2many:user_miniProfileBackgrounds;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CountryID  int      `json:"countryId"`
	Country    *Country `json:"country" gorm:"foreignKey:CountryID"`
	Games      []*Game  `json:"games" gorm:"many2many:game_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Friends    []*User  `json:"friends" gorm:"many2many:user_friends;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Items      []*Item  `json:"items" gorm:"many2many:user_items;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Experience int      `json:"experience"`
	IsSuspend  bool     `json:"isSuspend"`

	FeaturedBadgeID int      `json:"featuredBadgeId"`
	FeaturedBadge   *Badge   `json:"featuredBadge" gorm:"foreignKey:FeaturedBadgeID"`
	Badges          []*Badge `json:"badges" gorm:"many2many:user_badges;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UpdateUser struct {
	AccountName             string `json:"accountName"`
	ProfileName             string `json:"profileName"`
	RealName                string `json:"realName"`
	CustomURL               string `json:"customURL"`
	Summary                 string `json:"summary"`
	Avatar                  string `json:"avatar"`
	AvatarFrameID           int    `json:"avatarFrameId"`
	ProfileBackgroundID     int    `json:"profileBackgroundId"`
	MiniProfileBackgroundID int    `json:"miniProfileBackgroundId"`
	ThemeID                 int    `json:"themeId"`
	FeaturedBadgeID         int    `json:"featuredBadgeId"`
	CountryID               int    `json:"CountryId"`
}

type NewUser struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

type InputUser struct {
	ID int `json:"id"`
}
