package database

import (
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm"
	"time"
)

func Seed() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] SEEDING...")
	// SEED ALL NECESSARY TABLE

	SeedTag(db)
	SeedSystem(db)
	SeedReview(db)
	SeedReviewVote(db)
	SeedPublisher(db)
	SeedGenre(db)
	SeedGame(db)
	SeedDeveloper(db)
	SeedCountry(db)
	SeedUser(db)

	fmt.Println("[INFO] SEEDED.")
}

func SeedTag(db *gorm.DB) {
	tags := []model.Tag {
		{
			Name: "MOBA",
		},
		{
			Name: "Action",
		},
	}

	for _, tag := range tags {
		db.Create(&tag)
	}
}

func SeedSystem(db *gorm.DB) {
	systems := []model.System{
		{
			Os:        "Windows 7 or newer",
			Memory:    4,
			Graphics:  "nVidia GeForce 8600/9600GT, ATI/AMD Radeon HD2600/3600",
			Storage:   15,
		},
	}

	for _, system := range systems {
		db.Create(&system)
	}
}

func SeedReview(db *gorm.DB) {

}

func SeedReviewVote(db *gorm.DB) {

}

func SeedPublisher(db *gorm.DB) {
	publishers := []model.Publisher {
		{
			Name: "Valve",
		},
	}

	for _, publisher := range publishers {
		db.Create(&publisher)
	}
}

func SeedGenre(db *gorm.DB) {
	genres := []model.Genre {
		{
			Name: "Free to Play",
		},
		{
			Name: "Early Access",
		},
		{
			Name: "Action",
		},
		{
			Name: "Adventure",
		},
		{
			Name: "Casual",
		},
		{
			Name: "Indie",
		},
		{
			Name: "Massively Multiplayer",
		},
		{
			Name: "Racing",
		},
		{
			Name: "RPG",
		},
		{
			Name: "Simulation",
		},
		{
			Name: "Sports",
		},
		{
			Name: "Strategy",
		},
	}

	for _, genre := range genres {
		db.Create(&genre)
		db.Save(&genre)
	}
}

func SeedGame(db *gorm.DB) {
	games := []model.Game {
		{
			Name: "Dota 2",
			Description: "Every day, millions of players worldwide enter battle as one of over a hundred Dota heroes. And no matter if it's their 10th hour of play or 1,000th, there's always something new to discover. With regular updates that ensure a constant evolution of gameplay, features, and heroes, Dota 2 has truly taken on a life of its own.\n",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 12,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			OriginalPrice: 0,
			OnSale: false,
			DiscountPercentage: 0,
			GamePlayHour: 1500,
			PublisherID: 1,
			SystemID: 1,
		},
	}

	for _, game := range games {
		db.Create(&game)
		db.Save(&game)
	}
}

func SeedDeveloper(db *gorm.DB) {
	developers := []model.Developer {
		{
			Name: "Valve",
		},
	}

	for _, developer := range developers {
		db.Create(&developer)
	}
}

func SeedCountry(db *gorm.DB) {
	countries := []model.Country{
		{
			Name: "Indonesia",
		},
		{
			Name: "Singapore",
		},
		{
			Name: "Japan",
		},
	}

	for _, country := range countries {
		db.Create(&country)
	}
}

func SeedUser(db *gorm.DB) {
	users := []model.User{
		{
			AccountName: "nico",
			ProfileName: "nico tanzil",
			RealName: "Nico Tanzil",
			Email: "nico@mail.com",
			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
			Balance: 100,
			CustomURL: "nico",
			CountryID: 1,
		},
		{
			AccountName: "William",
			ProfileName: "William",
			RealName: "William Martin",
			Email: "will@mail.com",
			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
			Balance: 100,
			CustomURL: "william",
			CountryID: 2,
		},
		{
			AccountName: "Ricko",
			ProfileName: "Ricko",
			RealName: "Ricko Adrio",
			Email: "rick@mail.com",
			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
			Balance: 100,
			CustomURL: "rick",
			CountryID: 3,
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}