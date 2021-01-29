package model

import "time"

type ReviewVote struct {
	ID         	int 		`json:"id"`
	UserID     	int 		`json:"userId"`
	ReviewID   	int 		`json:"reviewId"`
	VoteStatus 	bool   		`json:"voteStatus"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}
