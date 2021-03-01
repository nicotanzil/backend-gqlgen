package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateSuspensionRequest(ctx context.Context, request *model.InputSuspensionRequest) (*model.SuspensionRequest, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SuspensionRequests(ctx context.Context) ([]*model.SuspensionRequest, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var requests []*model.SuspensionRequest

	db.Preload(clause.Associations).Find(&requests)

	return requests, nil
}
