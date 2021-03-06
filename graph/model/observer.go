package model

type observer interface {
	sendEmail(string)
}