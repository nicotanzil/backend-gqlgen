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

func (r *mutationResolver) AddCommentByReviewID(ctx context.Context, reviewID int, userID int, comment string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityGameReviewComment model.CommunityGameReviewComment

	communityGameReviewComment = model.CommunityGameReviewComment{
		User:        &model.User{ID: userID},
		Review:      &model.CommunityGameReview{ID: reviewID},
		Description: comment,
	}

	db.Create(&communityGameReviewComment)
	return true, nil
}

func (r *queryResolver) CommunityGameReviewComments(ctx context.Context) ([]*model.CommunityGameReviewComment, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityGameReviewComments []*model.CommunityGameReviewComment

	db.Preload(clause.Associations).Find(&communityGameReviewComments)
	return communityGameReviewComments, nil
}

func (r *queryResolver) GetCommunityGameReviewCommentByReviewID(ctx context.Context, reviewID int, page int) ([]*model.CommunityGameReviewComment, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityGameReviewComment []*model.CommunityGameReviewComment

	db.Preload(clause.Associations).
		Limit(providers.COMMUNITY_GAME_REVIEW_PAGINATION).
		Offset(providers.COMMUNITY_GAME_REVIEW_PAGINATION*(page-1)).
		Where("review_id = ?", reviewID).
		Find(&communityGameReviewComment)

	return communityGameReviewComment, nil
}

func (r *queryResolver) GetTotalCommentByReviewID(ctx context.Context, reviewID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(model.CommunityGameReviewComment{}).Where("review_id = ?", reviewID).Count(&count)

	return int(count), nil
}
