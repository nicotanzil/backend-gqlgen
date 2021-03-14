package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateTransaction(ctx context.Context, transaction *model.InputTransactionHeader) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var newTransaction model.TransactionHeader

	newTransaction = model.TransactionHeader{
		PaymentType: &model.PaymentType{ID: transaction.PaymentTypeID},
		Sender:      &model.User{ID: transaction.SenderID},
		Receiver:    &model.User{ID: transaction.ReceiverID},
		Total:       transaction.Total,
	}

	db.Create(&newTransaction)

	var latest model.TransactionHeader
	db.Last(&latest)

	var details []*model.TransactionDetail

	for i := 0; i < len(transaction.TransactionDetails); i++ {
		detail := model.TransactionDetail{
			TransactionHeader: &model.TransactionHeader{ID: latest.ID},
			Game:              &model.Game{ID: transaction.TransactionDetails[i]},
		}
		details = append(details, &detail)
		db.Create(&detail)
	}

	latest.TransactionDetails = details
	db.Save(&latest)

	// Clear cart
	db.Where("user_id = ?", transaction.SenderID).Delete(&model.Cart{})

	// Deduct wallet balance
	var sender model.User
	db.Where("id = ?", transaction.SenderID).
		First(&sender)

	prev := sender.Balance
	prev -= float64(latest.Total)
	sender.Balance = prev

	// Add points shop balance to user
	point := latest.Total / 15000
	sender.Points += point * 100
	db.Save(&sender)

	// Set gift to complete
	var completedGifts model.Gift
	db.Where("sender_id = ?", transaction.SenderID).
		Where("is_complete = false").
		First(&completedGifts)

	completedGifts.IsComplete = true
	db.Save(&completedGifts)

	// Send Game to user
	var receiver model.User
	db.Where("id = ?", transaction.ReceiverID).
		First(&receiver)

	var addGames []*model.Game
	for i := 0; i < len(transaction.TransactionDetails); i++ {
		var addGame model.Game
		db.First(&addGame, transaction.TransactionDetails[i])
		addGames = append(addGames, &addGame)
	}

	receiver.Games = addGames
	db.Save(&receiver)

	// Send Purchase receipt email
	var latest2 model.TransactionHeader
	db.Preload(clause.Associations).Last(&latest2)

	SendPurchaseReceipt(latest2)

	return true, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func SendPurchaseReceipt(transaction model.TransactionHeader) {
	mailjetClient := mailjet.NewMailjetClient("afd1d885ca87fb244abd14ddb1fe0f97", "368a25b4ce7ccbc881316cb49f7623f3")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "foobarbaz234@gmail.com",
				Name:  "Staem Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: transaction.Sender.Email,
					Name:  "Nico",
				},
			},
			Subject:  "Greetings from Staem.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<div style=\"padding: 10px 20px\">\n  " +
				"<div>Hello " + transaction.Sender.ProfileName + "</div>\n  " +
				"<div>Thank you for your recent transaction on Staem.</div>\n  " +
				"<div>Account name: " + transaction.Sender.AccountName + "</div>\n  " +
				"<div>Payment method: " + transaction.PaymentType.Name + "</div>\n  " +
				"<div>Transaction ID: " + strconv.Itoa(transaction.ID) + "</div>\n  " +
				"<div>Your total for this transaction: Rp. " + strconv.Itoa(transaction.Total) + "</div>\n</div>\n",
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}
