package model

import (
	"time"
)

type System struct {
	ID        int    	`json:"id" gorm:"primaryKey"`
	Os        string    `json:"os"`
	Memory    int       `json:"memory"`
	Graphics  string    `json:"graphics"`
	Storage   int       `json:"storage"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}


type NewSystem struct {
	Os       string `json:"os"`
	Memory   int    `json:"memory"`
	Graphics string `json:"graphics"`
	Storage  int    `json:"storage"`
}