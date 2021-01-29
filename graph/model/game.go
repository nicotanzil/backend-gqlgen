package model

import (
	"time"
)

type Game struct {
	ID                 	int       		`json:"id" gorm:"primaryKey"`
	Name               	string       	`json:"name"`
	Description        	string       	`json:"description"`
	ReleaseDate        	time.Time    	`json:"releaseDate"`
	Genres             	[]*Genre    	`json:"genres" gorm:"many2many:game_genres;"`
	Tags               	[]*Tag    		`json:"tags" gorm:"many2many:game_tags;"`
	OriginalPrice      	float64      	`json:"originalPrice"`
	OnSale             	bool         	`json:"onSale"`
	DiscountPercentage 	int          	`json:"discountPercentage"`
	GamePlayHour       	float64      	`json:"gamePlayHour"`
	GameReviews        	[]*Review    	`json:"gameReviews" gorm:"many2many:game_reviews;"`
	Developers         	[]*Developer 	`json:"developers" gorm:"many2many:game_developers;"`
	PublisherID         int				`json:"publisher"`
	SystemID			int				`json:"system"`
	CreatedAt          	time.Time    	`json:"createdAt"`
	UpdatedAt          	time.Time    	`json:"updatedAt"`
	DeletedAt          	*time.Time   	`json:"deletedAt"`
}

//func init() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&Game{})
//	db.AutoMigrate(&Game{})
//
//	SeedGames()
//}
//
//
//func SeedGames() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	games := []Game {
//		{
//			Name: "Dota 2",
//			Description: "Every day, millions of players worldwide enter battle as one of over a hundred Dota heroes. And no matter if it's their 10th hour of play or 1,000th, there's always something new to discover. With regular updates that ensure a constant evolution of gameplay, features, and heroes, Dota 2 has truly taken on a life of its own.\n",
//			ReleaseDate: time.Now(),
//			Genres: []*Genre{
//				{
//					ID: 1,
//				},
//				{
//					ID: 2,
//				},
//				{
//					ID: 5,
//				},
//			},
//			Tags: []*Tag{
//				{
//					ID: 1,
//				},
//				{
//					ID: 2,
//				},
//			},
//			OriginalPrice: 0,
//			OnSale: false,
//			DiscountPercentage: 0,
//			GamePlayHour: 1500,
//		},
//	}
//
//	for _, game := range games {
//		db.Create(&game)
//		db.Save(&game)
//	}
//}

