package model

import "time"

type WalletCode struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Code      string    `json:"code"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}