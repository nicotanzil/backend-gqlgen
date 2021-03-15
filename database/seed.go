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
	dbClose, _ := db.DB()
	defer dbClose.Close()

	fmt.Println("[INFO] SEEDING...")
	// SEED ALL NECESSARY TABLE

	SeedBadge(db)
	SeedProfileBackground(db)
	SeedMiniProfileBackground(db)
	SeedAvatarFrames(db)
	SeedTheme(db)
	SeedChatSticker(db)
	SeedAnimatedAvatar(db)

	SeedAdmin(db)
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
	SeedSuspensionRequest(db)
	SeedFriendRequest(db)
	SeedUserReport(db)

	SeedCommunityArtPost(db)
	SeedCommunityArtPostReview(db)

	SeedCommunityGameReview(db)
	SeedCommunityGameReviewComment(db)

	SeedGameDiscussion(db)
	SeedGameDiscussionReply(db)

	SeedCart(db)
	SeedWishlist(db)

	SeedPaymentType(db)
	SeedGift(db)
	SeedTransactionHeader(db)
	SeedTransactionDetail(db)

	SeedItemType(db)
	SeedItem(db)
	SeedBid(db)
	SeedItemTransaction(db)
	SeedSellListing(db)
	SeedBuyListing(db)
	SeedNewItemNotification(db)

	SeedWalletCode(db)

	fmt.Println("[INFO] SEEDED.")
}

func SeedBadge(db *gorm.DB) {
	badges := []model.Badge{
		{
			Name: "Community Patron",
			Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fbadges%2Fcommunity%20patron.png?alt=media&token=89bd7b7c-31eb-4640-b10a-5f6f4dba086d",
			Xp:   0,
		},
		{
			Name: "Years of Service",
			Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fbadges%2Fyears%20of%20service.png?alt=media&token=d6f98e67-d7b5-4223-be5e-5a59a83cf6b6",
			Xp:   100,
		},
		{
			Name: "Game Mechanic",
			Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fbadges%2Fgame%20mechanic.png?alt=media&token=192f2fd3-111c-4b0e-97fa-f1f1007ef3a6",
			Xp:   500,
		},
		{
			Name: "Ganker",
			Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fbadges%2Fganker.png?alt=media&token=00f8c813-0a12-41e9-b2aa-11aa745b982a",
			Xp:   550,
		},
	}

	for _, badge := range badges {
		db.Create(&badge)
	}
}

func SeedProfileBackground(db *gorm.DB) {
	profileBackgrounds := []model.ProfileBackground{
		{
			Name:  "Default",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fusers%2Fdefault%2Fprofile_background.jpg?alt=media&token=d5b3f65c-37d3-45b8-ad72-7bb74948985f",
			Price: 0,
		},
		{
			Name:  "Cabin Lake",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fprofile_backgrounds%2Fcabin%20lake.jpg?alt=media&token=7d2bf329-6dd8-4d3d-a913-34f321dbc6c5",
			Price: 500,
		},
		{
			Name:  "Cliff's Edge",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fprofile_backgrounds%2Fcliff's%20edge.jpg?alt=media&token=6ca20d50-5f48-436e-9925-fab701e307f6",
			Price: 500,
		},
		{
			Name:  "Dying Light",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fprofile_backgrounds%2Fdying%20light.jpg?alt=media&token=8206d03b-84f0-412b-af8b-df79abbef4e2",
			Price: 1000,
		},
		{
			Name:  "Portrait",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fprofile_backgrounds%2Fportrait.jpg?alt=media&token=437912bc-331f-4be9-b345-eeb7e2806fde",
			Price: 500,
		},
		{
			Name:  "Rose Petals",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fprofile_backgrounds%2Frose%20petals.jpg?alt=media&token=1bb4be96-4b73-4bbe-89b4-7ce41e2a1e8e",
			Price: 750,
		},
	}

	for _, profileBackground := range profileBackgrounds {
		db.Create(&profileBackground)
	}
}

func SeedMiniProfileBackground(db *gorm.DB) {
	miniProfileBackgrounds := []model.MiniProfileBackground{
		{
			Name:  "Default",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fmini_profile_backgrounds%2Fdefault.gif?alt=media&token=5c0c7bd7-195a-4098-b251-d392b1ff3f74",
			Price: 0,
		},
		{
			Name:  "Retro",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fmini_profile_backgrounds%2FRetro.gif?alt=media&token=f3d4d523-9fea-47f7-b6f6-5187990b2579",
			Price: 1000,
		},
		{
			Name:  "Sunset",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fmini_profile_backgrounds%2Fsunset.gif?alt=media&token=a339f7cd-099d-4b2d-ad20-987bfb654394",
			Price: 500,
		},
		{
			Name:  "Universe",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fmini_profile_backgrounds%2Funiverse.gif?alt=media&token=e4dae938-f563-42f4-a5e8-706710cf8881",
			Price: 1000,
		},
	}

	for _, miniProfileBackground := range miniProfileBackgrounds {
		db.Create(&miniProfileBackground)
	}
}

func SeedAvatarFrames(db *gorm.DB) {
	avatarFrames := []model.AvatarFrame{
		{
			Name:  "Chase",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Favatar_frames%2Fchase.png?alt=media&token=72bca6bd-df64-485e-890e-467372d554dc",
			Price: 2000,
		},
		{
			Name:  "Glitch",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Favatar_frames%2Fglitch.png?alt=media&token=cf488486-e76b-49b3-8c42-78c4efd25dd5",
			Price: 1500,
		},
		{
			Name:  "Neon Glitch",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Favatar_frames%2Fneon%20glitch.png?alt=media&token=20d201fd-fcce-4926-a39e-8df02630ca20",
			Price: 2000,
		},
		{
			Name:  "Shadow Fire",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Favatar_frames%2Fshadow%20fire.png?alt=media&token=4557315a-390b-4535-a581-b09b45ab426c",
			Price: 2000,
		},
	}

	for _, avatarFrame := range avatarFrames {
		db.Create(&avatarFrame)
	}
}

func SeedTheme(db *gorm.DB) {
	themes := []model.Theme{
		{
			Name:  "Default Theme",
			Color: "#2C3144",
		},
		{
			Name:  "Summer",
			Color: "#A66F17",
		},
		{
			Name:  "Midnight",
			Color: "#16145E",
		},
		{
			Name:  "Steel",
			Color: "#313944",
		},
		{
			Name:  "Cosmic",
			Color: "#5A215F",
		},
		{
			Name:  "Dark Mode",
			Color: "#131313",
		},
	}

	for _, theme := range themes {
		db.Create(&theme)
	}
}

func SeedChatSticker(db *gorm.DB) {
	chatStickers := []model.ChatSticker{
		{
			Name:  "Chicken Dance",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fchat-sticker%2Fchicken%20dance.png?alt=media&token=dfda536e-c6f0-45b3-95a8-990d7cefdef7",
			Price: 1000,
		},
		{
			Name:  "Goose Chase",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fchat-sticker%2Fgoose%20chase.png?alt=media&token=95454cd7-f75d-4b61-b1bb-c98ef4691f30",
			Price: 1000,
		},
		{
			Name:  "Pyro",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fchat-sticker%2Fpyro.png?alt=media&token=3bbce705-e807-48e9-8faa-02a84266b4d7",
			Price: 1000,
		},
		{
			Name:  "Rock On",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fchat-sticker%2Frock%20on.png?alt=media&token=7bd44412-ea08-4aec-acc0-bdb534e31b45",
			Price: 1000,
		},
		{
			Name:  "Toast",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fchat-sticker%2Ftoast.png?alt=media&token=3ec82f5f-cfbe-4f68-8616-54aff957bd2f",
			Price: 1000,
		},
	}

	for _, chatSticker := range chatStickers {
		db.Create(&chatSticker)
	}
}

