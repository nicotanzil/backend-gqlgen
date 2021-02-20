package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	genres := []*model.Genre{}
	tags := []*model.Tag{}

	for i := 0; i < len(input.Genres); i++ {
		curr := model.Genre{ID: input.Genres[i].ID}
		genres = append(genres, &curr)
	}

	for i := 0; i < len(input.Tags); i++ {
		curr := model.Tag{ID: input.Tags[i].ID}
		tags = append(tags, &curr)
	}

	var newGame model.Game

	newGame = model.Game{
		Name:               input.Name,
		Description:        input.Description,
		ReleaseDate:        time.Now(),
		Genres:             genres,
		Tags:               tags,
		OriginalPrice:      input.OriginalPrice,
		OnSale:             input.OnSale,
		DiscountPercentage: input.DiscountPercentage,
		GamePlayHour:       0,
		GameReviews:        nil,
		Developers:         nil,
		PublisherID:        input.PublisherID,
		SystemID:           input.SystemID,
		Users:              nil,
		Banner:             input.Banner,
		Video:              input.Video,
		Image1:             input.Image1,
		Image2:             input.Image2,
		Image3:             input.Image3,
		Image4:             input.Image4,
	}

	db.Create(&newGame)
	return &newGame, nil
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var game model.Game

	db.Where("id = ?", id).First(&game)

	db.Delete(&model.Game{}, id)
	return &game, nil
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

	db.Preload(clause.Associations).Limit(providers.ADMIN_GAME_PAGINATION).Offset(providers.ADMIN_GAME_PAGINATION * (*page - 1)).Find(&games)

	return games, nil
}

func (r *queryResolver) GetTotalGame(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var count int64

	db.Model(&model.Game{}).Count(&count)

	return int(count), nil
}
