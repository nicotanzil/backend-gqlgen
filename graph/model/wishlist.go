package model

type Wishlist struct {

	UserID int   `json:"userId" gorm:"primaryKey"`
	User   *User `json:"user" gorm:"foreignKey:UserID"`
	GameID int   `json:"gameId" gorm:"primaryKey"`
	Game   *Game `json:"game" gorm:"foreignKey:GameID"`
	observerList []observer
}