func SeedAnimatedAvatar(db *gorm.DB) {
	animatedAvatars := []model.AnimatedAvatar{
		{
			Name:  "Ignis Sanat",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fignis%20sanat.gif?alt=media&token=2550c338-3f48-4003-a1c3-0078d5fbad7e",
			Price: 3000,
		},
		{
			Name:  "Shadow Demon",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fshadow%20demon.gif?alt=media&token=b127ad3d-9f43-4139-af47-18c013ae8fb0",
			Price: 3000,
		},
		{
			Name:  "Shadow Fiend",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fshadow%20fiend.gif?alt=media&token=d3aa6193-d9a3-4b13-b4bd-6b1b86813479",
			Price: 3000,
		},
		{
			Name:  "Space Sphere",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fspace%20sphere.gif?alt=media&token=4eb27ffc-b9b5-4d67-bff5-520aeaf70f67",
			Price: 3000,
		},
		{
			Name:  "Wheatley",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fwheatley.gif?alt=media&token=4bb97ee0-08a1-4dae-92b8-5d56bdce0865",
			Price: 3000,
		},
		{
			Name:  "Wizard Skull",
			Link:  "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fanimated_avatars%2Fwizard%20skull.gif?alt=media&token=7816d970-0bd8-457a-ad6e-e56fdf3cf77a",
			Price: 3000,
		},
	}

	for _, avatar := range animatedAvatars {
		db.Create(&avatar)
	}
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
	gameVideos := []model.GameVideo{
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fvideos%2F1.mp4?alt=media&token=915b8004-f5d1-4279-81fc-2c1f7e56f102"},
		{GameID: 1, Link: "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fvideos%2F2.mp4?alt=media&token=4cbc40d4-0f03-41e7-97f4-1c3154d34bbc"},
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
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 30).Unix()),
		},
		{
			DiscountPercentage: 50,
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 3).Unix()),
		},
		{
			DiscountPercentage: 10,
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 7).Unix()),
		},
		{
			DiscountPercentage: 75,
			ValidUntil:         int(time.Now().Add(time.Hour * 24 * 1).Unix()),
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
			OriginalPrice: 0,
			PromoID:       1,
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
			OriginalPrice: 562000,
			PromoID:       2,
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
			OriginalPrice: 650000,
			PromoID:       3,
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
			OriginalPrice: 0,
			PromoID:       1,
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
			OriginalPrice: 300000,
			PromoID:       4,
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
			OriginalPrice: 89999,
			PromoID:       1,
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
			OriginalPrice: 0,
			PromoID:       1,
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
	}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
	//for _, game := range games {
	//	db.Create(&game)
	//}
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
		{Latitude: 42.546245, Longitude: 1.601554, Name: "Andorra"},
		{Latitude: 23.424076, Longitude: 53.847818, Name: "United Arab Emirates"},
		{Latitude: 33.93911, Longitude: 67.709953, Name: "Afghanistan"},
		{Latitude: 17.060816, Longitude: -61.796428, Name: "Antigua and Barbuda"},
		{Latitude: 18.220554, Longitude: -63.068615, Name: "Anguilla"},
		{Latitude: 41.153332, Longitude: 20.168331, Name: "Albania"},
		{Latitude: 40.069099, Longitude: 45.038189, Name: "Armenia"},
		{Latitude: 12.226079, Longitude: -69.060087, Name: "Netherlands Antilles"},
		{Latitude: -11.202692, Longitude: 17.873887, Name: "Angola"},
		{Latitude: -75.250973, Longitude: -0.071389, Name: "Antarctica"},
		{Latitude: -38.416097, Longitude: -63.616672, Name: "Argentina"},
		{Latitude: -14.270972, Longitude: -170.132217, Name: "American Samoa"},
		{Latitude: 47.516231, Longitude: 14.550072, Name: "Austria"},
		{Latitude: -25.274398, Longitude: 133.775136, Name: "Australia"},
		{Latitude: 12.52111, Longitude: -69.968338, Name: "Aruba"},
		{Latitude: 40.143105, Longitude: 47.576927, Name: "Azerbaijan"},
		{Latitude: 43.915886, Longitude: 17.679076, Name: "Bosnia and Herzegovina"},
		{Latitude: 13.193887, Longitude: -59.543198, Name: "Barbados"},
		{Latitude: 23.684994, Longitude: 90.356331, Name: "Bangladesh"},
		{Latitude: 50.503887, Longitude: 4.469936, Name: "Belgium"},
		{Latitude: 12.238333, Longitude: -1.561593, Name: "Burkina Faso"},
		{Latitude: 42.733883, Longitude: 25.48583, Name: "Bulgaria"},
		{Latitude: 25.930414, Longitude: 50.637772, Name: "Bahrain"},
		{Latitude: -3.373056, Longitude: 29.918886, Name: "Burundi"},
		{Latitude: 9.30769, Longitude: 2.315834, Name: "Benin"},
		{Latitude: 32.321384, Longitude: -64.75737, Name: "Bermuda"},
		{Latitude: 4.535277, Longitude: 114.727669, Name: "Brunei"},
		{Latitude: -16.290154, Longitude: -63.588653, Name: "Bolivia"},
		{Latitude: -14.235004, Longitude: -51.92528, Name: "Brazil"},
		{Latitude: 25.03428, Longitude: -77.39628, Name: "Bahamas"},
		{Latitude: 27.514162, Longitude: 90.433601, Name: "Bhutan"},
		{Latitude: -54.423199, Longitude: 3.413194, Name: "Bouvet Island"},
		{Latitude: -22.328474, Longitude: 24.684866, Name: "Botswana"},
		{Latitude: 53.709807, Longitude: 27.953389, Name: "Belarus"},
		{Latitude: 17.189877, Longitude: -88.49765, Name: "Belize"},
		{Latitude: 56.130366, Longitude: -106.346771, Name: "Canada"},
		{Latitude: -12.164165, Longitude: 96.870956, Name: "Cocos [Keeling] Islands"},
		{Latitude: -4.038333, Longitude: 21.758664, Name: "Congo [DRC]"},
		{Latitude: 6.611111, Longitude: 20.939444, Name: "Central African Republic"},
		{Latitude: -0.228021, Longitude: 15.827659, Name: "Congo [Republic]"},
		{Latitude: 46.818188, Longitude: 8.227512, Name: "Switzerland"},
		{Latitude: 7.539989, Longitude: -5.54708, Name: "Côte d'Ivoire"},
		{Latitude: -21.236736, Longitude: -159.777671, Name: "Cook Islands"},
		{Latitude: -35.675147, Longitude: -71.542969, Name: "Chile"},
		{Latitude: 7.369722, Longitude: 12.354722, Name: "Cameroon"},
		{Latitude: 35.86166, Longitude: 104.195397, Name: "China"},
		{Latitude: 4.570868, Longitude: -74.297333, Name: "Colombia"},
		{Latitude: 9.748917, Longitude: -83.753428, Name: "Costa Rica"},
		{Latitude: 21.521757, Longitude: -77.781167, Name: "Cuba"},
		{Latitude: 16.002082, Longitude: -24.013197, Name: "Cape Verde"},
		{Latitude: -10.447525, Longitude: 105.690449, Name: "Christmas Island"},
		{Latitude: 35.126413, Longitude: 33.429859, Name: "Cyprus"},
		{Latitude: 49.817492, Longitude: 15.472962, Name: "Czech Republic"},
		{Latitude: 51.165691, Longitude: 10.451526, Name: "Germany"},
		{Latitude: 11.825138, Longitude: 42.590275, Name: "Djibouti"},
		{Latitude: 56.26392, Longitude: 9.501785, Name: "Denmark"},
		{Latitude: 15.414999, Longitude: -61.370976, Name: "Dominica"},
		{Latitude: 18.735693, Longitude: -70.162651, Name: "Dominican Republic"},
		{Latitude: 28.033886, Longitude: 1.659626, Name: "Algeria"},
		{Latitude: -1.831239, Longitude: -78.183406, Name: "Ecuador"},
		{Latitude: 58.595272, Longitude: 25.013607, Name: "Estonia"},
		{Latitude: 26.820553, Longitude: 30.802498, Name: "Egypt"},
		{Latitude: 24.215527, Longitude: -12.885834, Name: "Western Sahara"},
		{Latitude: 15.179384, Longitude: 39.782334, Name: "Eritrea"},
		{Latitude: 40.463667, Longitude: -3.74922, Name: "Spain"},
		{Latitude: 9.145, Longitude: 40.489673, Name: "Ethiopia"},
		{Latitude: 61.92411, Longitude: 25.748151, Name: "Finland"},
		{Latitude: -16.578193, Longitude: 179.414413, Name: "Fiji"},
		{Latitude: -51.796253, Longitude: -59.523613, Name: "Falkland Islands [Islas Malvinas]"},
		{Latitude: 7.425554, Longitude: 150.550812, Name: "Micronesia"},
		{Latitude: 61.892635, Longitude: -6.911806, Name: "Faroe Islands"},
		{Latitude: 46.227638, Longitude: 2.213749, Name: "France"},
		{Latitude: -0.803689, Longitude: 11.609444, Name: "Gabon"},
		{Latitude: 55.378051, Longitude: -3.435973, Name: "United Kingdom"},
		{Latitude: 12.262776, Longitude: -61.604171, Name: "Grenada"},
		{Latitude: 42.315407, Longitude: 43.356892, Name: "Georgia"},
		{Latitude: 3.933889, Longitude: -53.125782, Name: "French Guiana"},
		{Latitude: 49.465691, Longitude: -2.585278, Name: "Guernsey"},
		{Latitude: 7.946527, Longitude: -1.023194, Name: "Ghana"},
		{Latitude: 36.137741, Longitude: -5.345374, Name: "Gibraltar"},
		{Latitude: 71.706936, Longitude: -42.604303, Name: "Greenland"},
		{Latitude: 13.443182, Longitude: -15.310139, Name: "Gambia"},
		{Latitude: 9.945587, Longitude: -9.696645, Name: "Guinea"},
		{Latitude: 16.995971, Longitude: -62.067641, Name: "Guadeloupe"},
		{Latitude: 1.650801, Longitude: 10.267895, Name: "Equatorial Guinea"},
		{Latitude: 39.074208, Longitude: 21.824312, Name: "Greece"},
		{Latitude: -54.429579, Longitude: -36.587909, Name: "South Georgia and the South Sandwich Islands"},
		{Latitude: 15.783471, Longitude: -90.230759, Name: "Guatemala"},
		{Latitude: 13.444304, Longitude: 144.793731, Name: "Guam"},
		{Latitude: 11.803749, Longitude: -15.180413, Name: "Guinea-Bissau"},
		{Latitude: 4.860416, Longitude: -58.93018, Name: "Guyana"},
		{Latitude: 31.354676, Longitude: 34.308825, Name: "Gaza Strip"},
		{Latitude: 22.396428, Longitude: 114.109497, Name: "Hong Kong"},
		{Latitude: -53.08181, Longitude: 73.504158, Name: "Heard Island and McDonald Islands"},
		{Latitude: 15.199999, Longitude: -86.241905, Name: "Honduras"},
		{Latitude: 45.1, Longitude: 15.2, Name: "Croatia"},
		{Latitude: 18.971187, Longitude: -72.285215, Name: "Haiti"},
		{Latitude: 47.162494, Longitude: 19.503304, Name: "Hungary"},
		{Latitude: -0.789275, Longitude: 113.921327, Name: "Indonesia"},
		{Latitude: 53.41291, Longitude: -8.24389, Name: "Ireland"},
		{Latitude: 31.046051, Longitude: 34.851612, Name: "Israel"},
		{Latitude: 54.236107, Longitude: -4.548056, Name: "Isle of Man"},
		{Latitude: 20.593684, Longitude: 78.96288, Name: "India"},
		{Latitude: -6.343194, Longitude: 71.876519, Name: "British Indian Ocean Territory"},
		{Latitude: 33.223191, Longitude: 43.679291, Name: "Iraq"},
		{Latitude: 32.427908, Longitude: 53.688046, Name: "Iran"},
		{Latitude: 64.963051, Longitude: -19.020835, Name: "Iceland"},
		{Latitude: 41.87194, Longitude: 12.56738, Name: "Italy"},
		{Latitude: 49.214439, Longitude: -2.13125, Name: "Jersey"},
		{Latitude: 18.109581, Longitude: -77.297508, Name: "Jamaica"},
		{Latitude: 30.585164, Longitude: 36.238414, Name: "Jordan"},
		{Latitude: 36.204824, Longitude: 138.252924, Name: "Japan"},
		{Latitude: -0.023559, Longitude: 37.906193, Name: "Kenya"},
		{Latitude: 41.20438, Longitude: 74.766098, Name: "Kyrgyzstan"},
		{Latitude: 12.565679, Longitude: 104.990963, Name: "Cambodia"},
		{Latitude: -3.370417, Longitude: -168.734039, Name: "Kiribati"},
		{Latitude: -11.875001, Longitude: 43.872219, Name: "Comoros"},
		{Latitude: 17.357822, Longitude: -62.782998, Name: "Saint Kitts and Nevis"},
		{Latitude: 40.339852, Longitude: 127.510093, Name: "North Korea"},
		{Latitude: 35.907757, Longitude: 127.766922, Name: "South Korea"},
		{Latitude: 29.31166, Longitude: 47.481766, Name: "Kuwait"},
		{Latitude: 19.513469, Longitude: -80.566956, Name: "Cayman Islands"},
		{Latitude: 48.019573, Longitude: 66.923684, Name: "Kazakhstan"},
		{Latitude: 19.85627, Longitude: 102.495496, Name: "Laos"},
		{Latitude: 33.854721, Longitude: 35.862285, Name: "Lebanon"},
		{Latitude: 13.909444, Longitude: -60.978893, Name: "Saint Lucia"},
		{Latitude: 47.166, Longitude: 9.555373, Name: "Liechtenstein"},
		{Latitude: 7.873054, Longitude: 80.771797, Name: "Sri Lanka"},
		{Latitude: 6.428055, Longitude: -9.429499, Name: "Liberia"},
		{Latitude: -29.609988, Longitude: 28.233608, Name: "Lesotho"},
		{Latitude: 55.169438, Longitude: 23.881275, Name: "Lithuania"},
		{Latitude: 49.815273, Longitude: 6.129583, Name: "Luxembourg"},
		{Latitude: 56.879635, Longitude: 24.603189, Name: "Latvia"},
		{Latitude: 26.3351, Longitude: 17.228331, Name: "Libya"},
		{Latitude: 31.791702, Longitude: -7.09262, Name: "Morocco"},
		{Latitude: 43.750298, Longitude: 7.412841, Name: "Monaco"},
		{Latitude: 47.411631, Longitude: 28.369885, Name: "Moldova"},
		{Latitude: 42.708678, Longitude: 19.37439, Name: "Montenegro"},
		{Latitude: -18.766947, Longitude: 46.869107, Name: "Madagascar"},
		{Latitude: 7.131474, Longitude: 171.184478, Name: "Marshall Islands"},
		{Latitude: 41.608635, Longitude: 21.745275, Name: "Macedonia [FYROM]"},
		{Latitude: 17.570692, Longitude: -3.996166, Name: "Mali"},
		{Latitude: 21.913965, Longitude: 95.956223, Name: "Myanmar [Burma]"},
		{Latitude: 46.862496, Longitude: 103.846656, Name: "Mongolia"},
		{Latitude: 22.198745, Longitude: 113.543873, Name: "Macau"},
		{Latitude: 17.33083, Longitude: 145.38469, Name: "Northern Mariana Islands"},
		{Latitude: 14.641528, Longitude: -61.024174, Name: "Martinique"},
		{Latitude: 21.00789, Longitude: -10.940835, Name: "Mauritania"},
		{Latitude: 16.742498, Longitude: -62.187366, Name: "Montserrat"},
		{Latitude: 35.937496, Longitude: 14.375416, Name: "Malta"},
		{Latitude: -20.348404, Longitude: 57.552152, Name: "Mauritius"},
		{Latitude: 3.202778, Longitude: 73.22068, Name: "Maldives"},
		{Latitude: -13.254308, Longitude: 34.301525, Name: "Malawi"},
		{Latitude: 23.634501, Longitude: -102.552784, Name: "Mexico"},
		{Latitude: 4.210484, Longitude: 101.975766, Name: "Malaysia"},
		{Latitude: -18.665695, Longitude: 35.529562, Name: "Mozambique"},
		{Latitude: -22.95764, Longitude: 18.49041, Name: "Namibia"},
		{Latitude: -20.904305, Longitude: 165.618042, Name: "New Caledonia"},
		{Latitude: 17.607789, Longitude: 8.081666, Name: "Niger"},
		{Latitude: -29.040835, Longitude: 167.954712, Name: "Norfolk Island"},
		{Latitude: 9.081999, Longitude: 8.675277, Name: "Nigeria"},
		{Latitude: 12.865416, Longitude: -85.207229, Name: "Nicaragua"},
		{Latitude: 52.132633, Longitude: 5.291266, Name: "Netherlands"},
		{Latitude: 60.472024, Longitude: 8.468946, Name: "Norway"},
		{Latitude: 28.394857, Longitude: 84.124008, Name: "Nepal"},
		{Latitude: -0.522778, Longitude: 166.931503, Name: "Nauru"},
		{Latitude: -19.054445, Longitude: -169.867233, Name: "Niue"},
		{Latitude: -40.900557, Longitude: 174.885971, Name: "New Zealand"},
		{Latitude: 21.512583, Longitude: 55.923255, Name: "Oman"},
		{Latitude: 8.537981, Longitude: -80.782127, Name: "Panama"},
		{Latitude: -9.189967, Longitude: -75.015152, Name: "Peru"},
		{Latitude: -17.679742, Longitude: -149.406843, Name: "French Polynesia"},
		{Latitude: -6.314993, Longitude: 143.95555, Name: "Papua New Guinea"},
		{Latitude: 12.879721, Longitude: 121.774017, Name: "Philippines"},
		{Latitude: 30.375321, Longitude: 69.345116, Name: "Pakistan"},
		{Latitude: 51.919438, Longitude: 19.145136, Name: "Poland"},
		{Latitude: 46.941936, Longitude: -56.27111, Name: "Saint Pierre and Miquelon"},
		{Latitude: -24.703615, Longitude: -127.439308, Name: "Pitcairn Islands"},
		{Latitude: 18.220833, Longitude: -66.590149, Name: "Puerto Rico"},
		{Latitude: 31.952162, Longitude: 35.233154, Name: "Palestinian Territories"},
		{Latitude: 39.399872, Longitude: -8.224454, Name: "Portugal"},
		{Latitude: 7.51498, Longitude: 134.58252, Name: "Palau"},
		{Latitude: -23.442503, Longitude: -58.443832, Name: "Paraguay"},
		{Latitude: 25.354826, Longitude: 51.183884, Name: "Qatar"},
		{Latitude: -21.115141, Longitude: 55.536384, Name: "Réunion"},
		{Latitude: 45.943161, Longitude: 24.96676, Name: "Romania"},
		{Latitude: 44.016521, Longitude: 21.005859, Name: "Serbia"},
		{Latitude: 61.52401, Longitude: 105.318756, Name: "Russia"},
		{Latitude: -1.940278, Longitude: 29.873888, Name: "Rwanda"},
		{Latitude: 23.885942, Longitude: 45.079162, Name: "Saudi Arabia"},
		{Latitude: -9.64571, Longitude: 160.156194, Name: "Solomon Islands"},
		{Latitude: -4.679574, Longitude: 55.491977, Name: "Seychelles"},
		{Latitude: 12.862807, Longitude: 30.217636, Name: "Sudan"},
		{Latitude: 60.128161, Longitude: 18.643501, Name: "Sweden"},
		{Latitude: 1.352083, Longitude: 103.819836, Name: "Singapore"},
		{Latitude: -24.143474, Longitude: -10.030696, Name: "Saint Helena"},
		{Latitude: 46.151241, Longitude: 14.995463, Name: "Slovenia"},
		{Latitude: 77.553604, Longitude: 23.670272, Name: "Svalbard and Jan Mayen"},
		{Latitude: 48.669026, Longitude: 19.699024, Name: "Slovakia"},
		{Latitude: 8.460555, Longitude: -11.779889, Name: "Sierra Leone"},
		{Latitude: 43.94236, Longitude: 12.457777, Name: "San Marino"},
		{Latitude: 14.497401, Longitude: -14.452362, Name: "Senegal"},
		{Latitude: 5.152149, Longitude: 46.199616, Name: "Somalia"},
		{Latitude: 3.919305, Longitude: -56.027783, Name: "Suriname"},
		{Latitude: 0.18636, Longitude: 6.613081, Name: "São Tomé and Príncipe"},
		{Latitude: 13.794185, Longitude: -88.89653, Name: "El Salvador"},
		{Latitude: 34.802075, Longitude: 38.996815, Name: "Syria"},
		{Latitude: -26.522503, Longitude: 31.465866, Name: "Swaziland"},
		{Latitude: 21.694025, Longitude: -71.797928, Name: "Turks and Caicos Islands"},
		{Latitude: 15.454166, Longitude: 18.732207, Name: "Chad"},
		{Latitude: -49.280366, Longitude: 69.348557, Name: "French Southern Territories"},
		{Latitude: 8.619543, Longitude: 0.824782, Name: "Togo"},
		{Latitude: 15.870032, Longitude: 100.992541, Name: "Thailand"},
		{Latitude: 38.861034, Longitude: 71.276093, Name: "Tajikistan"},
		{Latitude: -8.967363, Longitude: -171.855881, Name: "Tokelau"},
		{Latitude: -8.874217, Longitude: 125.727539, Name: "Timor-Leste"},
		{Latitude: 38.969719, Longitude: 59.556278, Name: "Turkmenistan"},
		{Latitude: 33.886917, Longitude: 9.537499, Name: "Tunisia"},
		{Latitude: -21.178986, Longitude: -175.198242, Name: "Tonga"},
		{Latitude: 38.963745, Longitude: 35.243322, Name: "Turkey"},
		{Latitude: 10.691803, Longitude: -61.222503, Name: "Trinidad and Tobago"},
		{Latitude: -7.109535, Longitude: 177.64933, Name: "Tuvalu"},
		{Latitude: 23.69781, Longitude: 120.960515, Name: "Taiwan"},
		{Latitude: -6.369028, Longitude: 34.888822, Name: "Tanzania"},
		{Latitude: 48.379433, Longitude: 31.16558, Name: "Ukraine"},
		{Latitude: 1.373333, Longitude: 32.290275, Name: "Uganda"},
		{Latitude: 37.09024, Longitude: -95.712891, Name: "United States"},
		{Latitude: -32.522779, Longitude: -55.765835, Name: "Uruguay"},
		{Latitude: 41.377491, Longitude: 64.585262, Name: "Uzbekistan"},
		{Latitude: 41.902916, Longitude: 12.453389, Name: "Vatican City"},
		{Latitude: 12.984305, Longitude: -61.287228, Name: "Saint Vincent and the Grenadines"},
		{Latitude: 6.42375, Longitude: -66.58973, Name: "Venezuela"},
		{Latitude: 18.420695, Longitude: -64.639968, Name: "British Virgin Islands"},
		{Latitude: 18.335765, Longitude: -64.896335, Name: "U.S. Virgin Islands"},
		{Latitude: 14.058324, Longitude: 108.277199, Name: "Vietnam"},
		{Latitude: -15.376706, Longitude: 166.959158, Name: "Vanuatu"},
		{Latitude: -13.768752, Longitude: -177.156097, Name: "Wallis and Futuna"},
		{Latitude: -13.759029, Longitude: -172.104629, Name: "Samoa"},
		{Latitude: 42.602636, Longitude: 20.902977, Name: "Kosovo"},
		{Latitude: 15.552727, Longitude: 48.516388, Name: "Yemen"},
		{Latitude: -12.8275, Longitude: 45.166244, Name: "Mayotte"},
		{Latitude: -30.559482, Longitude: 22.937506, Name: "South Africa"},
		{Latitude: -13.133897, Longitude: 27.849332, Name: "Zambia"},
		{Latitude: -19.015438, Longitude: 29.154857, Name: "Zimbabwe"},
	}

	for _, country := range countries {
		db.Create(&country)
	}
}

