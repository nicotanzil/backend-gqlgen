package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateBid(ctx context.Context, typeID int, userID int) (*model.Bid, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var bid model.Bid

	bid = model.Bid{
		ItemType: &model.ItemType{ID: typeID},
		User:     &model.User{ID: userID},
	}

	db.Create(&bid)

	return &bid, nil
}
