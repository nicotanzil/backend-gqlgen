package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *queryResolver) GetVotesByReviewID(ctx context.Context, input *string) ([]*model.ReviewVote, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var reviewVotes []*model.ReviewVote

	db.Where("id = ?", input).Find(&reviewVotes)

	return reviewVotes, nil
}
