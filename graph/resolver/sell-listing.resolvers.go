package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateSellListing(ctx context.Context, itemID int, sell int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var sellListing model.SellListing

	db.First(&sellListing, "item_id = ?", itemID)

	if sellListing.ItemID != 0 {
		// Item already exists in
		return false, nil
	}

	var newSellListing model.SellListing

	newSellListing = model.SellListing{
		Item: &model.Item{ID: itemID},
		Sell: sell,
	}

	db.Create(&newSellListing)

	return true, nil
}

func (r *queryResolver) TopTransactionItemTypes(ctx context.Context, page int) ([]*model.TopTransactionItem, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var types []*model.TopTransactionItem

	/**
	SELECT ity.id AS item_type_id, ity.name AS item_name, ity.link AS item_link, COUNT(it.item_id) AS transaction_count
	FROM item_transactions it
	JOIN items i ON it.item_id = i.id
	JOIN item_types ity ON ity.id = i.item_type_id
	WHERE ity.id IN (
		SELECT i2.item_type_id
		FROM sell_listings sl2 JOIN items i2 ON sl2.item_id = i2.id
		GROUP BY i2.item_type_id
	)
	GROUP BY ity.id, ity.name, ity.link
	ORDER BY COUNT(it.item_id) DESC

	*/

	db.Raw("SELECT ity.id AS item_type_id, ity.name AS item_name, ity.link AS item_link, gm.name AS game_name, COUNT(it.item_id) AS transaction_count"+
		"\nFROM item_transactions it "+
		"\nJOIN items i ON it.item_id = i.id"+
		"\nJOIN item_types ity ON ity.id = i.item_type_id"+
		"\nJOIN games gm ON gm.id = ity.game_id"+
		"\nWHERE ity.id IN ("+
		"\n\tSELECT i2.item_type_id"+
		"\n\tFROM sell_listings sl2 JOIN items i2 ON sl2.item_id = i2.id"+
		"\n\tGROUP BY i2.item_type_id"+
		"\n)"+
		"\nGROUP BY ity.id, ity.name, ity.link, gm.name"+
		"\nORDER BY COUNT(it.item_id) DESC"+
		"\nOFFSET ?"+
		"LIMIT ?", providers.MARKET_LIST_PAGINATION*(page-1), providers.MARKET_LIST_PAGINATION).
		Scan(&types)

	return types, nil
}

func (r *queryResolver) GetTotalTopTransactionItemTypes(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Raw("SELECT ity.id AS item_type_id, ity.name AS item_name, ity.link AS item_link, gm.name AS game_name, COUNT(it.item_id) AS transaction_count" +
		"\nFROM item_transactions it " +
		"\nJOIN items i ON it.item_id = i.id" +
		"\nJOIN item_types ity ON ity.id = i.item_type_id" +
		"\nJOIN games gm ON gm.id = ity.game_id" +
		"\nWHERE ity.id IN (" +
		"\n\tSELECT i2.item_type_id" +
		"\n\tFROM sell_listings sl2 JOIN items i2 ON sl2.item_id = i2.id" +
		"\n\tGROUP BY i2.item_type_id" +
		"\n)" +
		"\nGROUP BY ity.id, ity.name, ity.link, gm.name" +
		"\nORDER BY COUNT(it.item_id) DESC").
		Count(&count)

	return int(count), nil
}
