package model

import "time"

type CommunityArtPost struct {
	ID          int                       `json:"id" gorm:"primaryKey"`
	Description string                    `json:"description"`
	Link        string                    `json:"link"`
	IsImage     bool                      `json:"isImage"`
	UserID      int                       `json:"userId;"`
	User        *User                     `json:"user" gorm:"foreignKey:UserID;"`
	Like        int                       `json:"like"`
	Reviews     []*CommunityArtPostReview `json:"reviews" gorm:"foreignKey:PostID;"`
	CreatedAt   time.Time                 `json:"createdAt"`
	UpdatedAt   time.Time                 `json:"updatedAt"`
	DeletedAt   *time.Time                `json:"deletedAt"`
}
