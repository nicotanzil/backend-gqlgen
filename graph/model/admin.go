package model

import "time"

type Admin struct {
	ID          int       `json:"id"`
	AccountName string    `json:"accountName"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}