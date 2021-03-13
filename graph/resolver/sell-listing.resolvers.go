package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
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
	fmt.Println(sellListing.ItemID)
	if sellListing.ItemID != 0 {
		// Item already exists in
		return false, nil
	}

	fmt.Println(sellListing.ItemID)
	var newSellListing model.SellListing

	newSellListing = model.SellListing{
		Item: &model.Item{ID: itemID},
		Sell: sell,
	}

	db.Create(&newSellListing)

	// Validate if there is any buy listing with higher or equal price with this sell listing
	var buyListingValidate model.BuyListing
	db.Preload(clause.Associations).
		Where("buy >= ?", newSellListing.Sell).
		First(&buyListingValidate)

	if buyListingValidate.BidID != 0 {
		var bid model.Bid
		var item model.Item
		db.Preload(clause.Associations).First(&item, newSellListing.Item.ID)
		db.Preload(clause.Associations).First(&bid, buyListingValidate.Bid.ID)
		var buyer model.User
		var seller model.User
		db.Preload(clause.Associations).First(&buyer, bid.User.ID)
		db.Preload(clause.Associations).First(&seller, item.User.ID)

		// Deduct balance from buyer's wallet (+tax)
		buyer.Balance = buyer.Balance - float64(newSellListing.Sell) + (float64(newSellListing.Sell) * 0.1)

		// Add balance to seller's wallet
		seller.Balance = seller.Balance + float64(newSellListing.Sell)

		db.Save(&buyer)
		db.Save(&seller)

		// Create Transaction
		var newItemTransaction model.ItemTransaction
		newItemTransaction = model.ItemTransaction{
			Item:   &model.Item{ID: newSellListing.Item.ID},
			Seller: &model.User{ID: seller.ID},
			Buyer:  &model.User{ID: buyer.ID},
			Price:  newSellListing.Sell,
		}
		db.Create(&newItemTransaction)

		// Remove the item from seller's inventory
		// Add the item to buyer's inventory
		// Change item's ownership
		item.User = &model.User{ID: buyer.ID}
		db.Save(&item)

		// Delete sellListing and buyListing
		db.Delete(model.SellListing{}, "item_id = ?", newSellListing.Item.ID)
		db.Delete(model.BuyListing{}, "bid_id = ?", buyListingValidate.Bid.ID)
	}

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

func (r *queryResolver) GetLowestSellListings(ctx context.Context) ([]*model.SellListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var sellListings []*model.SellListing

	db.Preload("Item.ItemType").
		Preload("Item.User").
		Order("sell asc").
		Limit(providers.DEFAULT_BUY_LISTING).
		Find(&sellListings)

	return sellListings, nil
}

func (r *queryResolver) GetSellListingsData(ctx context.Context, typeID int) ([]*model.LowestSellListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var lowest []*model.LowestSellListing

	//SELECT sell AS price, COUNT(sell) AS listing_count
	//FROM sell_listings sl
	//JOIN items i ON i.id = sl.item_id
	//JOIN item_types it ON i.item_type_id = it.id
	//WHERE it.id = 1
	//GROUP BY sell
	//ORDER BY sell asc
	//LIMIT 5

	db.Raw("SELECT sell AS price, COUNT(sell) AS listing_count"+
		"\nFROM sell_listings sl"+
		"\nJOIN items i ON i.id = sl.item_id"+
		"\nJOIN item_types it ON i.item_type_id = it.id"+
		"\nWHERE it.id = ?"+
		"\nGROUP BY sell"+
		"\nORDER BY sell asc"+
		"\nLIMIT 5", typeID).Scan(&lowest)

	return lowest, nil
}

func (r *queryResolver) GetSellListingsByUser(ctx context.Context, userID int, typeID int) ([]*model.SellListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var sellListings []*model.SellListing

	db.Preload(clause.Associations).
		Joins("JOIN items i ON sell_listings.item_id = i.id").
		Joins("JOIN users u ON i.user_id = u.id").
		Joins("JOIN item_types it ON i.item_type_id = it.id").
		Where("u.id = ?", userID).
		Where("it.id = ?", typeID).
		Find(&sellListings)

	return sellListings, nil
}
