package model

import "time"

type AnimatedAvatar struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
