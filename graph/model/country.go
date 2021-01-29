package model

import (
	"time"
)

type Country struct {
	ID   int 				`json:"id"`
	Name string 			`json:"name"`
	CreatedAt time.Time 	`json:"createdAt"`
	UpdatedAt time.Time 	`json:"updatedAt"`
	DeletedAt *time.Time 	`json:"deletedAt"`
}

//func init() {
//	//db, err := database.Connect()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//db.Migrator().DropTable(&Country{})
//	//db.AutoMigrate(&Country{})
//
//	SeedCountries()
//}
//
//func GetCountryIDByName(name string) int {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	var temp Country
//	db.Where("name = ?", name).First(&temp)
//
//	return temp.ID
//}
//
//func SeedCountries(){
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	countries := []Country{
//		{
//			Name: "Indonesia",
//		},
//		{
//			Name: "Singapore",
//		},
//		{
//			Name: "Japan",
//		},
//	}
//
//	for _, country := range countries {
//		db.Create(&country)
//	}
//}