func SeedUser(db *gorm.DB) {
	users := []model.User{
		{
			AccountName: "nico",
			ProfileName: "nico tanzil",
			RealName:    "Nico Tanzil",
			Email:       "tanzilclementius@gmail.com",
			Password:    providers.HashPassword("password"),
			Points:      15000,
			Balance:     1000000,
			CustomURL:   "nico",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
				{ID: 4},
			},
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			MiniProfileBackgroundID: 1,
			AvatarFrameID:           1,
			ThemeID:                 1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},

			CountryID: 99,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			Experience: 550,

			Badges: []*model.Badge{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			FeaturedBadge: &model.Badge{ID: 3},
		},
		{
			AccountName: "william",
			ProfileName: "William",
			RealName:    "William Martin",
			Email:       "will@mail.com",
			Password:    providers.HashPassword("password"),
			Balance:     100,
			Points:      15000,
			CustomURL:   "william",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 3},
			},
			MiniProfileBackgroundID: 1,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
			},
			AvatarFrameID: 1,
			ThemeID:       1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},

			CountryID: 99,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			Friends: []*model.User{
				{ID: 1},
			},
			FeaturedBadge: &model.Badge{ID: 1},
			Badges: []*model.Badge{
				{ID: 1}, {ID: 2},
			},
		},
		{
			AccountName: "ricko",
			ProfileName: "Ricko",
			RealName:    "Ricko Adrio",
			Email:       "rick@mail.com",
			Password:    providers.HashPassword("password"),
			Balance:     250000,
			CustomURL:   "rick",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 3},
			},
			MiniProfileBackgroundID: 3,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
			},
			AvatarFrameID: 1,
			ThemeID:       1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},
			CountryID:     125,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			IsSuspend:     true,
			FeaturedBadge: &model.Badge{ID: 1},
			Badges: []*model.Badge{
				{ID: 1}, {ID: 2},
			},
		},
		{
			AccountName: "leo",
			ProfileName: "Leo",
			RealName:    "Leonardo",
			Email:       "leo@mail.com",
			Password:    providers.HashPassword("password"),
			Balance:     150000,
			CustomURL:   "leo",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 3},
			},
			MiniProfileBackgroundID: 2,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
			},
			AvatarFrameID: 1,
			ThemeID:       1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},
			Friends: []*model.User{
				{ID: 1},
			},
			CountryID:     35,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			IsSuspend:     false,
			FeaturedBadge: &model.Badge{ID: 2},
			Badges: []*model.Badge{
				{ID: 1}, {ID: 2},
			},
		},
		{
			AccountName: "elisa",
			ProfileName: "Elisa",
			RealName:    "Elisa Pat",
			Email:       "elisa@mail.com",
			Password:    providers.HashPassword("password"),
			Balance:     150000,
			CustomURL:   "elisa",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 3},
			},
			MiniProfileBackgroundID: 1,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
			},
			AvatarFrameID: 1,
			ThemeID:       1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},
			CountryID:     115,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			IsSuspend:     false,
			FeaturedBadge: &model.Badge{ID: 2},
			Badges: []*model.Badge{
				{ID: 1}, {ID: 2},
			},
		},
		{
			AccountName: "nando",
			ProfileName: "Fernando",
			RealName:    "Fernando FM",
			Email:       "nando@mail.com",
			Password:    providers.HashPassword("password"),
			Balance:     150000,
			CustomURL:   "nando",
			Summary:     "No information given.",
			Avatar:      firebase_data.Avatar,

			ProfileBackgrounds: []*model.ProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			ProfileBackgroundID: 1,
			MiniProfileBackgrounds: []*model.MiniProfileBackground{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			MiniProfileBackgroundID: 2,
			AvatarFrames: []*model.AvatarFrame{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			AvatarFrameID: 2,
			ThemeID:       1,
			AnimatedAvatars: []*model.AnimatedAvatar{
				{ID: 1},
			},
			ChatStickers: []*model.ChatSticker{
				{ID: 1},
			},
			CountryID:     75,
			Games: []*model.Game{
				{ID: 1}, {ID: 2}, {ID: 6}, {ID: 7},
			},
			IsSuspend:     false,
			FeaturedBadge: &model.Badge{ID: 1},
			Badges: []*model.Badge{
				{ID: 1}, {ID: 2},
			},
		},
	}

	for _, user := range users {
		db.Create(&user)
	}

	var user model.User
	db.First(&user, "id = ?", 1)
	user.Friends = []*model.User{
		{ID: 2},
		{ID: 4},
	}
	db.Save(&user)
}

