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
	db.Exec("DROP TABLE user_badges")
	db.Exec("DROP TABLE game_users")
	db.Exec("DROP TABLE game_developers")
	db.Exec("DROP TABLE game_genres")
	db.Exec("DROP TABLE game_reviews")
	db.Exec("DROP TABLE game_tags")

	UserMigrate(db)
	GameMigrate(db)
	TransactionMigrate(db)

	fmt.Println("[INFO] MIGRATED.")
}

func UserMigrate(db *gorm.DB) {
	// DROPPING ALL TABLE
	db.Migrator().DropTable(
		&model.Admin{},
		&model.UserReport{},
		&model.FriendRequest{},
		&model.UserComment{},
		&model.User{},
		&model.Country{},
		&model.Otp{},
		&model.SuspensionRequest{},
		&model.Badge{},
	)

	// CREATING ALL TABLE
	db.AutoMigrate(
		&model.Badge{},
		&model.Admin{},
		&model.SuspensionRequest{},
		&model.Country{},
		&model.User{},
		&model.UserComment{},
		&model.FriendRequest{},
		&model.UserReport{},
		&model.Otp{},
	)

}

func GameMigrate(db *gorm.DB) {
	db.Migrator().DropTable(
		&model.GameVideo{},
		&model.GameImage{},
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
		&model.Developer{},
		&model.GameImage{},
		&model.GameVideo{},
		&model.Game{},
	)
}

func TransactionMigrate(db *gorm.DB) {
	db.Migrator().DropTable(
		&model.Cart{},
	)

	db.AutoMigrate(
		&model.Cart{},
	)
}

