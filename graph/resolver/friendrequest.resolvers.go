package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateFriendRequest(ctx context.Context, requesterID int, requestedID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	request := model.FriendRequest{
		Requester: &model.User{ID: requesterID},
		Requested: &model.User{ID: requestedID},
		Status:    "Pending",
	}

	db.Create(&request)

	return true, nil
}

func (r *mutationResolver) AcceptFriendRequest(ctx context.Context, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var request model.FriendRequest

	db.First(&request, id)

	request.Status = "Accepted"
	db.Save(&request)

	return true, nil
}

func (r *queryResolver) FriendRequests(ctx context.Context) ([]*model.FriendRequest, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var requests []*model.FriendRequest

	db.Preload(clause.Associations).Find(&requests)

	return requests, nil
}

func (r *queryResolver) GetFriendRequestByRequestedID(ctx context.Context, id int) ([]*model.FriendRequest, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var requests []*model.FriendRequest

	db.Preload(clause.Associations).Where("requested_id = ? AND status LIKE ?", id, "Pending").Find(&requests)

	return requests, nil
}

func (r *queryResolver) ValidateFriendRequestExists(ctx context.Context, requesterID int, requestedID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var request model.FriendRequest

	db.Where("requested_id = ? AND requester_id = ? AND status LIKE ?", requestedID, requesterID, "Pending").First(&request)

	if request.ID != 0 {
		return true, nil
	}

	return false, nil
}
