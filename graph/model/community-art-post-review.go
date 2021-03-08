package model

import "time"

type CommunityArtPostReview struct {
	ID        int               `json:"id" gorm:"primaryKey"`
	PostID    int               `json:"postId"`
	Post      *CommunityArtPost `json:"post" gorm:"foreignKey:PostID;"`
	UserID    int               `json:"userId"`
	User      *User             `json:"user" gorm:"foreignKey:UserID"`
	Comment   string            `json:"comment"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	DeletedAt *time.Time        `json:"deletedAt"`
}
