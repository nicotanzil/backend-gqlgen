package model

import (
	"time"
)

type Tag struct {
	ID   		int 		`json:"id" gorm:"primaryKey"`
	Name 		string 		`json:"name"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}
