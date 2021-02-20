package model

import "time"

type NewPromo struct {
	DiscountPercentage int `json:"discountPercentage"`
	ValidUntil         int `json:"validUntil"`
}

type Promo struct {
	ID                 int       `json:"id"`
	DiscountPercentage int       `json:"discountPercentage"`
	ValidUntil         int       `json:"validUntil"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	DeletedAt          *time.Time `json:"deletedAt"`
}