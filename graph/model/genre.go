package model

import (
	"time"
)

type Genre struct {
	ID          int      	`json:"id" gorm:"primaryKey"`
	Name        string    	`json:"name"`
	Description string    	`json:"description"`
	Games       []*Game   	`json:"games" gorm:"many2many:game_genres;"`
	CreatedAt   time.Time 	`json:"createdAt"`
	UpdatedAt   time.Time 	`json:"updatedAt"`
	DeletedAt   *time.Time 	`json:"deletedAt"`
}

type NewGenre struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InputGenre struct {
	ID int `json:"id"`
}

