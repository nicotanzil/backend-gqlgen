package model

type Item struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	ItemTypeID int       `json:"itemTypeId"`
	ItemType   *ItemType `json:"itemType" gorm:"foreignKey:ItemTypeID"`
	UserID     int       `json:"userId"`
	User       *User     `json:"user" gorm:"foreignKey:UserID"`
}
