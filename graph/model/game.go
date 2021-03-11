package model

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"log"
	"strconv"
	"time"
)

type Game struct {
	ID            int          `json:"id" gorm:"primaryKey"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	ReleaseDate   time.Time    `json:"releaseDate"`
	Genres        []*Genre     `json:"genres" gorm:"many2many:game_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags          []*Tag       `json:"tags" gorm:"many2many:game_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OriginalPrice float64      `json:"originalPrice"`
	PromoID       int          `json:"promoId"`
	Promo         *Promo       `json:"promo" gorm:"foreignKey:PromoID;"`
	GamePlayHour  float64      `json:"gamePlayHour"`
	GameReviews   []*Review    `json:"gameReviews" gorm:"many2many:game_reviews;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Developers    []*Developer `json:"developers" gorm:"many2many:game_developers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PublisherID   int          `json:"publisherId;"`
	Publisher     *Publisher   `json:"publisher" gorm:"foreignKey:PublisherID;"`
	SystemID      int          `json:"systemId"`
	System        *System      `json:"system" gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Users         []*User      `json:"users" gorm:"many2many:game_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Banner string       `json:"banner"`
	Video  []*GameVideo `json:"video" gorm:"foreignKey:GameID"`
	Images []*GameImage `json:"images" gorm:"foreignKey:GameID"`

	Discussions   []*GameDiscussion `json:"discussions" gorm:"foreignKey:GameID"`
	Items         []*Item           `json:"items" gorm:"foreignKey:GameID"`


	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type NewGame struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Genres        []*InputGenre     `json:"genres"`
	Tags          []*InputTag       `json:"tags"`
	OriginalPrice float64           `json:"originalPrice"`
	Developers    []*InputDeveloper `json:"developers"`
	PublisherID   int               `json:"publisherId"`
	SystemID      int               `json:"systemId"`
}

func SendEmailPromo(game Game, user User) {
	finalPrice := int(game.OriginalPrice) - int(game.OriginalPrice) * game.Promo.DiscountPercentage/100
	mailjetClient := mailjet.NewMailjetClient("afd1d885ca87fb244abd14ddb1fe0f97", "368a25b4ce7ccbc881316cb49f7623f3")
	messagesInfo := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From: &mailjet.RecipientV31{
				Email: "foobarbaz234@gmail.com",
				Name:  "Staem Marketing",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: user.Email,
					Name:  "Nico",
				},
			},
			Subject:  "Greetings from Staem.",
			TextPart: "My first Mailjet email",
			HTMLPart: "<h2>One of your wishlist item is on sale!</h2>" +
				"<br><img src=\"" + game.Banner +"\">" +
				"<br><h3>" + string(game.Name) + " is now on " + strconv.Itoa(game.Promo.DiscountPercentage) + "% sale at Rp." + strconv.Itoa(finalPrice) + "</h3>",
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
