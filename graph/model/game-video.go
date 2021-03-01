package model

import "time"

type GameVideo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	GameID    int       `json:"gameId"`
	Link      string    `json:"link"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type InputGameVideo struct {
	GameID int    `json:"gameId"`
	Link   string `json:"link"`
}