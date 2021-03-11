package model

type Item struct {
	ID      int     `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	Link    string  `json:"link"`
	GameID  int     `json:"gameId"`
	Game    *Game   `json:"game" gorm:"foreignKey:GameID"`
	Users    []*User `json:"user" gorm:"many2many:user_items;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
