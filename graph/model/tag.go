package model

import (
	"time"
)

type Tag struct {
	ID   		int 		`json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}



//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&Tag{})
//	db.AutoMigrate(&Tag{})
//
//	SeedTags()
//}
//
//func SeedTags() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	tags := []Tag {
//		{
//			Name: "MOBA",
//		},
//		{
//			Name: "Action",
//		},
//	}
//
//	for _, tag := range tags {
//		db.Create(&tag)
//	}
//}
