package model

type TransactionDetail struct {
	TransactionHeaderID int                `json:"transactionHeaderId"`
	TransactionHeader   *TransactionHeader `json:"transactionHeader" gorm:"foreignKey:TransactionHeaderID"`
	GameID              int                `json:"gameId"`
	Game                *Game              `json:"game" gorm:"foreignKey:GameID"`
}
