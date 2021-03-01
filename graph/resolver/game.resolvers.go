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
	developers := []*model.Developer{}

	for i := 0; i < len(input.Genres); i++ {
		curr := model.Genre{ID: input.Genres[i].ID}
		genres = append(genres, &curr)
	}

	for i := 0; i < len(input.Tags); i++ {
		curr := model.Tag{ID: input.Tags[i].ID}
		tags = append(tags, &curr)
	}

	for i := 0; i < len(input.Developers); i++ {
		curr := model.Developer{ID: input.Developers[i].ID}
		developers = append(developers, &curr)
	}

	var newGame model.Game

	newGame = model.Game{
		Name:          input.Name,
		Description:   input.Description,
		ReleaseDate:   time.Now(),
		Genres:        genres,
		Tags:          tags,
		OriginalPrice: input.OriginalPrice,
		PromoID:       input.Promo.ID,
		GamePlayHour:  0,
		GameReviews:   nil,
		Developers:    developers,
		PublisherID:   input.PublisherID,
		SystemID:      input.SystemID,
		Users:         nil,
	}

	db.Create(&newGame)
	return &newGame, nil
}

func (r *mutationResolver) UpdateGame(ctx context.Context, id int, input model.NewGame) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	genres := []*model.Genre{}
	tags := []*model.Tag{}
	developers := []*model.Developer{}

	for i := 0; i < len(input.Genres); i++ {
		curr := model.Genre{ID: input.Genres[i].ID}
		genres = append(genres, &curr)
	}

	for i := 0; i < len(input.Tags); i++ {
		curr := model.Tag{ID: input.Tags[i].ID}
		tags = append(tags, &curr)
	}

	for i := 0; i < len(input.Developers); i++ {
		curr := model.Developer{ID: input.Developers[i].ID}
		developers = append(developers, &curr)
	}

	var game model.Game

	db.Where("id = ?", id).First(&game)

	game.Name = input.Name
	game.Description = input.Description
	game.Genres = genres
	game.Tags = tags
	game.OriginalPrice = input.OriginalPrice
	game.PromoID = input.Promo.ID
	game.Developers = developers
	game.PublisherID = input.PublisherID
	game.SystemID = input.SystemID

	db.Save(&game)

	return &game, nil
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	db.Exec("DELETE FROM game_images WHERE game_id = ?", id)
	db.Exec("DELETE FROM game_videos WHERE game_id = ?", id)

	var game model.Game

	db.Where("id = ?", id).First(&game)

	db.Delete(&model.Game{}, id)
	return &game, nil
}

func (r *mutationResolver) InsertGameBanner(ctx context.Context, id int, link string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var game model.Game

	db.Where("id = ?", id).First(&game)

	game.Banner = link
	db.Save(&game)
	return true, nil
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

func (r *queryResolver) GameByID(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var game model.Game

	db.Preload(clause.Associations).Where("id = ?", id).First(&game)

	return &game, nil
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

func (r *queryResolver) GameSearch(ctx context.Context, keyword string) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*model.Game

	keyword = "%" + keyword + "%"
	db.Preload(clause.Associations).Limit(providers.GAME_SEARCH_LIMIT).Where("name LIKE ?", keyword).Find(&games)

	return games, nil
}

func (r *queryResolver) GameSearchPage(ctx context.Context, keyword string, page int, price int, tag []*model.InputTag) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var games []*model.Game
	var tagsId []int

	if len(tag) > 0 {
		for i := 0; i < len(tag); i++ {
			tagsId = append(tagsId, tag[i].ID)
		}

		keyword = "%" + keyword + "%"
		db.Preload(clause.Associations).
			Limit(providers.GAME_SEARCH_PAGE_LIMIT).
			Offset(providers.GAME_SEARCH_PAGE_LIMIT*(page-1)).
			Joins("JOIN game_tags ON game_tags.game_id = games.id").
			Where("name LIKE ?", keyword).
			Where("original_price <= ?", price).
			Where("game_tags.tag_id IN ?", tagsId).
			Find(&games)
	} else {
		keyword = "%" + keyword + "%"
		db.Preload(clause.Associations).
			Limit(providers.GAME_SEARCH_PAGE_LIMIT).
			Offset(providers.GAME_SEARCH_PAGE_LIMIT*(page-1)).
			Where("name LIKE ?", keyword).
			Where("original_price <= ?", price).
			Find(&games)
	}

	return games, nil
}
