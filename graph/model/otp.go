package model

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
	"math/rand"
	"time"
)

type Otp struct {
	ID          int    		`json:"id"`
	Code        string    	`json:"code"`
	Email       string    	`json:"email"`
	CountryId 	int    		`json:"countryId"`
	ValidUntil  time.Time 	`json:"validUntil"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}

type NewOtp struct {
	Email       string 	`json:"email"`
	CountryId 	int 	`json:"countryId"`
}


func SendEmail(otp Otp) {
	mailjetClient := mailjet.NewMailjetClient("afd1d885ca87fb244abd14ddb1fe0f97", "368a25b4ce7ccbc881316cb49f7623f3")
	messagesInfo := []mailjet.InfoMessagesV31 {
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "foobarbaz234@gmail.com",
				Name: "Staem Support",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31 {
					Email: otp.Email,
					Name: "Nico",
				},
			},
			Subject: "Greetings from Staem.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h2>Hello, welcome to Staem</h2>" +
				"<br><h3>To continue creating your new Staem account, please insert this One Time Password (OTP) below in your registration form. </h3>" +
				"<br><h1><b>" + otp.Code + "</b></h1>" +
				"<br>OTP valid until " + otp.ValidUntil.Format(time.RFC850),
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


func GenerateRandomCode() string{
	var (
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + numberSet
	)
	var password string

	for i := 0; i < 5; i++ {
		idx := rand.Intn(len(allCharSet))
		password += string(allCharSet[idx])
	}

	return password
}