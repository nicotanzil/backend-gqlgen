package model

import (
	"time"
)

type Developer struct {
	ID   		int 		`json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}

type NewDeveloper struct {
	Name string `json:"name"`
}

