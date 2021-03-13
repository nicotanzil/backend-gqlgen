package model

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type ItemTransaction struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	ItemID    int        `json:"itemId"`
	Item      *Item      `json:"item" gorm:"foreignKey:ItemID"`
	SellerID  int        `json:"sellerId"`
	Seller    *User      `json:"seller" gorm:"foreignKey:SellerID"`
	BuyerID   int        `json:"buyerId"`
	Buyer     *User      `json:"buyer" gor:"foreignKey:BuyerID"`
	Price     int        `json:"price"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt gorm.DeletedAt
}

func SendItemPurchaseReceipt(transaction ItemTransaction) {
	fmt.Println("Email: ", transaction.Buyer.Email)
	mailjetClient := mailjet.NewMailjetClient("afd1d885ca87fb244abd14ddb1fe0f97", "368a25b4ce7ccbc881316cb49f7623f3")
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "foobarbaz234@gmail.com",
				Name: "Staem Financial",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: transaction.Buyer.Email,
					Name: "Nico",
				},
			},
			Subject: "Greetings from Staem.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h2>Thank you for your purchase at Staem!</h2>" +
				"<br><h3>Purchased Item: Item(" + strconv.Itoa(transaction.Item.ID) + ")</h3>" +
				"<br><h1><b>Total price: " + strconv.Itoa(transaction.Price) + "</b></h1>" +
				"<br>Tax (10%): Rp. " + strconv.Itoa(transaction.Price / 10),
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}

func SendItemSellEmail(transaction ItemTransaction) {
	mailjetClient := mailjet.NewMailjetClient("afd1d885ca87fb244abd14ddb1fe0f97", "368a25b4ce7ccbc881316cb49f7623f3")
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "foobarbaz234@gmail.com",
				Name: "Staem Financial",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: transaction.Seller.Email,
					Name: "Nico",
				},
			},
			Subject: "Greetings from Staem.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h2>Thank you for selling your item at Staem!</h2>" +
				"<br><h3>Sold Item: Item(" + strconv.Itoa(transaction.Item.ID) + ")</h3>" +
				"<br><h1><b>Total price: " + strconv.Itoa(transaction.Price) + "</b></h1>" +
				"<br>Tax (10%): Rp. " + strconv.Itoa(transaction.Price / 10),
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo }
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
}