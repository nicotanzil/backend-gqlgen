package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) BuyProfileBackground(ctx context.Context, userID int, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	// Validate wallet
	var user model.User
	db.Preload(clause.Associations).First(&user, userID)
	var item model.ProfileBackground
	db.Preload(clause.Associations).First(&item, id)

	if user.Points < item.Price {
		return false, nil
	}

	// Deduct wallet
	user.Points = user.Points - item.Price

	// Allocate item
	db.Exec("INSERT INTO user_profile_backgrounds VALUES (?, ?)", userID, id)

	db.Save(&user)

	return true, nil
}

func (r *queryResolver) ProfileBackgrounds(ctx context.Context) ([]*model.ProfileBackground, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var backgrounds []*model.ProfileBackground

	db.Find(&backgrounds)

	return backgrounds, nil
}

func (r *queryResolver) ExcludeProfileBackground(ctx context.Context, userID int) ([]*model.ProfileBackground, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var owned []int

	db.Raw("select profile_background_id from "+
		"user_profile_backgrounds where user_id = ?", userID).
		Scan(&owned)

	var backgrounds []*model.ProfileBackground

	db.Preload(clause.Associations).
		Where("id NOT IN ?", owned).
		Find(&backgrounds)

	return backgrounds, nil
}
