package model

import "time"

type GameDiscussion struct {
	ID          int                    `json:"id" gorm:"primaryKey"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	GameID      int                    `json:"gameId"`
	Game        *Game                  `json:"game" gorm:"foreignKey:GameID"`
	UserID      int                    `json:"userId"`
	User        *User                  `json:"user" gorm:"foreignKey:UserID"`
	Replies     []*GameDiscussionReply `json:"replies" gorm:"foreignKey:DiscussionID"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	DeletedAt   *time.Time             `json:"deletedAt"`
}
