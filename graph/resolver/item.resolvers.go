package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var items []*model.Item

	db.Preload(clause.Associations).
		Find(&items)

	return items, nil
}

func (r *queryResolver) ItemsPaginate(ctx context.Context, userID int, gameID int, page int, keyword string) ([]*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var items []*model.Item

	keyword = "%" + keyword + "%"

	db.Preload(clause.Associations).
		Limit(providers.INVENTORY_PAGINATION).
		Offset(providers.INVENTORY_PAGINATION*(page-1)).
		Joins("JOIN user_items ui ON ui.item_id = items.id").
		Where("ui.user_id = ?", userID).
		Where("items.game_id = ?", gameID).
		Where("LOWER(items.name) LIKE LOWER(?)", keyword).
		Find(&items)

	return items, nil
}

func (r *queryResolver) GetTotalItems(ctx context.Context, userID int, gameID int, keyword string) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64
	keyword = "%" + keyword + "%"
	db.Model(model.Item{}).
		Joins("JOIN user_items ui ON ui.item_id = items.id").
		Where("ui.user_id = ?", userID).
		Where("items.game_id = ?", gameID).
		Where("LOWER(items.name) LIKE LOWER(?)", keyword).
		Count(&count)

	return int(count), nil
}

func (r *queryResolver) GetItemCategoriesByUser(ctx context.Context, userID int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	/**
	SELECT gm.id, gm.name, gm.banner
	FROM items i JOIN user_items ui ON i.id = ui.item_id
	JOIN games gm ON gm.id = i.game_id
	WHERE user_id = 1
	GROUP BY gm.id
	*/

	db.Model(model.Item{}).
		Select("gm.id, gm.name, gm.banner").
		Joins("JOIN user_items ui ON items.id = ui.item_id").
		Joins("JOIN games gm ON gm.id = items.game_id").
		Where("ui.user_id = ?", userID).
		Group("gm.id").
		Order("gm.id asc").
		Find(&games)

	return games, nil
}

func (r *queryResolver) GetItemByKeywordGame(ctx context.Context, userID int, gameID int, keyword string) ([]*model.Item, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var items []*model.Item

	/**
	SELECT *
	FROM items i JOIN user_items ui ON i.id = ui.item_id
	WHERE ui.user_id = 1 AND i.game_id = 1 AND i.name LIKE '%a%'
	*/

	keyword = "%" + keyword + "%"

	db.Model(model.Item{}).
		Joins("JOIN user_items ui ON items.id = ui.item_id").
		Where("ui.user_id = ?", userID).
		Where("items.game_id = ?", gameID).
		Where("LOWER(items.name) LIKE LOWER(?)", keyword).
		Find(&items)

	return items, nil
}
