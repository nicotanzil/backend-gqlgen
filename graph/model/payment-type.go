package model

type PaymentType struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}