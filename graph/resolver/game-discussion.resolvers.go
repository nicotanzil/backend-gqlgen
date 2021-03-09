package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) InsertDiscussionThread(ctx context.Context, title string, description string, gameID int, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var discussion model.GameDiscussion

	discussion = model.GameDiscussion{
		Title:       title,
		Description: description,
		Game:        &model.Game{ID: gameID},
		User:        &model.User{ID: userID},
	}

	db.Create(&discussion)
	return true, nil
}

func (r *queryResolver) GameDiscussions(ctx context.Context) ([]*model.GameDiscussion, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var discussions []*model.GameDiscussion

	db.Preload(clause.Associations).
		Find(&discussions)

	return discussions, nil
}

func (r *queryResolver) GetGameDiscussionByDiscussionID(ctx context.Context, id int) (*model.GameDiscussion, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var review model.GameDiscussion

	db.Preload("Game").
		Preload("User").
		Preload("Replies.User").
		Where("id = ?", id).
		First(&review)

	return &review, nil
}
