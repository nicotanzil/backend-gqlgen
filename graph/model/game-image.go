package model

import "time"

type GameImage struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	GameID    int       `json:"gameId"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type InputGameImage struct {
	GameID int    `json:"gameId"`
	Link   string `json:"link"`
}
