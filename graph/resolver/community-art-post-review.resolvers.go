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

func (r *mutationResolver) AddCommentByPostID(ctx context.Context, postID int, userID int, comment string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityArtPostReview model.CommunityArtPostReview

	communityArtPostReview = model.CommunityArtPostReview{
		Post:    &model.CommunityArtPost{ID: postID},
		User:    &model.User{ID: userID},
		Comment: comment,
	}

	db.Create(&communityArtPostReview)
	return true, nil
}

func (r *queryResolver) CommunityArtPostReviews(ctx context.Context) ([]*model.CommunityArtPostReview, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityArtPostReviews []*model.CommunityArtPostReview

	db.Preload(clause.Associations).Find(&communityArtPostReviews)

	return communityArtPostReviews, nil
}

func (r *queryResolver) GetCommunityArtPostReviewsByPostID(ctx context.Context, postID int, page int) ([]*model.CommunityArtPostReview, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityArtPostReviews []*model.CommunityArtPostReview

	db.Preload(clause.Associations).
		Limit(providers.COMMUNITY_ARTWORK_PAGINATION).
		Offset(providers.COMMUNITY_ARTWORK_PAGINATION*(page-1)).
		Where("post_id = ?", postID).Find(&communityArtPostReviews)

	return communityArtPostReviews, nil
}

func (r *queryResolver) GetTotalReviewsByPostID(ctx context.Context, postID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(model.CommunityArtPostReview{}).Where("post_id = ?", postID).Count(&count)

	return int(count), nil
}
