package model

import (
	"time"
)

type Publisher struct {
	ID   		int `json:"id" gorm:"primaryKey"`
	Name 		string `json:"name"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	DeletedAt 	*time.Time `json:"deletedAt"`
}

type NewPublisher struct {
	Name string `json:"name"`
}