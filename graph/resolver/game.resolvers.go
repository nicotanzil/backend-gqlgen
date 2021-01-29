package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/generated"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *gameResolver) Publisher(ctx context.Context, obj *model.Game) (*model.Publisher, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var publisher model.Publisher

	db.Where("id = ?", obj.PublisherID).First(&publisher)

	return &publisher, nil
}

func (r *gameResolver) System(ctx context.Context, obj *model.Game) (*model.System, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var system model.System

	db.Where("id = ?", obj.SystemID).First(&system)

	return &system, nil
}

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*model.Game

	db.Preload("Genres").Preload("Publisher").Preload("System").Find(&games)

	return games, nil
}

// Game returns generated.GameResolver implementation.
func (r *Resolver) Game() generated.GameResolver { return &gameResolver{r} }

type gameResolver struct{ *Resolver }
