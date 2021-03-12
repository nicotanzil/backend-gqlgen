package model

type ItemType struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	Link    string  `json:"link"`
	GameID  int     `json:"gameId"`
	Game    *Game   `json:"game" gorm:"foreignKey:GameID"`
	Items   []*Item `json:"items" gorm:"foreignKey:ItemTypeID"`
}
