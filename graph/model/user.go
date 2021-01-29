package model

import (
	"time"
)

type User struct {
	ID             int   		`json:"id"`
	AccountName    string   	`json:"accountName"`
	ProfileName    string   	`json:"profileName"`
	RealName       string   	`json:"realName"`
	Email          string   	`json:"email"`
	Password       string   	`json:"password"`
	Balance        float64  	`json:"balance"`
	CustomURL      string   	`json:"customURL"`
	ProfilePicture string   	`json:"profilePicture"`
	CountryID      int 			`json:"country"`
	CreatedAt 	   time.Time 	`json:"createdAt"`
	UpdatedAt 	   time.Time 	`json:"updatedAt"`
	DeletedAt 	   *time.Time 	`json:"deletedAt"`
}



//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&User{})
//	db.AutoMigrate(&User{})
//
//	SeedUsers()
//}
//
//
//func SeedUsers() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//
//	users := []User{
//		{
//			AccountName: "nico",
//			ProfileName: "nico tanzil",
//			RealName: "Nico Tanzil",
//			Email: "nico@mail.com",
//			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
//			Balance: 100,
//			CustomURL: "nico",
//			CountryID: 1,
//		},
//		{
//			AccountName: "William",
//			ProfileName: "William",
//			RealName: "William Martin",
//			Email: "will@mail.com",
//			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
//			Balance: 100,
//			CustomURL: "william",
//			CountryID: 2,
//		},
//		{
//			AccountName: "Ricko",
//			ProfileName: "Ricko",
//			RealName: "Ricko Adrio",
//			Email: "rick@mail.com",
//			Password: "$2a$14$LHJO1WVSn/HK7lEtm164aOoVoGv5yTJpXorfttKNE0HniMU.CqPlK",
//			Balance: 100,
//			CustomURL: "rick",
//			CountryID: 3,
//		},
//	}
//
//	for _, user := range users {
//		db.Create(&user)
//	}
//}
