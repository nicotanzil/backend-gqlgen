package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) InsertGameImage(ctx context.Context, gameImages []*model.InputGameImage) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(gameImages); i++ {
		var newGameImage model.GameImage

		newGameImage = model.GameImage{
			GameID: gameImages[i].GameID,
			Link:   gameImages[i].Link,
		}

		db.Create(&newGameImage)
	}

	return true, nil
}

func (r *mutationResolver) UpdateGameImage(ctx context.Context, id int, link string) (*model.GameImage, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var gameImage model.GameImage

	db.Where("id = ?", id).First(&gameImage)

	gameImage.Link = link

	return &gameImage, nil
}

func (r *queryResolver) GameImages(ctx context.Context) ([]*model.GameImage, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var gameImages []*model.GameImage

	db.Find(&gameImages)

	return gameImages, nil
}
