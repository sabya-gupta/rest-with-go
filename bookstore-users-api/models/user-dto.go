package models

import (
	"strings"

	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string
	Email       string
	DateCreated string
	Status      string
	Password    string
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.RestBadRequestError("Email Not Valid")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.RestBadRequestError("Password Not Valid")
	}

	return nil
}
