package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/nicotanzil/backend-gqlgen/app/errorlist"
	"github.com/nicotanzil/backend-gqlgen/app/middleware"
	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (string, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var user model.User

	fmt.Println(input.AccountName, input.Password)

	// Get user by profile name
	db.Where("account_name = ?", input.AccountName).First(&user)

	if user.AccountName == "" {
		fmt.Println("User not found")
		return "", errorlist.Auth_Invalid //user not found
	}

	// Compare password
	if !providers.CheckPasswordHash(input.Password, user.Password) {
		fmt.Println("Invalid password")
		return "", errorlist.Auth_Invalid //invalid password
	}

	//Create JWT Token
	token, err2 := providers.GenerateToken(user)

	if err2 != nil {
		fmt.Println("Generate token error")
		return "", err2
	}

	cookie := &http.Cookie{
		Name:  "access_token",
		Value: token,
		//Expires:  time.Now().Add(time.Minute * 5), // expiration time 5 minutes
		Expires:  time.Time{},
		HttpOnly: true,
	}

	w := *middleware.WForContext(ctx)
	http.SetCookie(w, cookie)

	// Send the JWT token
	return token, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	w := *middleware.WForContext(ctx)
	http.SetCookie(w, cookie)

	return true, nil
}

func (r *queryResolver) GetUserAuth(ctx context.Context) (*model.User, error) {
	user := middleware.ForContext(ctx)

	if user == nil {
		return &model.User{}, nil //user not found login as guest
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var fetchUser model.User

	db.Where("id = ?", user.ID).First(&fetchUser)

	return &fetchUser, nil
}
