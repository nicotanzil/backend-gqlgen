package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CommunityReviewHelpful(ctx context.Context, reviewID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityGameReview model.CommunityGameReview

	db.First(&communityGameReview, "id = ?", reviewID)

	count := communityGameReview.HelpfulCount
	count++
	communityGameReview.HelpfulCount = count
	db.Save(&communityGameReview)

	return count, nil
}

func (r *mutationResolver) CommunityReviewNotHelpful(ctx context.Context, reviewID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityGameReview model.CommunityGameReview

	db.First(&communityGameReview, "id = ?", reviewID)

	count := communityGameReview.HelpfulCount
	count--
	communityGameReview.HelpfulCount = count
	db.Save(&communityGameReview)

	return count, nil
}

func (r *queryResolver) CommunityGameReviews(ctx context.Context) ([]*model.CommunityGameReview, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var gameReviews []*model.CommunityGameReview

	db.Order("id").
		Preload(clause.Associations).Find(&gameReviews)

	return gameReviews, nil
}
