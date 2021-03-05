package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreatePublisher(ctx context.Context, input model.NewPublisher) (*model.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Publishers(ctx context.Context) ([]*model.Publisher, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var publishers []*model.Publisher

	db.Find(&publishers)

	return publishers, nil
}
