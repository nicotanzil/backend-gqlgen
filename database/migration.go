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
	dbClose, _ := db.DB()
	defer dbClose.Close()

	fmt.Println("[INFO] MIGRATING...")
	// MIGRATE ALL TABLE
	db.Exec("DROP TABLE user_friends")
	db.Exec("DROP TABLE user_badges")
	db.Exec("DROP TABLE user_avatar_frames")
	db.Exec("DROP TABLE user_profile_backgrounds")
	db.Exec("DROP TABLE user_mini_profile_backgrounds")

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

		&model.CommunityArtPostReview{},
		&model.CommunityArtPost{},
		&model.CommunityGameReviewComment{},
		&model.CommunityGameReview{},
		&model.GameDiscussionReply{},
		&model.GameDiscussion{},

		&model.Gift{},
		&model.TransactionHeader{},
		&model.TransactionDetail{},
		&model.PaymentType{},

		&model.Item{},
		&model.ItemType{},

		&model.UserReport{},
		&model.FriendRequest{},
		&model.UserComment{},
		&model.User{},
		&model.Country{},
		&model.Otp{},
		&model.SuspensionRequest{},
		&model.Badge{},
		&model.AvatarFrame{},
		&model.ProfileBackground{},
		&model.MiniProfileBackground{},
		&model.Theme{},
	)

	// CREATING ALL TABLE
	db.AutoMigrate(
		&model.Badge{},
		&model.AvatarFrame{},
		&model.ProfileBackground{},
		&model.MiniProfileBackground{},
		&model.Theme{},
		&model.Admin{},
		&model.SuspensionRequest{},
		&model.Country{},
		&model.User{},
		&model.UserComment{},
		&model.FriendRequest{},
		&model.UserReport{},

		&model.ItemType{},
		&model.Item{},

		&model.PaymentType{},
		&model.TransactionDetail{},
		&model.TransactionHeader{},
		&model.Gift{},

		&model.CommunityArtPost{},
		&model.CommunityArtPostReview{},
		&model.CommunityGameReview{},
		&model.CommunityGameReviewComment{},
		&model.GameDiscussion{},
		&model.GameDiscussionReply{},

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
		&model.Wishlist{},
		&model.ItemTransaction{},
	)

	db.AutoMigrate(
		&model.Cart{},
		&model.Wishlist{},
		&model.ItemTransaction{},
	)
}

