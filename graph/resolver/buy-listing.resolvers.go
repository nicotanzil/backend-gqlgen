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

func (r *mutationResolver) CreateBuyListing(ctx context.Context, bidID int, buy int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var buyListing model.BuyListing

	db.Preload(clause.Associations).First(&buyListing, "bid_id = ?", bidID)

	//if buyListing.BidID != 0 {
	//	// Item already exists in
	//	return false, nil
	//}

	var newBuyListing model.BuyListing

	newBuyListing = model.BuyListing{
		Bid: &model.Bid{ID: bidID},
		Buy: buy,
	}

	db.Create(&newBuyListing)

	// Validate if there is any sell listing with lower or equal price with this buy listing
	var sellListingValidate model.SellListing
	db.Preload(clause.Associations).
		Where("sell <= ?", newBuyListing.Buy).
		First(&sellListingValidate)

	if sellListingValidate.ItemID != 0 {
		var bid model.Bid
		var item model.Item
		db.Preload(clause.Associations).First(&bid, newBuyListing.Bid.ID)
		db.Preload(clause.Associations).First(&item, sellListingValidate.Item.ID)
		var buyer model.User
		var seller model.User
		db.Preload(clause.Associations).First(&buyer, bid.User.ID)
		db.Preload(clause.Associations).First(&seller, item.User.ID)

		// Deduct balance from buyer's wallet (+tax)
		buyer.Balance = buyer.Balance - float64(sellListingValidate.Sell) + (float64(sellListingValidate.Sell) * 0.1)

		// Add balance to seller's wallet
		seller.Balance = seller.Balance + float64(sellListingValidate.Sell)

		db.Save(&buyer)
		db.Save(&seller)

		// Create Transaction
		var newItemTransaction model.ItemTransaction
		newItemTransaction = model.ItemTransaction{
			Item:   &model.Item{ID: sellListingValidate.Item.ID},
			Seller: &model.User{ID: seller.ID},
			Buyer:  &model.User{ID: buyer.ID},
			Price:  sellListingValidate.Sell,
		}
		db.Create(&newItemTransaction)

		// Remove the item from seller's inventory
		// Add the item to buyer's inventory
		// Change item's ownership
		item.User = &model.User{ID: buyer.ID}
		db.Save(&item)

		// Delete sellListing and buyListing
		db.Delete(model.SellListing{}, "item_id = ?", sellListingValidate.Item.ID)
		db.Delete(model.BuyListing{}, "bid_id = ?", newBuyListing.Bid.ID)
	}

	return true, nil
}

func (r *queryResolver) GetHighestBuyListings(ctx context.Context) ([]*model.BuyListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var buyListings []*model.BuyListing

	db.Preload("Bid.ItemType").
		Preload("Bid.User").
		Order("buy desc").
		Limit(providers.DEFAULT_BUY_LISTING).
		Find(&buyListings)

	return buyListings, nil
}

func (r *queryResolver) GetBuyListingsData(ctx context.Context, typeID int) ([]*model.HighestBuyListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var highest []*model.HighestBuyListing

	//SELECT buy AS price, COUNT(buy) AS listing_count
	//FROM buy_listings bl
	//JOIN bids b ON b.id = bl.bid_id
	//JOIN item_types it ON b.item_type_id = it.id
	//WHERE it.id = 1
	//GROUP BY buy
	//ORDER BY buy desc
	//LIMIT 5

	db.Raw("SELECT buy AS price, COUNT(buy) AS listing_count"+
		"\nFROM buy_listings bl"+
		"\nJOIN bids b ON b.id = bl.bid_id"+
		"\nJOIN item_types it ON b.item_type_id = it.id"+
		"\nWHERE it.id = ?"+
		"\nGROUP BY buy"+
		"\nORDER BY buy desc"+
		"\nLIMIT 5", typeID).Scan(&highest)

	return highest, nil
}

func (r *queryResolver) GetBuyListingsByUser(ctx context.Context, userID int, typeID int) ([]*model.BuyListing, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var buyListings []*model.BuyListing

	db.Preload(clause.Associations).
		Joins("JOIN bids b ON buy_listings.bid_id = b.id").
		Joins("JOIN users u ON b.user_id = u.id").
		Joins("JOIN item_types it ON b.item_type_id = it.id").
		Where("u.id = ?", userID).
		Where("it.id = ?", typeID).
		Find(&buyListings)

	return buyListings, nil
}
