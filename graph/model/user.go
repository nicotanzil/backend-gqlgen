package model

import "time"

type User struct {
	ID                    int       `json:"id" gorm:"primaryKey"`
	AccountName           string    `json:"accountName"`
	ProfileName           string    `json:"profileName"`
	RealName              string    `json:"realName"`
	Email                 string    `json:"email"`
	Password              string    `json:"password"`
	Balance               float64   `json:"balance"`
	CustomURL             string    `json:"customURL"`
	Summary               string    `json:"summary"`
	ProfilePicture        string    `json:"profilePicture"`
	Avatar                string    `json:"avatar"`
	AvatarFrames          string    `json:"avatarFrames"`
	ProfileBackground     string    `json:"profileBackground"`
	MiniProfileBackground string    `json:"miniProfileBackground"`
	Theme                 string    `json:"theme"`
	CountryID      		  int 		`json:"countryId"`
	Country               *Country  `json:"country" gorm:"foreignKey:CountryID"`
	Games                 []*Game   `json:"games" gorm:"many2many:game_users;"`
	Friends               []*User   `json:"friends" gorm:"many2many:user_friends"`
	Experience            int       `json:"experience"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	DeletedAt             *time.Time `json:"deletedAt"`
}