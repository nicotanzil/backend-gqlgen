package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gorm.io/gorm/clause"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) BuyAvatarFrame(ctx context.Context, userID int, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	// Validate wallet
	var user model.User
	db.Preload(clause.Associations).First(&user, userID)
	var item model.AvatarFrame
	db.Preload(clause.Associations).First(&item, id)

	if user.Points < item.Price {
		return false, nil
	}

	// Deduct wallet
	user.Points = user.Points - item.Price

	// Allocate item
	db.Exec("INSERT INTO user_avatar_frames VALUES (?, ?)", userID, id)

	db.Save(&user)

	return true, nil
}

func (r *queryResolver) AvatarFrames(ctx context.Context) ([]*model.AvatarFrame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var frames []*model.AvatarFrame

	db.Find(&frames)

	return frames, nil
}

func (r *queryResolver) ExcludeAvatarFrames(ctx context.Context, userID int) ([]*model.AvatarFrame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var owned []int

	db.Raw("select avatar_frame_id from "+
		"user_avatar_frames where user_id = ?", userID).
		Scan(&owned)

	var frames []*model.AvatarFrame

	db.Preload(clause.Associations).
		Where("id NOT IN ?", owned).
		Find(&frames)

	return frames, nil
}
