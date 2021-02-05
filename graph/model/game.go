package model

import "time"

type Game struct {
	ID                 	int       		`json:"id" gorm:"primaryKey"`
	Name               	string       	`json:"name"`
	Description        	string       	`json:"description"`
	ReleaseDate        	time.Time    	`json:"releaseDate"`
	Genres             	[]*Genre    	`json:"genres" gorm:"many2many:game_genres;"`
	Tags               	[]*Tag    		`json:"tags" gorm:"many2many:game_tags;"`
	OriginalPrice      	float64      	`json:"originalPrice"`
	OnSale             	bool         	`json:"onSale"`
	DiscountPercentage 	int          	`json:"discountPercentage"`
	GamePlayHour       	float64      	`json:"gamePlayHour"`
	GameReviews        	[]*Review    	`json:"gameReviews" gorm:"many2many:game_reviews;"`
	Developers         	[]*Developer 	`json:"developers" gorm:"many2many:game_developers;"`
	PublisherID         int				`json:"publisherId"`
	Publisher          	*Publisher   	`json:"publisher" gorm:"foreignKey:PublisherID"`
	SystemID			int				`json:"systemId"`
	System             	*System      	`json:"system" gorm:"foreignKey:SystemID"`
	Users              	[]*User      	`json:"users" gorm:"many2many:game_users;"`

	Banner             string       `json:"banner"`
	Video              string       `json:"video"`
	Image1             string       `json:"image1"`
	Image2             string       `json:"image2"`
	Image3             string       `json:"image3"`
	Image4             string       `json:"image4"`

	CreatedAt          	time.Time    	`json:"createdAt"`
	UpdatedAt          	time.Time    	`json:"updatedAt"`
	DeletedAt          	*time.Time   	`json:"deletedAt"`
}

type NewGame struct {
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Genres             []*NewGenre     `json:"genres"`
	Tags               []*NewTag       `json:"tags"`
	OriginalPrice      float64         `json:"originalPrice"`
	OnSale             bool            `json:"onSale"`
	DiscountPercentage int             `json:"discountPercentage"`
	Developers         []*NewDeveloper `json:"developers"`
	Publisher          *NewPublisher   `json:"publisher"`
	System             *NewSystem      `json:"system"`
}
