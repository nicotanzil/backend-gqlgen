package model

import "time"

type Gift struct {
	ID         int        `json:"id"`
	SenderID   int        `json:"senderId"`
	Sender     *User      `json:"sender" gorm:"foreignKey:SenderID"`
	ReceiverID int        `json:"receiverId"`
	Receiver   *User      `json:"receiver" gorm:"foreignKey:ReceiverID"`
	FirstName  string     `json:"firstName"`
	Message    string     `json:"message"`
	Sentiment  string     `json:"sentiment"`
	Signature  string     `json:"signature"`
	IsComplete bool       `json:"isComplete"`
	IsOpen     bool       `json:"isOpen"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

type NewGift struct {
	SenderID   int    `json:"senderId"`
	ReceiverID int    `json:"receiverId"`
	FirstName  string `json:"firstName"`
	Message    string `json:"message"`
	Sentiment  string `json:"sentiment"`
	Signature  string `json:"signature"`
}