func SeedFriendRequest(db *gorm.DB) {
	friendRequests := []model.FriendRequest{
		{
			RequesterID: 3,
			RequestedID: 1,
			Status:      "Pending",
		},
		{
			RequesterID: 1,
			RequestedID: 5,
			Status:      "Ignored",
		},
		{
			RequesterID: 6,
			RequestedID: 1,
			Status:      "Pending",
		},
	}

	for _, report := range friendRequests {
		db.Create(&report)
	}
}

func SeedUserReport(db *gorm.DB) {
	userReports := []model.UserReport{
		{
			ReportedId:  1,
			ReporterId:  2,
			Description: "Toxic user!",
		},
		{
			ReportedId:  1,
			ReporterId:  3,
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
			Description: "Please unsuspend me!",
			User:        &model.User{ID: 3},
		},
	}

	for _, request := range suspensionRequests {
		db.Create(&request)
	}
}

func SeedCart(db *gorm.DB) {
	carts := []model.Cart{
		{
			User: &model.User{ID: 1},
			Game: &model.Game{ID: 2},
		},
		{
			User: &model.User{ID: 1},
			Game: &model.Game{ID: 3},
		},
	}

	for _, cart := range carts {
		db.Create(&cart)
	}
}

func SeedWishlist(db *gorm.DB) {
	wishlists := []model.Wishlist{
		{
			User: &model.User{ID: 1},
			Game: &model.Game{ID: 5},
		},
	}

	for _, wishlist := range wishlists {
		db.Create(&wishlist)
	}
}

