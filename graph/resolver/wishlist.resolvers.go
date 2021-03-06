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

func (r *mutationResolver) InsertGameToWishlist(ctx context.Context, gameID int, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	wishlist := model.Wishlist{
		User: &model.User{ID: userID},
		Game: &model.Game{ID: gameID},
	}

	db.Create(&wishlist)

	var game model.Game
	var user model.User

	db.Preload(clause.Associations).First(&game, gameID)
	db.Preload(clause.Associations).First(&user, userID)

	// Send email if game on sale
	fmt.Println("ID: ", game.Promo.ID)
	if game.Promo.ID != 1 {
		// Game on promo
		model.SendEmailPromo(game, user)
	}

	return true, nil
}

func (r *mutationResolver) RemoveGameFromWishlist(ctx context.Context, gameID int, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var wishlist model.Wishlist

	db.Where("user_id = ?", userID).Where("game_id = ?", gameID).Delete(&wishlist)
	return true, nil
}

func (r *queryResolver) Wishlists(ctx context.Context) ([]*model.Wishlist, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var wishlists []*model.Wishlist

	db.Preload("Game.Promo").Preload("User").Find(&wishlists)

	return wishlists, nil
}

func (r *queryResolver) GetWishlistByUserID(ctx context.Context, id int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var wishlists []*model.Wishlist
	var games []*model.Game

	db.Where("user_id = ?", id).Preload("Game.Promo").Preload("User").Find(&wishlists)
	var gameIds []int

	for i := 0; i < len(wishlists); i++ {
		gameIds = append(gameIds, wishlists[i].Game.ID)
	}
	db.Preload(clause.Associations).Where("id IN ?", gameIds).Find(&games)

	return games, nil
}
