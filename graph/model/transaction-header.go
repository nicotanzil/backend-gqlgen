package model

import "time"

type TransactionHeader struct {
	ID                 int                  `json:"id" gorm:"primaryKey"`
	PaymentTypeID      int                  `json:"paymentTypeId"`
	PaymentType        *PaymentType         `json:"paymentType" gorm:"foreignKey:PaymentTypeID"`
	SenderID           int                  `json:"senderId"`
	Sender             *User                `json:"sender" gorm:"foreignKey:SenderID"`
	ReceiverID         int                  `json:"receiverId"`
	Receiver           *User                `json:"receiver" gorm:"foreignKey:ReceiverID"`
	Total              int                  `json:"total"`
	TransactionDetails []*TransactionDetail `json:"transactionDetails" gorm:"foreignKey:TransactionHeaderID"`
	CreatedAt          time.Time            `json:"createdAt"`
	UpdatedAt          time.Time            `json:"updatedAt"`
	DeletedAt          *time.Time           `json:"deletedAt"`
}