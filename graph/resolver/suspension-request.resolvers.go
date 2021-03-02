package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateSuspensionRequest(ctx context.Context, request model.InputSuspensionRequest) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var suspensionRequest model.SuspensionRequest

	suspensionRequest = model.SuspensionRequest{
		Description: request.Description,
		UserID:      request.User.ID,
	}

	db.Create(&suspensionRequest)
	return true, nil
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

func (r *queryResolver) SuspensionRequestsByUserID(ctx context.Context, id int) ([]*model.SuspensionRequest, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var requests []*model.SuspensionRequest

	db.Preload(clause.Associations).Where("user_id = ?", id).Find(&requests)

	return requests, nil
}
