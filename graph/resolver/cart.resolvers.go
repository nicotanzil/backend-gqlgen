package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertGameToCart(ctx context.Context, gameID int, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	cart := model.Cart{
		User: &model.User{ID: userID},
		Game: &model.Game{ID: gameID},
	}
	db.Create(&cart)

	return true, nil
}

func (r *mutationResolver) RemoveGameFromCart(ctx context.Context, gameID int, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var cart model.Cart

	db.Where("user_id = ?", userID).Where("game_id = ?", gameID).Delete(&cart)
	return true, nil
}

func (r *queryResolver) Carts(ctx context.Context) ([]*model.Cart, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var carts []*model.Cart

	db.Preload("Game.Promo").Preload("User").Find(&carts)

	return carts, nil
}

func (r *queryResolver) GetCartGamesByUserID(ctx context.Context, id int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var carts []*model.Cart
	var games []*model.Game

	db.Where("user_id = ?", id).Preload("Game.Promo").Preload("User").Find(&carts)
	var gameIds []int

	for i := 0; i < len(carts); i++ {
		gameIds = append(gameIds, carts[i].Game.ID)
	}
	db.Preload(clause.Associations).Where("id IN ?", gameIds).Find(&games)

	return games, nil
}
