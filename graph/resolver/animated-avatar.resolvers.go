package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) BuyAnimatedAvatars(ctx context.Context, userID int, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	// Validate wallet
	var user model.User
	db.Preload(clause.Associations).First(&user, userID)
	var item model.AnimatedAvatar
	db.Preload(clause.Associations).First(&item, id)

	if user.Points < item.Price {
		return false, nil
	}

	// Deduct wallet
	user.Points = user.Points - item.Price

	// Allocate item
	db.Exec("INSERT INTO user_animated_avatars VALUES (?, ?)", userID, id)

	db.Save(&user)

	return true, nil
}

func (r *queryResolver) AnimatedAvatars(ctx context.Context) ([]*model.AnimatedAvatar, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var animatedAvatars []*model.AnimatedAvatar

	db.Preload(clause.Associations).Find(&animatedAvatars)

	return animatedAvatars, nil
}

func (r *queryResolver) ExcludeAnimatedAvatars(ctx context.Context, userID int) ([]*model.AnimatedAvatar, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var owned []int

	db.Raw("select animated_avatar_id from "+
		"user_animated_avatars where user_id = ?", userID).
		Scan(&owned)

	var avatars []*model.AnimatedAvatar

	db.Preload(clause.Associations).
		Where("id NOT IN ?", owned).
		Find(&avatars)

	return avatars, nil
}
