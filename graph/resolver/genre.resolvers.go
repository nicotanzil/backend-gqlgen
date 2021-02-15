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

func (r *mutationResolver) CreateGenre(ctx context.Context, input model.NewGenre) (*model.Genre, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Genres(ctx context.Context) ([]*model.Genre, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var genres []*model.Genre

	db.Preload(clause.Associations).Find(&genres)

	return genres, nil
}
