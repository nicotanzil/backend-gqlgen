package model

import "time"

type SuspensionRequest struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	UserID		int			`json:"userId"`
	User        *User     `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

type InputSuspensionRequest struct {
	Description string     `json:"description"`
	User        *InputUser `json:"user"`
}