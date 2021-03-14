package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *queryResolver) GetNewItemNotificationCount(ctx context.Context, userID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(&model.NewItemNotification{}).
		Joins("JOIN items i ON new_item_notifications.item_id = i.id").
		Where("i.user_id = ?", userID).
		Where("new_item_notifications.is_open = ?", false).
		Count(&count)

	return int(count), nil
}
