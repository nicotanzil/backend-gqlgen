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
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.NewUser, otp *model.NewOtp) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var newUser model.User

	newUser = model.User{
		AccountName: user.AccountName,
		ProfileName: user.AccountName,
		RealName:    "",
		Email:       otp.Email,
		Password:    providers.HashPassword(user.Password),
		Balance:     0,
		CustomURL:   user.AccountName,
		Summary:     "No information given.",
		Avatar:      firebase_data.Avatar,
		ProfileBackgrounds: []*model.ProfileBackground{
			{ID: 1},
		},
		ProfileBackgroundID: 1,
		MiniProfileBackgrounds: []*model.MiniProfileBackground{
			{ID: 1},
		},
		MiniProfileBackgroundID: 1,
		AvatarFrames: []*model.AvatarFrame{
			{ID: 1},
		},
		AvatarFrameID: 1,
		ThemeID:       1,
		Country:       &model.Country{ID: otp.CountryId},
	}

	db.Create(&newUser)
	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user model.UpdateUser) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var u model.User

	db.Where("account_name = ?", user.AccountName).First(&u)

	u.ProfileName = user.ProfileName
	u.RealName = user.RealName
	u.CustomURL = user.CustomURL
	u.Summary = user.Summary
	u.Avatar = user.Avatar
	u.AvatarFrameID = user.AvatarFrameID
	u.ProfileBackgroundID = user.ProfileBackgroundID
	u.MiniProfileBackgroundID = user.MiniProfileBackgroundID
	u.ThemeID = user.ThemeID
	u.FeaturedBadgeID = user.FeaturedBadgeID
	u.CountryID = user.CountryID

	db.Save(&u)

	return true, nil
}

func (r *mutationResolver) UpdateAccountSuspension(ctx context.Context, id int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var user model.User
	var request model.SuspensionRequest

	db.Preload(clause.Associations).Where("id = ?", id).First(&user)

	if user.IsSuspend {
		user.IsSuspend = false
		db.Where("user_id = ?", user.ID).Delete(&request)
	} else {
		user.IsSuspend = true
	}
	db.Save(&user)
	return true, nil
}

func (r *mutationResolver) AddFriend(ctx context.Context, userID int, friendID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var user model.User
	var friend model.User

	db.Preload(clause.Associations).First(&user, userID)
	db.Preload(clause.Associations).First(&friend, friendID)

	user.Friends = []*model.User{{ID: friendID}}
	friend.Friends = []*model.User{{ID: userID}}
	db.Save(&user)
	db.Save(&friend)

	return true, nil
}

func (r *mutationResolver) UpdateAccountGeneral(ctx context.Context, accountName string, profileName string, realName string, customURL string, countryID int, summary string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var u model.User

	db.Where("account_name = ?", accountName).First(&u)

	u.ProfileName = profileName
	u.RealName = realName
	u.CustomURL = customURL
	u.CountryID = countryID
	u.Summary = summary

	db.Save(&u)

	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var users []*model.User

	//db.Find(&users)	//populate the users variable with `SELECT * FROM users`

	db.Preload(clause.Associations).
		Find(&users) //populate the users with each corresponding countries

	return users, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id *int) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var user model.User

	db.Preload(clause.Associations).Where("id = ?", id).First(&user)

	return &user, nil
}

func (r *queryResolver) GetUseByAccountName(ctx context.Context, accountName string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var user model.User

	db.Preload(clause.Associations).Where("account_name = ?", accountName).First(&user)

	return &user, nil
}

func (r *queryResolver) GetTotalUser(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(&model.User{}).Count(&count)

	return int(count), nil
}

func (r *queryResolver) GetUserByURL(ctx context.Context, input *string) (*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var user model.User

	db.Preload("AvatarFrame").
		Preload("ProfileBackground").
		Preload("MiniProfileBackground").
		Preload("Theme").
		Preload("Country").
		Preload("FeaturedBadge").
		Preload("Badges").
		Preload("Games").
		Preload("Friends").
		Preload("Friends.MiniProfileBackground").
		Preload("Friends.AvatarFrame").
		Preload("Friends.ProfileBackground").
		Preload("Friends.FeaturedBadge").
		Where("custom_url = ?", input).First(&user)

	return &user, nil
}

func (r *queryResolver) GetUserPaginationAdmin(ctx context.Context, page int) ([]*model.User, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promos []*model.User

	db.Preload(clause.Associations).Limit(providers.ADMIN_USER_PAGINATION).Offset(providers.ADMIN_USER_PAGINATION * (page - 1)).Find(&promos)

	return promos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
