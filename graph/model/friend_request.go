package model

import "time"

type FriendRequest struct {
	ID        		int    	`json:"id" gorm:"primaryKey"`
	RequesterID 	int  	`json:"requesterId"`
	Requester 		*User  	`json:"requester" gorm:"foreignKey:RequesterID"`
	RequestedID 	int  	`json:"requestedId"`
	Requested 		*User  	`json:"requested" gorm:"foreignKey:RequestedID"`
	Status    		string 	`json:"status"` // Accepted, Pending, Ignored

	CreatedAt  time.Time  	`json:"createdAt"`
	UpdatedAt  time.Time  	`json:"updatedAt"`
	DeletedAt  *time.Time 	`json:"deletedAt"`
}
