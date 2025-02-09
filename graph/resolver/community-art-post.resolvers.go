package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateCommunityArtPost(ctx context.Context, input model.InputCommunityArtPost) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var post model.CommunityArtPost

	post = model.CommunityArtPost{
		Description: input.Description,
		Link:        input.Link,
		IsImage:     input.IsImage,
		User:        &model.User{ID: input.UserID},
	}

	db.Create(&post)
	return true, nil
}

func (r *mutationResolver) CommunityPostLike(ctx context.Context, postID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityPost model.CommunityArtPost

	db.First(&communityPost, "id = ?", postID)

	count := communityPost.Like
	count++
	communityPost.Like = count
	db.Save(&communityPost)

	return count, nil
}

func (r *mutationResolver) CommunityPostDislike(ctx context.Context, postID int) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityPost model.CommunityArtPost

	db.First(&communityPost, "id = ?", postID)

	count := communityPost.Like
	count--
	communityPost.Like = count
	db.Save(&communityPost)

	return count, nil
}

func (r *queryResolver) CommunityArtPosts(ctx context.Context) ([]*model.CommunityArtPost, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var communityArtPosts []*model.CommunityArtPost

	db.Order("id desc").
		Preload("User").
		Preload("Reviews.User").
		Find(&communityArtPosts)

	return communityArtPosts, nil
}
