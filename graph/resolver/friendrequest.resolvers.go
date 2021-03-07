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
	dbClose, _ := db.DB()
	defer dbClose.Close()

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
	dbClose, _ := db.DB()
	defer dbClose.Close()

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
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var requests []*model.FriendRequest

	db.Preload(clause.Associations).Find(&requests)

	return requests, nil
}

func (r *queryResolver) GetFriendRequestByRequestedID(ctx context.Context, id int) ([]*model.FriendRequest, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var requests []*model.FriendRequest

	db.Preload(clause.Associations).Where("requested_id = ? AND status LIKE ?", id, "Pending").Find(&requests)

	return requests, nil
}

func (r *queryResolver) ValidateFriendRequestExists(ctx context.Context, requesterID int, requestedID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var request model.FriendRequest

	if requestedID != requesterID {
		db.Where("requested_id = ? AND requester_id = ? AND status LIKE ? OR status LIKE ?", requestedID, requesterID, "Pending", "Ignored").First(&request)
	} else {
		request.ID = 0
	}

	if request.ID != 0 {
		return true, nil // Stranger Pending / Ignored
	}

	return false, nil // Stranger Declined / No Interaction
}

func (r *queryResolver) GetPendingFriendRequestCount(ctx context.Context, id int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var requestCount int64

	db.Model(model.FriendRequest{}).Where("requested_id = ? AND status = ?", id, "Pending").Count(&requestCount)

	return int(requestCount), nil
}
