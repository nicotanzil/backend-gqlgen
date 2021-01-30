package model

type Login struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}