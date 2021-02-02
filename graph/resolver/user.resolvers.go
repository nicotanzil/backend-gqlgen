package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	firebase_data "github.com/nicotanzil/backend-gqlgen/app/firebase-data"
	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/generated"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.NewUser, otp *model.NewOtp) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var newUser model.User

	newUser = model.User{
		AccountName:       user.AccountName,
		ProfileName:       user.AccountName,
		RealName:          "",
		Email:             otp.Email,
		Password:          providers.HashPassword(user.Password),
		Balance:           0,
		CustomURL:         user.AccountName,
		Summary:           "No information given.",
		Avatar:            firebase_data.Avatar,
		ProfileBackground: firebase_data.ProfileBackground,
		Country:           &model.Country{ID: otp.CountryId},
	}

	db.Create(&newUser)
	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user *model.UpdateUser) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var u model.User

	db.Where("account_name = ?", user.AccountName).First(&u)

	u.ProfileName = user.ProfileName
	u.RealName = user.RealName
	u.CustomURL = user.CustomURL
	u.Summary = user.Summary
	u.Avatar = user.Avatar
	u.AvatarFrames = user.AvatarFrames
	u.ProfileBackground = user.ProfileBackground
	u.MiniProfileBackground = user.MiniProfileBackground
	u.Theme = user.Theme
	u.CountryID = user.CountryID

	db.Save(&u)

	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var users []*model.User

	//db.Find(&users)	//populate the users variable with `SELECT * FROM users`

	db.Preload("Country").Preload("Games").Preload("Friends").Find(&users) //populate the users with each corresponding countries

	return users, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, input *string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	db.Preload("Country").Preload("Games").Preload("Friends").Where("id = ?", input).First(&user)

	return &user, nil
}

func (r *queryResolver) GetUserByURL(ctx context.Context, input *string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	db.Preload("Country").Preload("Games").Preload("Friends").Where("custom_url = ?", input).First(&user)

	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
