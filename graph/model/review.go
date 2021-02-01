package model

import "time"

type Review struct {
	ID        int    	`json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	Upvote    string    `json:"upvote"`
	Downvote  string    `json:"downvote"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}