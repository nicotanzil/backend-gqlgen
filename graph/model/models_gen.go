// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type TopReviewGame struct {
	GameID       int    `json:"game_id"`
	GameName     string `json:"game_name"`
	GamePrice    int    `json:"game_price"`
	GameBanner   string `json:"game_banner"`
	GameDiscount int    `json:"game_discount"`
	ReviewCount  int    `json:"review_count"`
}

type TopSellerGame struct {
	GameID        int    `json:"game_id"`
	GameName      string `json:"game_name"`
	GamePrice     int    `json:"game_price"`
	GameBanner    string `json:"game_banner"`
	GameDiscount  int    `json:"game_discount"`
	PurchaseCount int    `json:"purchase_count"`
}
