package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/app/providers"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*model.Game

	db.Preload(clause.Associations).Find(&games)

	return games, nil
}

func (r *queryResolver) GetGamePaginationAdmin(ctx context.Context, page *int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*model.Game

	left := (providers.ADMIN_GAME_PAGINATION * (*page-1)) + 1
	right := providers.ADMIN_GAME_PAGINATION * (*page)

	db.Preload(clause.Associations).Where("id >= ? AND id <= ?", left, right).Debug().Find(&games)

	return games, nil
}
