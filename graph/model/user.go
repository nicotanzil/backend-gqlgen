package model

import (
	"time"
)

type User struct {
	ID             int   		`json:"id"`
	AccountName    string   	`json:"accountName"`
	ProfileName    string   	`json:"profileName"`
	RealName       string   	`json:"realName"`
	Email          string   	`json:"email"`
	Password       string   	`json:"password"`
	Balance        float64  	`json:"balance"`
	CustomURL      string   	`json:"customURL"`
	ProfilePicture string    	`json:"profilePicture"`
	CountryID      int 			`json:"country"`
	CreatedAt 	   time.Time 	`json:"createdAt"`
	UpdatedAt 	   time.Time 	`json:"updatedAt"`
	DeletedAt 	   *time.Time 	`json:"deletedAt"`
}
