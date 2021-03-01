package model

import "time"

type Game struct {
	ID                 	int       		`json:"id" gorm:"primaryKey"`
	Name               	string       	`json:"name"`
	Description        	string       	`json:"description"`
	ReleaseDate        	time.Time    	`json:"releaseDate"`
	Genres             	[]*Genre    	`json:"genres" gorm:"many2many:game_genres;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tags               	[]*Tag    		`json:"tags" gorm:"many2many:game_tags;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OriginalPrice      	float64      	`json:"originalPrice"`
	PromoID				int 			`json:"promoId"`
	Promo         		*Promo       	`json:"promo" gorm:"foreignKey:PromoID;"`
	GamePlayHour       	float64      	`json:"gamePlayHour"`
	GameReviews        	[]*Review    	`json:"gameReviews" gorm:"many2many:game_reviews;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Developers         	[]*Developer 	`json:"developers" gorm:"many2many:game_developers;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PublisherID         int				`json:"publisherId;"`
	Publisher          	*Publisher   	`json:"publisher" gorm:"foreignKey:PublisherID;"`
	SystemID			int				`json:"systemId"`
	System             	*System      	`json:"system" gorm:"foreignKey:SystemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Users              	[]*User      	`json:"users" gorm:"many2many:game_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Banner             string       `json:"banner"`
	Video              []*GameVideo `json:"video" gorm:"foreignKey:GameID"`
	Images             []*GameImage `json:"images" gorm:"foreignKey:GameID"`

	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	DeletedAt          *time.Time    `json:"deletedAt"`
}

type NewGame struct {
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Genres        []*InputGenre     `json:"genres"`
	Tags          []*InputTag       `json:"tags"`
	OriginalPrice float64           `json:"originalPrice"`
	Promo         *InputPromo       `json:"promo"`
	Developers    []*InputDeveloper `json:"developers"`
	PublisherID   int               `json:"publisherId"`
	SystemID      int               `json:"systemId"`
}