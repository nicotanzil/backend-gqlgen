package model

import (
	"time"
)

type Genre struct {
	ID          int      	`json:"id" gorm:"primaryKey"`
	Name        string    	`json:"name"`
	Description string    	`json:"description"`
	Games       []*Game   	`json:"games" gorm:"many2many:game_genres;"`
	CreatedAt   time.Time 	`json:"createdAt"`
	UpdatedAt   time.Time 	`json:"updatedAt"`
	DeletedAt   *time.Time 	`json:"deletedAt"`
}

//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&Genre{})
//	db.AutoMigrate(&Genre{})
//
//	SeedGenres()
//}
//
//func SeedGenres() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	genres := []Genre {
//		{
//			Name: "Free to Play",
//		},
//		{
//			Name: "Early Access",
//		},
//		{
//			Name: "Action",
//		},
//		{
//			Name: "Adventure",
//		},
//		{
//			Name: "Casual",
//		},
//		{
//			Name: "Indie",
//		},
//		{
//			Name: "Massively Multiplayer",
//		},
//		{
//			Name: "Racing",
//		},
//		{
//			Name: "RPG",
//		},
//		{
//			Name: "Simulation",
//		},
//		{
//			Name: "Sports",
//		},
//		{
//			Name: "Strategy",
//		},
//	}
//
//	for _, genre := range genres {
//		db.Create(&genre)
//		db.Save(&genre)
//	}
//}