func SeedCommunityArtPost(db *gorm.DB) {
	communityArtPosts := []model.CommunityArtPost{
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F1.jpg?alt=media&token=c4636eb2-e642-4be2-9566-7460a1bb0d14",
			Description: "Sunset",
			IsImage:     true,
			User:        &model.User{ID: 1},
			Like:        5,
		},
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F2.jpg?alt=media&token=61260093-4558-4582-810c-4509b18bacea",
			Description: "Nature",
			IsImage:     true,
			User:        &model.User{ID: 1},
			Like:        -3,
		},
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F3.jpg?alt=media&token=878e2002-1684-4f64-a3e7-ffadd4656514",
			Description: "Holy tree",
			IsImage:     true,
			User:        &model.User{ID: 1},
			Like:        12,
		},
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F4.jpg?alt=media&token=65c255e5-9d45-4516-b1c2-3f05f98f11db",
			Description: "Lonely Road",
			IsImage:     true,
			User:        &model.User{ID: 2},
			Like:        50,
		},
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F1.mp4?alt=media&token=cc3bfb4c-1b7e-46ff-9ee4-389144321d1a",
			Description: "Video 1",
			IsImage:     false,
			User:        &model.User{ID: 2},
			Like:        25,
		},
		{
			Link:        "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fcommunity-art-post%2F2.mp4?alt=media&token=191ba8ae-cb0c-4da6-8ddb-153b6c92aab9",
			Description: "Video 2",
			IsImage:     false,
			User:        &model.User{ID: 3},
			Like:        0,
		},
	}

	for _, communityArtPost := range communityArtPosts {
		db.Create(&communityArtPost)
	}
}

