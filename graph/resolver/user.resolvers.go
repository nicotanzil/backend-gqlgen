package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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
		AccountName:    user.AccountName,
		ProfileName:    user.AccountName,
		RealName:       "",
		Email:          otp.Email,
		Password:       providers.HashPassword(user.Password),
		Balance:        0,
		CustomURL:      user.AccountName,
		ProfilePicture: "",
		CountryID:      otp.CountryId,
	}

	db.Create(&newUser)
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var users []*model.User

	//db.Find(&users)	//populate the users variable with `SELECT * FROM users`

	db.Find(&users) //populate the users with each corresponding countries

	return users, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, input *string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	db.Where("id = ?", input).First(&user)

	return &user, nil
}

func (r *userResolver) Country(ctx context.Context, obj *model.User) (*model.Country, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var country model.Country

	db.Where("id = ?", obj.CountryID).First(&country)

	return &country, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
