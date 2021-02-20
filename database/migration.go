package database

import (
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm"
)

func Migrate() {

	db, err := Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] MIGRATING...")
	// MIGRATE ALL TABLE
	db.Exec("DROP TABLE user_friends")
	db.Exec("DROP TABLE game_users")
	db.Exec("DROP TABLE game_developers")
	db.Exec("DROP TABLE game_genres")
	db.Exec("DROP TABLE game_reviews")
	db.Exec("DROP TABLE game_tags")

	UserMigrate(db)
	GameMigrate(db)

	fmt.Println("[INFO] MIGRATED.")
}

func UserMigrate(db *gorm.DB) {
	// DROPPING ALL TABLE
	db.Migrator().DropTable(
		&model.Admin{},
		&model.FriendRequest{},
		&model.UserComment{},
		&model.User{},
		&model.Country{},
		&model.Otp{},
	)

	// CREATING ALL TABLE
	db.AutoMigrate(
		&model.Admin{},
		&model.Country{},
		&model.User{},
		&model.UserComment{},
		&model.FriendRequest{},
		&model.Otp{},
	)

}

func GameMigrate(db *gorm.DB) {
	db.Migrator().DropTable(
		&model.Game{},
		&model.Promo{},
		&model.Developer{},
		&model.Genre{},
		&model.Publisher{},
		&model.ReviewVote{},
		&model.Review{},
		&model.System{},
		&model.Tag{},
	)

	// CREATING ALL TABLE
	db.AutoMigrate(
		&model.Tag{},
		&model.System{},
		&model.Review{},
		&model.ReviewVote{},
		&model.Publisher{},
		&model.Genre{},
		&model.Promo{},
		&model.Game{},
		&model.Developer{},

	)
}

