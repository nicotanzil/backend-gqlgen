package model

import (
	"time"
)

type System struct {
	ID        int    	`json:"id"`
	Os        string    `json:"os"`
	Memory    int       `json:"memory"`
	Graphics  string    `json:"graphics"`
	Storage   int       `json:"storage"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}



//func init() {
//
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	db.Migrator().DropTable(&System{})
//	db.AutoMigrate(&System{})
//
//	SeedSystem()
//}
//
//func SeedSystem() {
//	db, err := database.Connect()
//	if err != nil {
//		panic(err)
//	}
//
//	systems := []System{
//		{
//			Os:        "Windows 7 or newer",
//			Memory:    4,
//			Graphics:  "nVidia GeForce 8600/9600GT, ATI/AMD Radeon HD2600/3600",
//			Storage:   15,
//		},
//	}
//
//	for _, system := range systems {
//		db.Create(&system)
//	}
//}