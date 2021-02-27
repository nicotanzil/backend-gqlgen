package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) InsertGameVideo(ctx context.Context, gameVideos []*model.InputGameVideo) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(gameVideos); i++ {
		var newGameVideos model.GameVideo

		newGameVideos = model.GameVideo{
			GameID: gameVideos[i].GameID,
			Link:   gameVideos[i].Link,
		}

		db.Create(&newGameVideos)
	}

	return true, nil
}

func (r *mutationResolver) UpdateGameVideo(ctx context.Context, id int, link string) (*model.GameVideo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var gameVideo model.GameVideo

	db.Where("id = ?", id).First(&gameVideo)

	gameVideo.Link = link

	return &gameVideo, nil
}

func (r *queryResolver) GameVideos(ctx context.Context) ([]*model.GameVideo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var gameVideos []*model.GameVideo

	db.Find(&gameVideos)

	return gameVideos, nil
}
