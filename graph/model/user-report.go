package model

import "time"

type UserReport struct {
	ID          int       `json:"id"`
	ReportedId 	int		  `json:"ReportedId"`
	Reported    *User     `json:"Reported" gorm:"foreignKey:ReportedId;"`
	ReporterId 	int		  `json:"ReporterId"`
	Reporter    *User     `json:"Reporter" gorm:"foreignKey:ReporterId;"`
	Description string   `json:"Description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}