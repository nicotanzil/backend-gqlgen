package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) BuyChatSticker(ctx context.Context, userID int, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	// Validate wallet
	var user model.User
	db.Preload(clause.Associations).First(&user, userID)
	var item model.ChatSticker
	db.Preload(clause.Associations).First(&item, id)

	if user.Points < item.Price {
		return false, nil
	}

	// Deduct wallet
	user.Points = user.Points - item.Price

	// Allocate item
	db.Exec("INSERT INTO user_chat_stickers VALUES (?, ?)", userID, id)

	db.Save(&user)

	return true, nil
}

func (r *queryResolver) ChatStickers(ctx context.Context) ([]*model.ChatSticker, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var chatStickers []*model.ChatSticker

	db.Preload(clause.Associations).Find(&chatStickers)

	return chatStickers, nil
}

func (r *queryResolver) ExcludeChatStickers(ctx context.Context, userID int) ([]*model.ChatSticker, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var owned []int

	db.Raw("select chat_sticker_id from "+
		"user_chat_stickers where user_id = ?", userID).
		Scan(&owned)

	var stickers []*model.ChatSticker

	db.Preload(clause.Associations).
		Where("id NOT IN ?", owned).
		Find(&stickers)

	return stickers, nil
}
