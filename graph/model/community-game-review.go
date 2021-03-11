package model

import "time"

type CommunityGameReview struct {
	ID            int                           `json:"id" gorm:"primaryKey"`
	Description   string                        `json:"description"`
	UserID        int                           `json:"userId"`
	User          *User                         `json:"user" gorm:"foreignKey:UserID"`
	GameID        int                           `json:"gameId"`
	Game          *Game                         `json:"game" gorm:"foreignKey:GameID"`
	IsRecommended bool                          `json:"isRecommended"`
	Comments      []*CommunityGameReviewComment `json:"comments" gorm:"foreignKey:ReviewID"`
	HelpfulCount  int                           `json:"helpfulCount"`
	CreatedAt     time.Time                     `json:"createdAt"`
	UpdatedAt     time.Time                     `json:"updatedAt"`
	DeletedAt     *time.Time                    `json:"deletedAt"`
}

type CommunityGameReviewInput struct {
	Description   string `json:"description"`
	UserID        int    `json:"userId"`
	GameID        int    `json:"gameId"`
	IsRecommended bool   `json:"isRecommended"`
}
