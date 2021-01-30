package model

import (
	"time"
)

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
	PublisherID         int				`json:"publisher"`
	SystemID			int				`json:"system"`
	CreatedAt          	time.Time    	`json:"createdAt"`
	UpdatedAt          	time.Time    	`json:"updatedAt"`
	DeletedAt          	*time.Time   	`json:"deletedAt"`
}

