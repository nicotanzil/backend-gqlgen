package model

import "time"

type BuyListing struct {
	BidID     int       `json:"bidId"`
	Bid       *Bid      `json:"bid" gorm:"foreignKey:BidID"`
	Buy       int       `json:"buy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
