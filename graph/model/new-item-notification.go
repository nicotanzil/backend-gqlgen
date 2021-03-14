package model

type NewItemNotification struct {
	ID     int   `json:"id" gorm:"primaryKey"`
	ItemID int   `json:"itemId"`
	Item   *Item `json:"item" gorm:"foreignKey:ItemID"`
	IsOpen bool  `json:"isOpen"`
}
