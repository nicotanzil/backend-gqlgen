package database

import (
	"fmt"
	"github.com/nicotanzil/backend-gqlgen/app/firebase-data"
	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm"
	"time"
)

func Seed() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] SEEDING...")
	// SEED ALL NECESSARY TABLE

	SeedAdmin(db)
	SeedSuspensionRequest(db)
	SeedTag(db)
	SeedSystem(db)
	SeedReview(db)
	SeedReviewVote(db)
	SeedPublisher(db)
	SeedGenre(db)
	SeedDeveloper(db)
	SeedPromo(db)
	SeedGame(db)
	SeedGameImage(db)
	SeedGameVideo(db)
	SeedCountry(db)
	SeedUser(db)
	SeedUserReport(db)

	fmt.Println("[INFO] SEEDED.")
}

func SeedAdmin(db *gorm.DB) {
	admin := model.Admin{
		ID:          1,
		AccountName: "admin",
		Password:    providers.HashPassword("admin"),
	}

	db.Create(&admin)
}

func SeedGameImage(db *gorm.DB) {
	gameImages := []model.GameImage{
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2F1.jpg?alt=media&token=738cf71e-7978-482b-9b53-383a7518b7db"},
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2F2.jpg?alt=media&token=4303c33d-263f-4747-b92b-05627279be46"},
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2F3.jpg?alt=media&token=ff544eab-441f-4a63-8684-6440bb9d2206"},
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2F4.jpg?alt=media&token=36293e0b-cc7f-4898-a7bd-911ecba302a0"},
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2F5.jpg?alt=media&token=458390ea-5331-4628-ad98-7cc7f69f65cb"},
		{GameID: 2, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2F1.jpg?alt=media&token=7f5e64b4-9853-4df1-9fd4-143ee83c3578"},
		{GameID: 2, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2F2.jpg?alt=media&token=54d52e5f-73f1-46e0-879f-e3effbc901ad"},
		{GameID: 2, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2F3.jpg?alt=media&token=23254fc0-9328-4f9a-938b-d59b669712c3"},
		{GameID: 2, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2F4.jpg?alt=media&token=36086373-200b-4e71-89f2-e6a53a880fd0"},
		{GameID: 2, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2F5.jpg?alt=media&token=36b46e74-1ff6-4411-8da3-74526824ffc2"},
		{GameID: 3, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2F1.jpg?alt=media&token=61fff14c-aa9e-40e8-bee7-ff34a405bc18"},
		{GameID: 3, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2F2.jpg?alt=media&token=1be0a5d7-5adf-4a57-9e42-cfe1261f2157"},
		{GameID: 3, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2F3.jpg?alt=media&token=6fa25ff5-6998-4613-a7d0-39391288a758"},
		{GameID: 3, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2F4.jpg?alt=media&token=c6bb3113-ee28-461f-bc30-182a3036fe3d"},
		{GameID: 3, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2F5.jpg?alt=media&token=a15181b9-1d11-4861-bcf0-af3eed554bac"},
		{GameID: 4, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2F1.jpg?alt=media&token=ad074c45-0e4f-4894-82d4-2b953125d339"},
		{GameID: 4, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2F2.jpg?alt=media&token=270ae5ef-336f-4eb8-b311-51ec901bd87f"},
		{GameID: 4, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2F3.jpg?alt=media&token=accb5312-3c9a-4751-9994-ae092b6ea9d0"},
		{GameID: 4, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2F4.jpg?alt=media&token=92d7eed4-4ffa-40bb-b9cc-7dc781d438a3"},
		{GameID: 4, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2F5.jpg?alt=media&token=0b38b234-52bc-4325-a64a-4eebc86ff31c"},
		{GameID: 5, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2F1.jpg?alt=media&token=e532cf5d-6721-4871-b39c-99aa1cc30235"},
		{GameID: 5, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2F2.jpg?alt=media&token=eb183ce6-2dcb-4ad7-8f79-b2d076efc84d"},
		{GameID: 5, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2F3.jpg?alt=media&token=4bcd2864-82da-4bb7-8e91-3a646fee5c69"},
		{GameID: 5, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2F4.jpg?alt=media&token=2e2b118a-f449-4093-bf14-16bb60b4b031"},
		{GameID: 5, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2F5.jpg?alt=media&token=89ec1556-d73c-45be-894b-6b4bacfa18c8"},
		{GameID: 6, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2F1.jpg?alt=media&token=21ff03bb-d533-43b1-80a8-f59b184ba5fb"},
		{GameID: 6, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2F2.jpg?alt=media&token=e532061f-2532-45b9-be7f-8d2c80a539a7"},
		{GameID: 6, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2F3.jpg?alt=media&token=a167e1c6-9f07-4013-8c36-4b02a02af48a"},
		{GameID: 6, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2F4.jpg?alt=media&token=474429c0-0d47-43bb-ae97-0f5dfccc11f1"},
		{GameID: 6, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2F5.jpg?alt=media&token=3058b24e-cd33-446e-b855-fce0fe150e95"},
		{GameID: 7, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2F1.jpg?alt=media&token=432ac1d1-ad60-43c1-89c8-7efa5257acce"},
		{GameID: 7, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2F2.jpg?alt=media&token=af083f6d-3ce0-4789-b90e-c85370cdc8bb"},
		{GameID: 7, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2F3.jpg?alt=media&token=7e955225-1fa9-457a-9e76-f3630ef24056"},
		{GameID: 7, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2F4.jpg?alt=media&token=c97a66d4-1f26-4a84-8169-72b5491a618a"},
		{GameID: 7, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2F5.jpg?alt=media&token=aba239b2-f826-4680-8485-076b849f305c"},
	}

	for _, gameImage := range gameImages {
		db.Create(&gameImage)
	}

}

func SeedGameVideo(db *gorm.DB) {
	gameVideos := []model.GameVideo {
		{ GameID:   1, 	Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fvideos%2F1.mp4?alt=media&token=915b8004-f5d1-4279-81fc-2c1f7e56f102",},
		{ GameID: 	1,	Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fvideos%2F2.mp4?alt=media&token=4cbc40d4-0f03-41e7-97f4-1c3154d34bbc",},
	}

	for _, gameVideo := range gameVideos {
		db.Create(&gameVideo)
	}
}

func SeedTag(db *gorm.DB) {
	tags := []model.Tag{
		{
			Name: "Horror",
		},
		{
			Name: "Adventure",
		},
		{
			Name: "Indie",
		},
		{
			Name: "MOBA",
		},
		{
			Name: "Action",
		},
		{
			Name: "Third Person",
		},
		{
			Name: "RPG",
		},
		{
			Name: "Fighting",
		},
		{
			Name: "Multiplayer",
		},
		{
			Name: "Arcade",
		},
		{
			Name: "FPS",
		},
		{
			Name: "Mature Content",
		},
	}

	for _, tag := range tags {
		db.Create(&tag)
	}
}

func SeedSystem(db *gorm.DB) {
	systems := []model.System{
		{
			Os:       "Windows 7 or newer",
			Memory:   4,
			Graphics: "nVidia GeForce 8600/9600GT, ATI/AMD Radeon HD2600/3600",
			Storage:  15,
		},
	}

	for _, system := range systems {
		db.Create(&system)
	}
}

func SeedReview(db *gorm.DB) {

}

func SeedReviewVote(db *gorm.DB) {

}

func SeedPublisher(db *gorm.DB) {
	publishers := []model.Publisher{
		{
			Name: "Valve",
		},
		{
			Name: "Electronic Arts",
		},
		{
			Name: "Capcom",
		},
		{
			Name: "Rockstar Games",
		},
		{
			Name: "Re-Logic",
		},
	}

	for _, publisher := range publishers {
		db.Create(&publisher)
	}
}

func SeedGenre(db *gorm.DB) {
	genres := []model.Genre{
		{
			Name: "Free to Play",
		},
		{
			Name: "Early Access",
		},
		{
			Name: "Action",
		},
		{
			Name: "Adventure",
		},
		{
			Name: "Casual",
		},
		{
			Name: "Indie",
		},
		{
			Name: "Massively Multiplayer",
		},
		{
			Name: "Racing",
		},
		{
			Name: "RPG",
		},
		{
			Name: "Simulation",
		},
		{
			Name: "Sports",
		},
		{
			Name: "Strategy",
		},
	}

	for _, genre := range genres {
		db.Create(&genre)
		db.Save(&genre)
	}
}

func SeedPromo(db *gorm.DB) {
	promos := []model.Promo{
		{
			DiscountPercentage: 0,
			ValidUntil: int(time.Now().Add(time.Hour * 24 * 30).Unix()),
		},
		{
			DiscountPercentage: 50,
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 3).Unix()),
		},
		{
			DiscountPercentage: 10,
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 7).Unix()),
		},
	}

	for _, promo := range promos {
		db.Create(&promo)
		db.Save(&promo)
	}
}

func SeedGame(db *gorm.DB) {
	games := []model.Game{
		{
			Name:        "Dota 2",
			Description: "Every day, millions of players worldwide enter battle as one of over a hundred Dota heroes. And no matter if it's their 10th hour of play or 1,000th, there's always something new to discover. With regular updates that ensure a constant evolution of gameplay, features, and heroes, Dota 2 has truly taken on a life of its own.\n",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 12,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
			OriginalPrice:      0,
			PromoID: 1,
			Developers: []*model.Developer{
				{
					ID: 1,
				},
			},
			GamePlayHour: 1500,
			PublisherID:  1,
			SystemID:     1,
			Banner:       "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fbanner.jpg?alt=media&token=6fe38ac5-d74e-45f9-b1f3-5052403edfec",
		},
		{
			Name:        "STAR WARS Jedi: Fallen Order",
			Description: "A galaxy-spanning adventure awaits in Star Wars Jedi: Fallen Order, a new third-person action-adventure title from Respawn Entertainment. This narratively driven, single-player game puts you in the role of a Jedi Padawan who narrowly escaped the purge of Order 66 following the events of Episode 3: Revenge of the Sith. On a quest to rebuild the Jedi Order, you must pick up the pieces of your shattered past to complete your training, develop new powerful Force abilities and master the art of the iconic lightsaber - all while staying one step ahead of the Empire and its deadly Inquisitors.",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 5,
				},
				{
					ID: 7,
				},
				{
					ID: 10,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 2,
				},
				{
					ID: 5,
				},
			},
			OriginalPrice:      562000,
			PromoID: 2,
			Developers: []*model.Developer{
				{
					ID: 2,
				},
			},
			PublisherID: 2,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2Fbanner.jpg?alt=media&token=f12004fa-98c7-4238-b209-7437fa761197",
		},
		{
			Name:        "Street Fighter V",
			Description: "Experience the intensity of head-to-head battle with Street Fighter® V! Choose from 16 iconic characters, each with their own personal story and unique training challenges, then battle against friends online or offline with a robust variety of match options.\n\nEarn Fight Money in Ranked Matches, play for fun in Casual Matches or invite friends into a Battle Lounge and see who comes out on top! PlayStation 4 and Steam players can also play against each other thanks to cross-play compatibility!\n",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 8,
				},
				{
					ID: 2,
				},
				{
					ID: 1,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 6,
				},
				{
					ID: 7,
				},
			},
			OriginalPrice:      650000,
			PromoID: 2,
			Developers: []*model.Developer{
				{
					ID: 3,
				},
			},
			PublisherID: 3,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2Fbanner.jpg?alt=media&token=073881f4-a93e-4a05-9bc9-9efbb1401108",
		},
		{
			Name:        "Counter-Strike: Global Offensive",
			Description: "Counter-Strike: Global Offensive (CS: GO) expands upon the team-based action gameplay that it pioneered when it was launched 19 years ago.\n\nCS: GO features new maps, characters, weapons, and game modes, and delivers updated versions of the classic CS content (de_dust2, etc.).",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 2,
				},
				{
					ID: 5,
				},
				{
					ID: 10,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 2,
				},
				{
					ID: 3,
				},
				{
					ID: 5,
				},
				{
					ID: 12,
				},
			},
			OriginalPrice:      0,
			PromoID: 1,
			Developers: []*model.Developer{
				{
					ID: 1,
				},
				{
					ID: 4,
				},
			},
			PublisherID: 1,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2Fbanner.jpg?alt=media&token=71bf9807-e661-47f4-b19e-5b12b8f8eb79",
		},
		{
			Name:        "Grand Theft Auto V",
			Description: "When a young street hustler, a retired bank robber and a terrifying psychopath find themselves entangled with some of the most frightening and deranged elements of the criminal underworld, the U.S. government and the entertainment industry, they must pull off a series of dangerous heists to survive in a ruthless city in which they can trust nobody, least of all each other.",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 1,
				},
				{
					ID: 6,
				},
				{
					ID: 8,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 1,
				},
				{
					ID: 4,
				},
				{
					ID: 8,
				},
				{
					ID: 12,
				},
			},
			OriginalPrice:      300000,
			PromoID: 1,
			Developers: []*model.Developer{
				{
					ID: 5,
				},
			},
			PublisherID: 4,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2Fbanner.jpg?alt=media&token=dafa20ed-2139-45cc-8a59-38e245faac53",
		},
		{
			Name:        "Terraria",
			Description: "Dig, Fight, Explore, Build: The very world is at your fingertips as you fight for survival, fortune, and glory. Will you delve deep into cavernous expanses in search of treasure and raw materials with which to craft ever-evolving gear, machinery, and aesthetics? Perhaps you will choose instead to seek out ever-greater foes to test your mettle in combat? Maybe you will decide to construct your own city to house the host of mysterious allies you may encounter along your travels? ",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 4,
				},
				{
					ID: 8,
				},
				{
					ID: 12,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 2,
				},
				{
					ID: 3,
				},
				{
					ID: 7,
				},
			},
			OriginalPrice:      89999,
			PromoID: 1,
			Developers: []*model.Developer{
				{
					ID: 6,
				},
			},
			PublisherID: 5,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2Fbanner.jpg?alt=media&token=b403f94b-9ebc-496d-a8a6-36a34fed085e",
		},
		{
			Name:        "Underlords",
			Description: "NEXT GENERATION AUTO-BATTLER\nIn Dota Underlords, strategic decisions matter more than twitch reflexes. Underlords includes compelling singleplayer and multiplayer modes, and offers level progression with rewards. Play a strategic Standard game, a quick Knockout match, or co-op Duos match with a friend.\n\nSEASON ONE NOW AVAILABLE\nSeason One comes with a City Crawl full of content, a Battle Pass full of rewards, and multiple ways to play online or offline. Dota Underlords is now out of Early Access and ready to play!",
			ReleaseDate: time.Now(),
			Genres: []*model.Genre{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
				{
					ID: 10,
				},
			},
			Tags: []*model.Tag{
				{
					ID: 5,
				},
				{
					ID: 8,
				},
				{
					ID: 11,
				},
			},
			OriginalPrice:      0,
			PromoID: 1,
			Developers: []*model.Developer{
				{
					ID: 1,
				},
			},
			PublisherID: 1,
			SystemID:    1,
			Users:       nil,
			Banner:      "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F7%2Fbanner.jpg?alt=media&token=877a6401-3b03-4e09-8caa-ce2c9f25e2de",
		},
	}

	for _, game := range games {
		db.Create(&game)
		db.Save(&game)
	}
}

func SeedDeveloper(db *gorm.DB) {
	developers := []model.Developer{
		{
			Name: "Valve",
		},
		{
			Name: "Respawn Entertainment",
		},
		{
			Name: "Capcom",
		},
		{
			Name: "Hidden Path Entertainment",
		},
		{
			Name: "Rockstar North",
		},
		{
			Name: "Re-Logic",
		},
	}

	for _, developer := range developers {
		db.Create(&developer)
	}
}

func SeedCountry(db *gorm.DB) {
	countries := []model.Country{
		{Name: "Afghanistan"},
		{Name: "Åland Islands"},
		{Name: "Albania"},
		{Name: "Algeria"},
		{Name: "American Samoa"},
		{Name: "AndorrA"},
		{Name: "Angola"},
		{Name: "Anguilla"},
		{Name: "Antarctica"},
		{Name: "Antigua and Barbuda"},
		{Name: "Argentina"},
		{Name: "Armenia"},
		{Name: "Aruba"},
		{Name: "Australia"},
		{Name: "Austria"},
		{Name: "Azerbaijan"},
		{Name: "Bahamas"},
		{Name: "Bahrain"},
		{Name: "Bangladesh"},
		{Name: "Barbados"},
		{Name: "Belarus"},
		{Name: "Belgium"},
		{Name: "Belize"},
		{Name: "Benin"},
		{Name: "Bermuda"},
		{Name: "Bhutan"},
		{Name: "Bolivia"},
		{Name: "Bosnia and Herzegovina"},
		{Name: "Botswana"},
		{Name: "Bouvet Island"},
		{Name: "Brazil"},
		{Name: "British Indian Ocean Territory"},
		{Name: "Brunei Darussalam"},
		{Name: "Bulgaria"},
		{Name: "Burkina Faso"},
		{Name: "Burundi"},
		{Name: "Cambodia"},
		{Name: "Cameroon"},
		{Name: "Canada"},
		{Name: "Cape Verde"},
		{Name: "Cayman Islands"},
		{Name: "Central African Republic"},
		{Name: "Chad"},
		{Name: "Chile"},
		{Name: "China"},
		{Name: "Christmas Island"},
		{Name: "Cocos (Keeling) Islands"},
		{Name: "Colombia"},
		{Name: "Comoros"},
		{Name: "Congo"},
		{Name: "Congo, The Democratic Republic of the"},
		{Name: "Cook Islands"},
		{Name: "Costa Rica"},
		{Name: "Cote D\"Ivoire"},
		{Name: "Croatia"},
		{Name: "Cuba"},
		{Name: "Cyprus"},
		{Name: "Czech Republic"},
		{Name: "Denmark"},
		{Name: "Djibouti"},
		{Name: "Dominica"},
		{Name: "Dominican Republic"},
		{Name: "Ecuador"},
		{Name: "Egypt"},
		{Name: "El Salvador"},
		{Name: "Equatorial Guinea"},
		{Name: "Eritrea"},
		{Name: "Estonia"},
		{Name: "Ethiopia"},
		{Name: "Falkland Islands (Malvinas)"},
		{Name: "Faroe Islands"},
		{Name: "Fiji"},
		{Name: "Finland"},
		{Name: "France"},
		{Name: "French Guiana"},
		{Name: "French Polynesia"},
		{Name: "French Southern Territories"},
		{Name: "Gabon"},
		{Name: "Gambia"},
		{Name: "Georgia"},
		{Name: "Germany"},
		{Name: "Ghana"},
		{Name: "Gibraltar"},
		{Name: "Greece"},
		{Name: "Greenland"},
		{Name: "Grenada"},
		{Name: "Guadeloupe"},
		{Name: "Guam"},
		{Name: "Guatemala"},
		{Name: "Guernsey"},
		{Name: "Guinea"},
		{Name: "Guinea-Bissau"},
		{Name: "Guyana"},
		{Name: "Haiti"},
		{Name: "Heard Island and Mcdonald Islands"},
		{Name: "Holy See (Vatican City State)"},
		{Name: "Honduras"},
		{Name: "Hong Kong"},
		{Name: "Hungary"},
		{Name: "Iceland"},
		{Name: "India"},
		{Name: "Indonesia"},
		{Name: "Iran, Islamic Republic Of"},
		{Name: "Iraq"},
		{Name: "Ireland"},
		{Name: "Isle of Man"},
		{Name: "Israel"},
		{Name: "Italy"},
		{Name: "Jamaica"},
		{Name: "Japan"},
		{Name: "Jersey"},
		{Name: "Jordan"},
		{Name: "Kazakhstan"},
		{Name: "Kenya"},
		{Name: "Kiribati"},
		{Name: "Korea, Democratic People\"S Republic of"},
		{Name: "Korea, Republic of"},
		{Name: "Kuwait"},
		{Name: "Kyrgyzstan"},
		{Name: "Lao People\"S Democratic Republic"},
		{Name: "Latvia"},
		{Name: "Lebanon"},
		{Name: "Lesotho"},
		{Name: "Liberia"},
		{Name: "Libyan Arab Jamahiriya"},
		{Name: "Liechtenstein"},
		{Name: "Lithuania"},
		{Name: "Luxembourg"},
		{Name: "Macao"},
		{Name: "Macedonia, The Former Yugoslav Republic of"},
		{Name: "Madagascar"},
		{Name: "Malawi"},
		{Name: "Malaysia"},
		{Name: "Maldives"},
		{Name: "Mali"},
		{Name: "Malta"},
		{Name: "Marshall Islands"},
		{Name: "Martinique"},
		{Name: "Mauritania"},
		{Name: "Mauritius"},
		{Name: "Mayotte"},
		{Name: "Mexico"},
		{Name: "Micronesia, Federated States of"},
		{Name: "Moldova, Republic of"},
		{Name: "Monaco"},
		{Name: "Mongolia"},
		{Name: "Montserrat"},
		{Name: "Morocco"},
		{Name: "Mozambique"},
		{Name: "Myanmar"},
		{Name: "Namibia"},
		{Name: "Nauru"},
		{Name: "Nepal"},
		{Name: "Netherlands"},
		{Name: "Netherlands Antilles"},
		{Name: "New Caledonia"},
		{Name: "New Zealand"},
		{Name: "Nicaragua"},
		{Name: "Niger"},
		{Name: "Nigeria"},
		{Name: "Niue"},
		{Name: "Norfolk Island"},
		{Name: "Northern Mariana Islands"},
		{Name: "Norway"},
		{Name: "Oman"},
		{Name: "Pakistan"},
		{Name: "Palau"},
		{Name: "Palestinian Territory, Occupied"},
		{Name: "Panama"},
		{Name: "Papua New Guinea"},
		{Name: "Paraguay"},
		{Name: "Peru"},
		{Name: "Philippines"},
		{Name: "Pitcairn"},
		{Name: "Poland"},
		{Name: "Portugal"},
		{Name: "Puerto Rico"},
		{Name: "Qatar"},
		{Name: "Reunion"},
		{Name: "Romania"},
		{Name: "Russian Federation"},
		{Name: "RWANDA"},
		{Name: "Saint Helena"},
		{Name: "Saint Kitts and Nevis"},
		{Name: "Saint Lucia"},
		{Name: "Saint Pierre and Miquelon"},
		{Name: "Saint Vincent and the Grenadines"},
		{Name: "Samoa"},
		{Name: "San Marino"},
		{Name: "Sao Tome and Principe"},
		{Name: "Saudi Arabia"},
		{Name: "Senegal"},
		{Name: "Serbia and Montenegro"},
		{Name: "Seychelles"},
		{Name: "Sierra Leone"},
		{Name: "Singapore"},
		{Name: "Slovakia"},
		{Name: "Slovenia"},
		{Name: "Solomon Islands"},
		{Name: "Somalia"},
		{Name: "South Africa"},
		{Name: "South Georgia and the South Sandwich Islands"},
		{Name: "Spain"},
		{Name: "Sri Lanka"},
		{Name: "Sudan"},
		{Name: "SuriName"},
		{Name: "Svalbard and Jan Mayen"},
		{Name: "Swaziland"},
		{Name: "Sweden"},
		{Name: "Switzerland"},
		{Name: "Syrian Arab Republic"},
		{Name: "Taiwan, Province of China"},
		{Name: "Tajikistan"},
		{Name: "Tanzania, United Republic of"},
		{Name: "Thailand"},
		{Name: "Timor-Leste"},
		{Name: "Togo"},
		{Name: "Tokelau"},
		{Name: "Tonga"},
		{Name: "Trinidad and Tobago"},
		{Name: "Tunisia"},
		{Name: "Turkey"},
		{Name: "Turkmenistan"},
		{Name: "Turks and Caicos Islands"},
		{Name: "Tuvalu"},
		{Name: "Uganda"},
		{Name: "Ukraine"},
		{Name: "United Arab Emirates"},
		{Name: "United Kingdom"},
		{Name: "United States"},
		{Name: "United States Minor Outlying Islands"},
		{Name: "Uruguay"},
		{Name: "Uzbekistan"},
		{Name: "Vanuatu"},
		{Name: "Venezuela"},
		{Name: "Viet Nam"},
		{Name: "Virgin Islands, British"},
		{Name: "Virgin Islands, U.S."},
		{Name: "Wallis and Futuna"},
		{Name: "Western Sahara"},
		{Name: "Yemen"},
		{Name: "Zambia"},
		{Name: "Zimbabwe"},
	}

	for _, country := range countries {
		db.Create(&country)
	}
}

func SeedUser(db *gorm.DB) {
	users := []model.User{
		{
			AccountName:       "nico",
			ProfileName:       "nico tanzil",
			RealName:          "Nico Tanzil",
			Email:             "nico@mail.com",
			Password:          providers.HashPassword("password"),
			Balance:           100,
			CustomURL:         "nico",
			Summary:           "No information given.",
			Avatar:            firebase_data.Avatar,
			ProfileBackground: firebase_data.ProfileBackground,
			Country:           &model.Country{ID: 102},
			Games: []*model.Game{
				{
					ID: 1,
				},
			},
			Experience:        550,
			IsSuspend:         true,
			SuspensionRequest: &model.SuspensionRequest{ID: 2},
		},
		{
			AccountName:       "william",
			ProfileName:       "William",
			RealName:          "William Martin",
			Email:             "will@mail.com",
			Password:          providers.HashPassword("password"),
			Balance:           100,
			CustomURL:         "william",
			Summary:           "No information given.",
			Avatar:            firebase_data.Avatar,
			ProfileBackground: firebase_data.ProfileBackground,
			Country:           &model.Country{ID: 2},
			SuspensionRequest: &model.SuspensionRequest{ID: 1},
		},
		{
			AccountName:       "ricko",
			ProfileName:       "Ricko",
			RealName:          "Ricko Adrio",
			Email:             "rick@mail.com",
			Password:          providers.HashPassword("password"),
			Balance:           100,
			CustomURL:         "rick",
			Summary:           "No information given.",
			Avatar:            firebase_data.Avatar,
			ProfileBackground: firebase_data.ProfileBackground,
			Country:           &model.Country{ID: 3},
			SuspensionRequest: &model.SuspensionRequest{ID: 1},
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}

func SeedUserReport(db *gorm.DB) {
	userReports := []model.UserReport{
		{
			Reported:    &model.User{ID: 1},
			Reporter:    &model.User{ID: 2},
			Description: "Toxic user!",
		},
		{
			Reported:    &model.User{ID: 1},
			Reporter:    &model.User{ID: 3},
			Description: "Careful scammer!",
		},
	}

	for _, report := range userReports {
		db.Create(&report)
	}
}

func SeedSuspensionRequest(db *gorm.DB) {
	suspensionRequests := []model.SuspensionRequest{
		{
			Description: "No suspension request.",
		},
		{
			Description: "Please unsuspend me!",
		},
	}

	for _, request := range suspensionRequests {
		db.Create(&request)
	}
}
