package model

import "time"

type ItemTransaction struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	ItemID    int        `json:"itemId"`
	Item      *Item      `json:"item" gorm:"foreignKey:ItemID"`
	SellerID  int        `json:"sellerId"`
	Seller    *User      `json:"seller" gorm:"foreignKey:SellerID"`
	BuyerID   int        `json:"buyerId"`
	Buyer     *User      `json:"buyer" gor:"foreignKey:BuyerID"`
	Price     int        `json:"price"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
