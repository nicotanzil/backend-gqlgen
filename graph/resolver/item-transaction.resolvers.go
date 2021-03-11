package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/nicotanzil/backend-gqlgen/database"
	"gorm.io/gorm/clause"

	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *queryResolver) GetPreviousTransactionData(ctx context.Context, itemID int) ([]*model.ItemTransaction, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var transactions []*model.ItemTransaction

	db.Order("created_at asc").
		Where("item_id = ?", itemID).
		Where("created_at BETWEEN (now() - '1 week'::interval) AND now()").
		Preload(clause.Associations).
		Find(&transactions)

	return transactions, nil
}
