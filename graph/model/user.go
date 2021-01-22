package model

import (
	"github.com/nicotanzil/backend-gqlgen/database"
	"time"
)

type User struct {
	ID             string   `json:"id"`
	ProfileName    string   `json:"profileName"`
	RealName       string   `json:"realName"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	Balance        float64  `json:"balance"`
	CustomURL      string   `json:"customURL"`
	ProfilePicture string   `json:"profilePicture"`
	CountryID      string   `json:"countryId"`
	Country        *Country `json:"country"`
	CreatedAt      int      `json:"createdAt"`
	UpdatedAt      int      `json:"updatedAt"`
	DeletedAt      int      `json:"deletedAt"`
}

func init() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})

	SeedUsers()
}

func SeedUsers() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var country Country
	db.First(&country, 2)

	users := []User{
		{
			ProfileName: "UserProfileName_1",
			RealName: "UserRealName_1",
			Email: "user1@mail.com",
			Password: "password",
			Balance: 100,
			CustomURL: "user1",
			CountryID: "1",
			CreatedAt: int(time.Now().Unix()),
			UpdatedAt: int(time.Now().Unix()),
			DeletedAt: nil,
		},
		{
			ProfileName: "UserProfileName_2",
			RealName: "UserRealName_2",
			Email: "user2@mail.com",
			Password: "password",
			Balance: 100,
			CustomURL: "user2",
			CountryID: "2",
			CreatedAt: int(time.Now().Unix()),
			UpdatedAt: int(time.Now().Unix()),
			DeletedAt: nil,
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}