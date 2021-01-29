package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/app/errorlist"
	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.TokenDetail, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	// Get user by profile name
	db.Where("account_name = ?", input.AccountName).First(&user)

	if user.ProfileName == "" {
		return nil, errorlist.Auth_Invalid //user not found
	}

	// Compare password
	if !providers.CheckPasswordHash(input.Password, user.Password) {
		return nil, errorlist.Auth_Invalid //invalid password
	}

	//Create JWT Token
	td, error := providers.GenerateToken(user)

	if error != nil {
		return nil, error
	}

	// Send the JWT token
	return td, nil
}
