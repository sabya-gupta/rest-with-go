package models

import (
	"fmt"

	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/errors"
)

var userDB = make(map[int64]*User)

func GetUserById(id int64) (*User, *errors.RestError) {
	user := userDB[id]
	if user == nil {
		return nil, errors.RestNotFoundError(fmt.Sprintf("user %d not found", id))
	}
	return user, nil
}

func SaveUser(user *User) *errors.RestError {
	if userDB[user.Id] != nil {
		return errors.RestBadRequestError(fmt.Sprintf("user %d is already present", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
