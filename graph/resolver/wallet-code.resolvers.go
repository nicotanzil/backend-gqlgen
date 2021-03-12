package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *mutationResolver) RedeemCode(ctx context.Context, code string, userID int) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	var walletCode model.WalletCode

	db.First(&walletCode, "code = ?", code)

	if walletCode.ID != 0 {
		// Code exists
		balance := walletCode.Balance

		var user model.User

		db.First(&user, "id = ?", userID)
		prev := user.Balance

		prev = prev + float64(balance)

		user.Balance = prev
		db.Save(&user)

		db.Delete(model.WalletCode{}, "code = ?", code)

		return true, nil

	}
	return false, nil
}
