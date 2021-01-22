package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nicotanzil/backend-gqlgen/database"

	"github.com/nicotanzil/backend-gqlgen/graph/generated"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var users []*model.User

	//db.Find(&users)	//populate the users variable with `SELECT * FROM users`

	db.Preload("Country").Find(&users) //populate the users with each corresponding countries

	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
