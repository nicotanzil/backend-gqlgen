package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertCommentByReviewID(ctx context.Context, reviewID int, userID int, comment string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var reply model.GameDiscussionReply

	reply = model.GameDiscussionReply{
		Description:  comment,
		Discussion:   &model.GameDiscussion{ID: reviewID},
		User:         &model.User{ID: userID},
	}

	db.Create(&reply)
	return true, nil
}

func (r *queryResolver) GameDiscussionReplies(ctx context.Context) ([]*model.GameDiscussionReply, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var replies []*model.GameDiscussionReply

	db.Preload(clause.Associations).
		Find(&replies)

	return replies, nil
}
