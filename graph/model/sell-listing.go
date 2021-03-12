package model

import "time"

type SellListing struct {
	ItemID    int        `json:"itemId"`
	Item      *Item      `json:"item" gorm:"foreignKey:ItemID"`
	Sell      int        `json:"sell"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
