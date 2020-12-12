package models

type User struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string
	Email       string
	DateCreated string
}
