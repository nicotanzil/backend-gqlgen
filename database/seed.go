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
	SeedTag(db)
	SeedSystem(db)
	SeedReview(db)
	SeedReviewVote(db)
	SeedPublisher(db)
	SeedGenre(db)
	SeedDeveloper(db)
	SeedGame(db)
	SeedCountry(db)
	SeedUser(db)

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
			OnSale:             false,
			DiscountPercentage: 0,
			Developers: []*model.Developer {
				{
					ID:        1,
				},
			},
			GamePlayHour:       1500,
			PublisherID:        1,
			SystemID:           1,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F1%2Fbanner.jpg?alt=media&token=6fe38ac5-d74e-45f9-b1f3-5052403edfec",
		},
		{
			Name:               "STAR WARS Jedi: Fallen Order",
			Description:        "A galaxy-spanning adventure awaits in Star Wars Jedi: Fallen Order, a new third-person action-adventure title from Respawn Entertainment. This narratively driven, single-player game puts you in the role of a Jedi Padawan who narrowly escaped the purge of Order 66 following the events of Episode 3: Revenge of the Sith. On a quest to rebuild the Jedi Order, you must pick up the pieces of your shattered past to complete your training, develop new powerful Force abilities and master the art of the iconic lightsaber - all while staying one step ahead of the Empire and its deadly Inquisitors.",
			ReleaseDate:        time.Now(),
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
			OnSale:             true,
			DiscountPercentage: 50,
			Developers:         []*model.Developer {
				{
					ID: 2,
				},
			},
			PublisherID:        2,
			SystemID:           1,
			Users:              nil,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F2%2Fbanner.jpg?alt=media&token=f12004fa-98c7-4238-b209-7437fa761197",
			Video:              "",
			Image1:             "",
			Image2:             "",
			Image3:             "",
			Image4:             "",
		},
		{
			Name:               "Street Fighter V",
			Description:        "Experience the intensity of head-to-head battle with Street Fighter® V! Choose from 16 iconic characters, each with their own personal story and unique training challenges, then battle against friends online or offline with a robust variety of match options.\n\nEarn Fight Money in Ranked Matches, play for fun in Casual Matches or invite friends into a Battle Lounge and see who comes out on top! PlayStation 4 and Steam players can also play against each other thanks to cross-play compatibility!\n",
			ReleaseDate:        time.Now(),
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
			OnSale:             false,
			DiscountPercentage: 0,
			Developers:         []*model.Developer {
				{
					ID: 3,
				},
			},
			PublisherID:        3,
			SystemID:           1,
			Users:              nil,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F3%2Fbanner.jpg?alt=media&token=073881f4-a93e-4a05-9bc9-9efbb1401108",
			Video:              "",
			Image1:             "",
			Image2:             "",
			Image3:             "",
			Image4:             "",
		},
		{
			Name:               "Counter-Strike: Global Offensive",
			Description:        "Counter-Strike: Global Offensive (CS: GO) expands upon the team-based action gameplay that it pioneered when it was launched 19 years ago.\n\nCS: GO features new maps, characters, weapons, and game modes, and delivers updated versions of the classic CS content (de_dust2, etc.).",
			ReleaseDate:        time.Now(),
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
			OnSale:             false,
			DiscountPercentage: 0,
			Developers:         []*model.Developer {
				{
					ID: 1,
				},
				{
					ID: 4,
				},
			},
			PublisherID:        1,
			SystemID:           1,
			Users:              nil,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F4%2Fbanner.jpg?alt=media&token=71bf9807-e661-47f4-b19e-5b12b8f8eb79",
			Video:              "",
			Image1:             "",
			Image2:             "",
			Image3:             "",
			Image4:             "",
		},
		{
			Name:               "Grand Theft Auto V",
			Description:        "When a young street hustler, a retired bank robber and a terrifying psychopath find themselves entangled with some of the most frightening and deranged elements of the criminal underworld, the U.S. government and the entertainment industry, they must pull off a series of dangerous heists to survive in a ruthless city in which they can trust nobody, least of all each other.",
			ReleaseDate:        time.Now(),
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
			OnSale:             true,
			DiscountPercentage: 25,
			Developers:         []*model.Developer {
				{
					ID: 5,
				},
			},
			PublisherID:        4,
			SystemID:           1,
			Users:              nil,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F5%2Fbanner.jpg?alt=media&token=dafa20ed-2139-45cc-8a59-38e245faac53",
			Video:              "",
			Image1:             "",
			Image2:             "",
			Image3:             "",
			Image4:             "",
		},
		{
			Name:               "Terraria",
			Description:        "Dig, Fight, Explore, Build: The very world is at your fingertips as you fight for survival, fortune, and glory. Will you delve deep into cavernous expanses in search of treasure and raw materials with which to craft ever-evolving gear, machinery, and aesthetics? Perhaps you will choose instead to seek out ever-greater foes to test your mettle in combat? Maybe you will decide to construct your own city to house the host of mysterious allies you may encounter along your travels? ",
			ReleaseDate:        time.Now(),
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
			OnSale:             false,
			DiscountPercentage: 0,
			Developers:         []*model.Developer {
				{
					ID: 6,
				},
			},
			PublisherID:        5,
			SystemID:           1,
			Users:              nil,
			Banner:             "https://firebasestorage.googleapis.com/v0/b/staem-web.appspot.com/o/assets%2Fgames%2F6%2Fbanner.jpg?alt=media&token=b403f94b-9ebc-496d-a8a6-36a34fed085e",
			Video:              "",
			Image1:             "",
			Image2:             "",
			Image3:             "",
			Image4:             "",
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
			Experience: 550,
		},
		{
			AccountName:       "William",
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
		},
		{
			AccountName:       "Ricko",
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
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
