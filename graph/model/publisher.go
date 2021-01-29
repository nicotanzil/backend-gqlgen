package model

import (
	"time"
)

type Publisher struct {
	ID   		int `json:"id"`
	Name 		string `json:"name"`
	CreatedAt 	time.Time `json:"createdAt"`
	UpdatedAt 	time.Time `json:"updatedAt"`
	DeletedAt 	*time.Time `json:"deletedAt"`
}



//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&Publisher{})
//	db.AutoMigrate(&Publisher{})
//
//	SeedPublishers()
//}
//
//func SeedPublishers() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	publishers := []Publisher {
//		{
//			Name: "Valve",
//		},
//	}
//
//	for _, publisher := range publishers {
//		db.Create(&publisher)
//	}
//}