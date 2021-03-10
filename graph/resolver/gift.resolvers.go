package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateGift(ctx context.Context, gift *model.NewGift) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var insertGift model.Gift

	insertGift = model.Gift{
		Sender:     &model.User{ID: gift.SenderID},
		Receiver:   &model.User{ID: gift.ReceiverID},
		FirstName:  gift.FirstName,
		Message:    gift.Message,
		Sentiment:  gift.Sentiment,
		Signature:  gift.Signature,
		IsComplete: false,
	}

	db.Create(&insertGift)

	return true, nil
}

func (r *queryResolver) Gifts(ctx context.Context) ([]*model.Gift, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var gifts []*model.Gift

	db.Preload(clause.Associations).Find(&gifts)

	return gifts, nil
}

func (r *queryResolver) GetGiftBySenderID(ctx context.Context, id int) (*model.Gift, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var gift model.Gift

	db.Preload(clause.Associations).Where("sender_id = ?", id).First(&gift)

	return &gift, nil
}

func (r *queryResolver) GetGiftNotificationCount(ctx context.Context, receiverID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(model.Gift{}).
		Where("receiver_id = ?", receiverID).
		Where("is_open = false").
		Where("is_complete = true").
		Count(&count)

	return int(count), nil
}
