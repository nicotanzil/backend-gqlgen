package model

import "time"

type GameDiscussionReply struct {
	ID           int             `json:"id" gorm:"primaryKey"`
	Description  string          `json:"description"`
	DiscussionID int             `json:"discussionId"`
	Discussion   *GameDiscussion `json:"discussion" gorm:"foreignKey:DiscussionID"`
	UserID       int             `json:"userId"`
	User         *User           `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	DeletedAt    *time.Time       `json:"deletedAt"`
}
