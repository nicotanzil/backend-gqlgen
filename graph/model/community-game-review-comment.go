package model

import "time"

type CommunityGameReviewComment struct {
	ID          int                  `json:"id" gorm:"primaryKey"`
	UserID      int                  `json:"userId"`
	User        *User                `json:"user" gorm:"foreignKey:UserID"`
	ReviewID    int                  `json:"reviewId" gorm:"foreignKey:ReviewID"`
	Review      *CommunityGameReview `json:"review"`
	Description string               `json:"description"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	DeletedAt   *time.Time           `json:"deletedAt"`
}
