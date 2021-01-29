package model

import (
	"time"
)

type Developer struct {
	ID   		int 		`json:"id"`
	Name 		string 		`json:"name"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	DeletedAt 	*time.Time 	`json:"deletedAt"`
}

type NewDeveloper struct {
	Name string `json:"name"`
}

//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&Developer{})
//	db.AutoMigrate(&Developer{})
//
//	SeedDevelopers()
//}
//
//func SeedDevelopers() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	developers := []Developer {
//		{
//			Name: "Valve",
//		},
//	}
//
//	for _, developer := range developers {
//		db.Create(&developer)
//	}
//}