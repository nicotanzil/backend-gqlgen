package model

import (
	"github.com/nicotanzil/backend-gqlgen/database"
	"time"
)

type Country struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
	DeletedAt int    `json:"deletedAt"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&Country{})
	db.AutoMigrate(&Country{})

	SeedCountries()
}

func SeedCountries(){
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	countries := []Country{
		{
			Name: "Indonesia",

			CreatedAt: int(time.Now().Unix()),
			UpdatedAt: int(time.Now().Unix()),
			DeletedAt: nil,
		},
		{
			Name: "Singapore",

			CreatedAt: int(time.Now().Unix()),
			UpdatedAt: int(time.Now().Unix()),
			DeletedAt: nil,
		},
		{
			Name: "Japan",

			CreatedAt: int(time.Now().Unix()),
			UpdatedAt: int(time.Now().Unix()),
			DeletedAt: nil,
		},
	}

	for _, country := range countries {
		db.Create(&country)
	}
}