func SeedCommunityArtPostReview(db *gorm.DB) {
	communityArtPostReviews := []model.CommunityArtPostReview{
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 3},
			Comment: "Cool!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 2},
			User:    &model.User{ID: 2},
			Comment: "hehe!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro 1!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro 2!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro 3!",
		}, {
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro 4!",
		},
		{
			Post:    &model.CommunityArtPost{ID: 1},
			User:    &model.User{ID: 2},
			Comment: "Cool artwork bro 5!",
		},
	}

	for _, communityArtPostReview := range communityArtPostReviews {
		db.Create(&communityArtPostReview)
	}
}

func SeedCommunityGameReview(db *gorm.DB) {
	communityGameReviews := []model.CommunityGameReview{
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Don't play this game",
			User:          &model.User{ID: 2},
			Game:          &model.Game{ID: 1},
			IsRecommended: false,
			HelpfulCount:  73,
		},
		{
			Description:   "Really fun game 3",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 2},
			IsRecommended: true,
			HelpfulCount:  100,
		},
		{
			Description:   "Really fun game 4",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 3},
			IsRecommended: true,
			HelpfulCount:  126,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 2},
			Game:          &model.Game{ID: 2},
			IsRecommended: false,
			HelpfulCount:  67,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 3},
			Game:          &model.Game{ID: 3},
			IsRecommended: true,
			HelpfulCount:  56,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 6},
			IsRecommended: false,
			HelpfulCount:  23,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 7},
			IsRecommended: true,
			HelpfulCount:  456,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 4},
			IsRecommended: true,
			HelpfulCount:  45,
		}, {
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 4},
			IsRecommended: true,
			HelpfulCount:  45,
		}, {
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 4},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 4},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		}, {
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		}, {
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		}, {
			Description:   "Really fun game",
			User:          &model.User{ID: 1},
			Game:          &model.Game{ID: 1},
			IsRecommended: true,
			HelpfulCount:  45,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 3},
			Game:          &model.Game{ID: 3},
			IsRecommended: true,
			HelpfulCount:  56,
		},
		{
			Description:   "Really fun game",
			User:          &model.User{ID: 3},
			Game:          &model.Game{ID: 3},
			IsRecommended: true,
			HelpfulCount:  56,
		},
	}

	for _, communityGameReview := range communityGameReviews {
		db.Create(&communityGameReview)
	}
}

func SeedCommunityGameReviewComment(db *gorm.DB) {
	communityGameReviewComments := []model.CommunityGameReviewComment{
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 2",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 3",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 4",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 5",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 6",
		},
		{
			User:        &model.User{ID: 2},
			Review:      &model.CommunityGameReview{ID: 1},
			Description: "great review I find this really helpful 7",
		},
	}

	for _, communityGameReviewComment := range communityGameReviewComments {
		db.Create(&communityGameReviewComment)
	}
}

func SeedGameDiscussion(db *gorm.DB) {
	gameDiscussions := []model.GameDiscussion{
		{
			Title:       "How to report players in Dota 2 ",
			Description: "Some users on Dota 2 community discussions have misconceptions about how the report system works. Here is a guide on how to use the report feature and other relevant information about the report system.\n",
			Game:        &model.Game{ID: 1},
			User:        &model.User{ID: 1},
		},
		{
			Title:       "HELP! - Q: \"Is this site a scam?\" A: Probably! ",
			Description: "First of all, compliments for checking this thread and being vigilant.\n\nThe short answer is: Until proven otherwise, assume the website is indeed a scam.\n\nThere are many more scam sites out there than legitimate sites and it's impossible for us to list them all.\n\nWhile many scam sites are blocked (shown as {LINK REMOVED} ) it's easy to make a new site and spam this, especially when often the accounts hijacked are used to spread even more links, or if certain sites use referral programs to trick you into spreading the website even more.\n\nScam sites are often professionally built, have a convincing proposition and are well operated.\nThey also come in many forms and continuously evolve, as a result of various steps taken to protect you from these sites.",
			Game:        &model.Game{ID: 1},
			User:        &model.User{ID: 2},
		},
		{
			Title:       "Trading is NOT allowed in this sub-forum.",
			Description: "Why isn’t it allowed?\n\nIt's simple.\n\nFirstly, there is a dedicated Trading Sub Forum available.\n\nSecondly, there are thousands of trade topics created in the trading sub forum each day. The disruption caused by a large of them spilling over into this forum would be unacceptable.\n\nFinally, there are some users simply attempting to game the system. Trade topics are commonly duplicated in the trading sub forum and this one. That just isn’t fair to users doing the right thing keeping their trade topics in the right place.",
			Game:        &model.Game{ID: 1},
			User:        &model.User{ID: 1},
		},
		{
			Title:       "Terraria Bugs ( Post Here! ) ",
			Description: "Edit in 2021: Thank you to the people who have given awards, very kind of you! I hope this thread has helped keep the discussions a cleaner place, and has helped the community come together in order to solve eachothers problems. Remember; if you've solved your problem, there may be tens or hundreds of others that have that same problem! Drop a comment and help them solve it too! :) You can search your problem with keywords via the discussions search once someone has posted about it. Stay safe everyone ❤️\n\nPotential resolution save fix\n1) If in fullscreen press alt-enter to go to windowed mode\n2) Change the resolution to the desired setting and DO NOT PRESS APPLY\n3) Press alt-enter\nshould be fixed\n\n- Game crashes after 5ish minutes on multiplayer server, (lost connection). Allows immediate reconnection. Windows 8.1 - Also noted to be happening with Windows XP 32bit. Windows 7 Home Premium\n\n- Worlds have been deleted / Unaccessable after 1.3 update. (In-Game, Worlds still show in the game folder when accessed through computer.)\n\n- Players spawning into a new world with brand new character and recieving 2 achievements ( Heart fruit and Mana Star upgrade).\n\n- Problems changing the resolution from 1440x900 to 800x600.",
			Game:        &model.Game{ID: 6},
			User:        &model.User{ID: 1},
		},
		{
			Title:       "Terraria 1.4.1 Hotfixes! Information inside! (updated as hotfixes release) ",
			Description: "Hello everyone!\n\nAs with any release, new bugs will arise - so this will be your one-stop shop for the changelogs that are handed down with each Terraria 1.4.1 hotfix... the first of which went live just a few moments ago!\n\n",
			Game:        &model.Game{ID: 6},
			User:        &model.User{ID: 2},
		},
		{
			Title:       "How to make a server!!! ",
			Description: "Will look into updating this more when I have time. I'm a busy person so if yall could just chill the F out I would be greatful. Have a nice day.\n\n\nThank you all for the warm feed back and I'm happy that this thread has been helpful for so many of you. From now on please refrain from adding me I check this semi regularly and there are also plenty of other people that lurk here with helpful advice.\n\nNOTE: If I do actually take the time to respond/add/help you be respectful or I will either ignore you or block you. I'm taking my own personal time to do these things and it is a choice not and resposibility.\n\nFor simple knowledge of in game things your best friend is>>> http://wiki.terrariaonline.com/Terraria_Wiki\nEvery pro uses this I keep it open 24/7 while I play.\n\n\n\nSide note to all newbie server makers, NO your IP is not 127.0.0.1, nor will it be 192.168.0.0-200, these are internal IP addresses and are useless to anyone else. Go to www.whatismyip.com and THAT is your ip.\n",
			Game:        &model.Game{ID: 6},
			User:        &model.User{ID: 3},
		},
		{
			Title:       "Help! Is this site a scam? Probably. ",
			Description: "First of all, compliments for being vigilant and applying some critical thinking. Till we figure out if this is a legitimate site not trying to scam you, we should assume they are for all intents and purposes.\n\nThere are many more scam sites out there than legitimate sites and it's impossible for us to list them all.\n\nWhile we do block poor sites (shown as {LINK REMOVED} ) it's easy to make a new site and spam this, especially when often the accounts hijacked are used to spread even more links, or if certain sites use referral programs to trick you into spreading your site even more.\n\nScam sites are often professionally built, have a convincing proposition and are well operated.\nThey also come in many forms and continuously evolve, as a result of various steps taken to protect you from these sites.\n",
			Game:        &model.Game{ID: 4},
			User:        &model.User{ID: 1},
		},
		{
			Title:       "Mythbuster: The unofficial FAQ ",
			Description: "As people who visited the forums frequently have undoubtedly noticed, there are some recurring themes that often lead different people to make the same threads.\n\nIn some cases these people are simply poorly informed, in some cases it's persisting myths, and other times it's simply a case of not realising a search function and rules exist.\n\nWe plan on posting each and any of these topics as individual, open threads so that people can discussions there in and link to them.\n",
			Game:        &model.Game{ID: 4},
			User:        &model.User{ID: 1},
		},
		{
			Title:       "Trading is just not allowed in this forum. ",
			Description: "It's simple.\n\nFirstly, there is a dedicated Trading Sub Forum available.\n\nSecondly, there are thousands of trade topics created in the trading sub forum each day. The disruption caused by a large of them spilling over into this forum would be unacceptable.\n\nFinally, there are some users simply attempting to game the system. Trade topics are commonly duplicated in the trading sub forum and this one. That just isn’t fair to users doing the right thing keeping their trade topics in the right place.",
			Game:        &model.Game{ID: 4},
			User:        &model.User{ID: 2},
		},
	}

	for _, gameDiscussion := range gameDiscussions {
		db.Create(&gameDiscussion)
	}
}

