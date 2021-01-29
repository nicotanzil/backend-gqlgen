package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateDeveloper(ctx context.Context, input model.NewDeveloper) (*model.Developer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Developers(ctx context.Context) ([]*model.Developer, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var developers []*model.Developer

	db.Find(&developers)

	return developers, nil
}
