package migration

import (
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/database"
)

func Migrate() {

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] MIGRATING...")
	// MIGRATE ALL TABLE
	db.Exec("DROP TABLE game_developers")
	db.Exec("DROP TABLE game_genres")
	db.Exec("DROP TABLE game_reviews")
	db.Exec("DROP TABLE game_tags")

	fmt.Println("[INFO] MIGRATED.")
}

//func UserMigrate(db *gorm.DB) {
//	// DROPPING ALL TABLE
//	db.Migrator().DropTable(
//		&model.User{},
//		&model.Country{},
//		&model.Otp{},
//	)
//
//	// CREATING ALL TABLE
//	db.AutoMigrate(
//		&model.User{},
//		&model.Country{},
//		&model.Otp{},
//		)
//
//}
//
//func GameMigrate(db *gorm.DB) {
//
//	// DROPPING ALL TABLE
//	db.Exec("DROP TABLE game_developers")
//	db.Exec("DROP TABLE game_genres")
//	db.Exec("DROP TABLE game_reviews")
//	db.Exec("DROP TABLE game_tags")
//	db.Migrator().DropTable(
//		&model.Game{},
//		&model.Developer{},
//		&model.Genre{},
//		&model.Publisher{},
//		&model.ReviewVote{},
//		&model.Review{},
//		&model.System{},
//		&model.Tag{},
//	)
//
//	// CREATING ALL TABLE
//	db.AutoMigrate(
//		&model.Developer{},
//		&model.Game{},
//		&model.Genre{},
//		&model.Publisher{},
//		&model.ReviewVote{},
//		&model.Review{},
//		&model.System{},
//		&model.Tag{},
//	)
//}