func SeedGameDiscussionReply(db *gorm.DB) {
	gameDiscussionReplies := []model.GameDiscussionReply{
		{
			Description: "nice thread! very useful and informative",
			Discussion:  &model.GameDiscussion{ID: 1},
			User:        &model.User{ID: 2},
		},
		{
			Description: "nice thread! very useful and informative 2",
			Discussion:  &model.GameDiscussion{ID: 1},
			User:        &model.User{ID: 1},
		},
		{
			Description: "nice thread! very useful and informative 3",
			Discussion:  &model.GameDiscussion{ID: 1},
			User:        &model.User{ID: 2},
		}, {
			Description: "nice thread! very useful and informative 4",
			Discussion:  &model.GameDiscussion{ID: 2},
			User:        &model.User{ID: 1},
		},
		{
			Description: "nice thread! very useful and informative 5",
			Discussion:  &model.GameDiscussion{ID: 2},
			User:        &model.User{ID: 3},
		},
		{
			Description: "nice thread! very useful and informative 6",
			Discussion:  &model.GameDiscussion{ID: 2},
			User:        &model.User{ID: 2},
		},
		{
			Description: "nice thread! very useful and informative 7",
			Discussion:  &model.GameDiscussion{ID: 3},
			User:        &model.User{ID: 1},
		},
		{
			Description: "nice thread! very useful and informative 8",
			Discussion:  &model.GameDiscussion{ID: 3},
			User:        &model.User{ID: 2},
		},
		{
			Description: "nice thread! very useful and informative 9",
			Discussion:  &model.GameDiscussion{ID: 3},
			User:        &model.User{ID: 3},
		},
		{
			Description: "nice thread! very useful and informative 10",
			Discussion:  &model.GameDiscussion{ID: 4},
			User:        &model.User{ID: 1},
		}, {
			Description: "nice thread! very useful and informative 11",
			Discussion:  &model.GameDiscussion{ID: 4},
			User:        &model.User{ID: 2},
		},
	}

	for _, gameDiscussionReply := range gameDiscussionReplies {
		db.Create(&gameDiscussionReply)
	}
}

func SeedPaymentType(db *gorm.DB) {
	paymentTypes := []model.PaymentType{
		{Name: "My Staem Wallet"},
		{Name: "Visa"},
		{Name: "MasterCard"},
		{Name: "eClub Points"},
	}

	for _, paymentType := range paymentTypes {
		db.Create(&paymentType)
	}
}

func SeedTransactionHeader(db *gorm.DB) {
	transactionHeaders := []*model.TransactionHeader{
		{
			PaymentType: &model.PaymentType{ID: 1},
			Sender:      &model.User{ID: 1},
			Receiver:    &model.User{ID: 2},
			Total:       150000,
		},
		{
			PaymentType: &model.PaymentType{ID: 1},
			Sender:      &model.User{ID: 2},
			Receiver:    &model.User{ID: 3},
			Total:       150000,
		},
		{
			PaymentType: &model.PaymentType{ID: 1},
			Sender:      &model.User{ID: 1},
			Receiver:    &model.User{ID: 1},
			Total:       150000,
		},
		{
			PaymentType: &model.PaymentType{ID: 1},
			Sender:      &model.User{ID: 2},
			Receiver:    &model.User{ID: 2},
			Total:       150000,
		},
		{
			PaymentType: &model.PaymentType{ID: 1},
			Sender:      &model.User{ID: 3},
			Receiver:    &model.User{ID: 2},
			Total:       150000,
		},
	}

	for _, transactionHeader := range transactionHeaders {
		db.Create(&transactionHeader)
	}
}

func SeedTransactionDetail(db *gorm.DB) {
	transactions := []*model.TransactionDetail{
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 3},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 4},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 5},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 1},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 2},
		}, {
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 3},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 1},
			Game:              &model.Game{ID: 6},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 2},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 3},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 4},
			Game:              &model.Game{ID: 2},
		},
		{
			TransactionHeader: &model.TransactionHeader{ID: 5},
			Game:              &model.Game{ID: 2},
		},
	}

	for _, transaction := range transactions {
		db.Create(&transaction)
	}
}

func SeedGift(db *gorm.DB) {
	gifts := []*model.Gift{
		{
			Sender:     &model.User{ID: 2},
			Receiver:   &model.User{ID: 1},
			FirstName:  "hello",
			Message:    "this is a gift",
			Sentiment:  "XOXO",
			Signature:  "sig",
			IsComplete: true,
			IsOpen:     false,
		},
	}

	for _, gift := range gifts {
		db.Create(&gift)
	}
}

