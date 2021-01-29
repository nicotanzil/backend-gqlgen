package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) CreateOtp(ctx context.Context, input model.NewOtp) (*model.Otp, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var temp model.Otp

	db.Where("email = ?", input.Email).First(&temp)

	if temp.ID != 0 {
		fmt.Println("OTP code already sent")
		return &temp, nil
	}

	otp := model.Otp{
		Code:       model.GenerateRandomCode(),
		Email:      input.Email,
		CountryId:  input.CountryId,
		ValidUntil: time.Now().Add(time.Minute * 5), //OTP valid for 5 minutes
	}

	db.Create(&otp)

	model.SendEmail(otp) //Send OTP email

	return &otp, nil
}

func (r *mutationResolver) UpdateOtp(ctx context.Context, code string) (*model.Otp, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Otps(ctx context.Context) ([]*model.Otp, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var otps []*model.Otp

	db.Find(&otps)

	return otps, nil
}

func (r *queryResolver) GetOtpByCode(ctx context.Context, code *string) (*model.Otp, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var otp model.Otp

	db.Where("code = ? AND is_valid = ?", code, true).First(&otp)

	//Delete OTP

	db.Save(&otp)

	return &otp, nil
}
