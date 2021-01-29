package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var reviews []*model.Review

	db.Find(&reviews)

	return reviews, nil
}