func SeedItemType(db *gorm.DB) {
	itemTypes := []model.ItemType{
		{
			Name:    "Profane Union",
			Summary: "Already an unnatural mingling of two separate minds inhabiting a single form, N'aix finds nothing distasteful in augmenting that form ever further to suit their twisted needs. ",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2Fprofane%20union.png?alt=media&token=84fdcdcb-c6e7-435f-a432-4eba978f1814",
			Game:    &model.Game{ID: 1},
		},
		{
			Name:    "Profane Union 123",
			Summary: "Already an unnatural mingling of two separate minds inhabiting a single form, N'aix finds nothing distasteful in augmenting that form ever further to suit their twisted needs. ",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2Fprofane%20union.png?alt=media&token=84fdcdcb-c6e7-435f-a432-4eba978f1814",
			Game:    &model.Game{ID: 1},
		},
		{
			Name:    "Razor Spines of the Sunken Gaoler",
			Summary: "During his long-ago tenure as a Dark Reef Enforcer, Slardar sharpened his spines breaking up escape attempts and quelling constant insurrections.",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2FRazor%20Spines%20of%20the%20Sunken%20Gaoler.png?alt=media&token=cc0c7647-cfe1-4d31-bb7a-02c9aa0eafbb",
			Game:    &model.Game{ID: 1},
		},
		{
			Name:    "Plates of the Sunken Gaoler",
			Summary: "During his long-ago tenure as a Dark Reef Enforcer, Slardar sharpened his spines breaking up escape attempts and quelling constant insurrections.",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2FBlastforge%20Exhaler.png?alt=media&token=3a886e22-77b0-43c2-86cc-221a14fc969e",
			Game:    &model.Game{ID: 1},
		},
		{
			Name:    "Lost Shield",
			Summary: "Carved from the oldest trees of the thousand-league wood, this shield will protect old bones from attack.",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2FLost%20Hills%20Shield.png?alt=media&token=b7281274-4911-4800-9e84-13ec358ae170",
			Game:    &model.Game{ID: 2},
		},
		{
			Name:    "The International 2017 Emoticon Pack II",
			Summary: "Unlocks 7 emoticons earned via The International Battle Pass 2017.",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2FLost%20Hills%20Shield.png?alt=media&token=b7281274-4911-4800-9e84-13ec358ae170",
			Game:    &model.Game{ID: 3},
		},
		{
			Name:    "Yulsaria's Mantle",
			Summary: "In an age now lost to time, the Frozen Witch Yulsaria ruled the whitelands, summoning blizzards and hail storms upon those who displeased her, while an army of ice golems roamed the lands to snuff out all warmth. In time, her southward expansion angered the Eldwurm Slyrak who, in his terrible rage, melted Yulsaria's armies with his endless flame before conquering the Frozen Witch herself. Now, centuries later, shifts in the ice have uncovered yet another shard of her empire: her frosty mantle. ",
			Link:    "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fitems%2FYulsaria's%20Mantle.png?alt=media&token=83ce7f54-1001-4174-8605-114c8a1ee717",
			Game:    &model.Game{ID: 4},
		},
	}

	for _, itemType := range itemTypes {
		db.Create(&itemType)
	}
}

func SeedItem(db *gorm.DB) {
	items := []model.Item{
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 2},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 2},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 2},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 3},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 3},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 3},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 4},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 4},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 4},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 5},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 6},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 6},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 7},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 7},
			User:     &model.User{ID: 3},
		},
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, item := range items {
		db.Create(&item)
	}
	for _, item := range items {
		db.Create(&item)
	}
}

func SeedBid(db *gorm.DB) {
	bids := []model.Bid{
		//{
		//	ItemType: &model.ItemType{ID: 1},
		//	User:     &model.User{ID: 1},
		//},
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 2},
		},
		{
			ItemType: &model.ItemType{ID: 1},
			User:     &model.User{ID: 3},
		},
		{
			ItemType: &model.ItemType{ID: 2},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 3},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 4},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 5},
			User:     &model.User{ID: 1},
		},
		{
			ItemType: &model.ItemType{ID: 7},
			User:     &model.User{ID: 1},
		},
	}

	for _, bid := range bids {
		db.Create(&bid)
	}
	for _, bid := range bids {
		db.Create(&bid)
	}
	for _, bid := range bids {
		db.Create(&bid)
	}
	for _, bid := range bids {
		db.Create(&bid)
	}
	for _, bid := range bids {
		db.Create(&bid)
	}
	for _, bid := range bids {
		db.Create(&bid)
	}
}

func SeedItemTransaction(db *gorm.DB) {
	transactions := []model.ItemTransaction{
		{
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 7),
		},
		{
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     235000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 6),
		},
		{
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     175000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 5),
		}, {
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     350000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		}, {
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     375000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 3),
		}, {
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     200000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 2),
		},
		{
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     250000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 1),
		},
		{
			Item:      &model.Item{ID: 1},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     250000,
			CreatedAt: time.Now().Add(-time.Hour),
		},
		{
			Item:      &model.Item{ID: 10},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 7),
		},
		{
			Item:      &model.Item{ID: 2},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 6),
		},
		{
			Item:      &model.Item{ID: 15},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 5),
		},
		{
			Item:      &model.Item{ID: 35},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 50},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 85},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 75},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 3},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 4},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 4},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 10},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 36},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
		{
			Item:      &model.Item{ID: 54},
			Seller:    &model.User{ID: 1},
			Buyer:     &model.User{ID: 2},
			Price:     150000,
			CreatedAt: time.Now().Add(-time.Hour * 24 * 4),
		},
	}

	for _, transaction := range transactions {
		db.Create(&transaction)
	}
}

func SeedSellListing(db *gorm.DB) {
	sellListings := []model.SellListing{
		//{
		//	Item: &model.Item{ID: 1},
		//	Sell: 250,
		//},
		{
			Item: &model.Item{ID: 2},
			Sell: 350000,
		},

		{
			Item: &model.Item{ID: 3},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 4},
			Sell: 400000,
		},
		{
			Item: &model.Item{ID: 5},
			Sell: 175000,
		},
		{
			Item: &model.Item{ID: 6},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 7},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 8},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 10},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 15},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 19},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 20},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 21},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 22},
			Sell: 160000,
		},
		{
			Item: &model.Item{ID: 30},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 40},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 50},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 60},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 70},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 80},
			Sell: 150000,
		},
		{
			Item: &model.Item{ID: 90},
			Sell: 150000,
		},
	}
	for _, sellListing := range sellListings {
		db.Create(&sellListing)
	}
}

func SeedBuyListing(db *gorm.DB) {
	buyListings := []model.BuyListing{
		//{
		//	Bid: &model.Bid{ID: 1},
		//	Buy: 125,
		//},
		{
			Bid: &model.Bid{ID: 2},
			Buy: 140000,
		},
		{
			Bid: &model.Bid{ID: 3},
			Buy: 140000,
		},
		{
			Bid: &model.Bid{ID: 4},
			Buy: 100000,
		},
		{
			Bid: &model.Bid{ID: 5},
			Buy: 100000,
		},
		{
			Bid: &model.Bid{ID: 6},
			Buy: 124000,
		},
		{
			Bid: &model.Bid{ID: 7},
			Buy: 150000,
		},
		{
			Bid: &model.Bid{ID: 8},
			Buy: 110000,
		},
		{
			Bid: &model.Bid{ID: 9},
			Buy: 110000,
		},
		{
			Bid: &model.Bid{ID: 10},
			Buy: 110000,
		},
		{
			Bid: &model.Bid{ID: 11},
			Buy: 90000,
		},
		{
			Bid: &model.Bid{ID: 17},
			Buy: 100000,
		},
		{
			Bid: &model.Bid{ID: 18},
			Buy: 90000,
		},
		{
			Bid: &model.Bid{ID: 19},
			Buy: 90000,
		},
		{
			Bid: &model.Bid{ID: 25},
			Buy: 85000,
		},
		{
			Bid: &model.Bid{ID: 26},
			Buy: 85000,
		},
		{
			Bid: &model.Bid{ID: 27},
			Buy: 95000,
		},
	}

	for _, buyListing := range buyListings {
		db.Create(&buyListing)
	}
}

func SeedWalletCode(db *gorm.DB) {
	walletCodes := []model.WalletCode{
		{
			Code:    "abc",
			Balance: 15000,
		},
		{
			Code:    "qwe",
			Balance: 15000,
		},
		{
			Code:    "asd",
			Balance: 15000,
		},
	}
	for _, code := range walletCodes {
		db.Create(&code)
	}
}

func SeedNewItemNotification(db *gorm.DB) {
	newItemNotifications := []model.NewItemNotification {
		{
			Item:   &model.Item{ID: 2},
			IsOpen: false,
		},
	}

	for _, notification := range newItemNotifications {
		db.Create(&notification)
	}
}
