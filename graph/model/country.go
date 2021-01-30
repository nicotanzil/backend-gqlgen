package model

import (
	"time"
)

type Country struct {
	ID   int 				`json:"id"`
	Name string 			`json:"name"`
	CreatedAt time.Time 	`json:"createdAt"`
	UpdatedAt time.Time 	`json:"updatedAt"`
	DeletedAt *time.Time 	`json:"deletedAt"`
}
