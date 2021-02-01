package model

import "time"

type UserComment struct {
	ID          	int       	`json:"id"`
	UserID      	int       	`json:"userId"`
	User        	*User     	`json:"user" gorm:"foreignKey:UserID"`
	CommentatorID 	int     	`json:"commentatorId"`
	Commentator 	*User     	`json:"commentator" gorm:"foreignKey:CommentatorID"`
	Content     	string    	`json:"content"`
	CreatedAt   	time.Time 	`json:"createdAt"`
	UpdatedAt   	time.Time 	`json:"updatedAt"`
	DeletedAt   	*time.Time 	`json:"deletedAt"`
}
