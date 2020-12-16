package services

import (
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/models/bookstoreuser"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"
)

func CreateUser(user *bookstoreuser.User) (*bookstoreuser.User, *errors.RestError) {
	valErr := user.Validate()
	if valErr != nil {
		return nil, valErr
	}
	err := bookstoreuser.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(iD int64) (*bookstoreuser.User, *errors.RestError) {
	return bookstoreuser.GetUserById(iD)
}

func UpdateUser(user *bookstoreuser.User) *errors.RestError {
	return bookstoreuser.UpdateUser(user)
}

func DeleteUser(id int64) *errors.RestError {
	return bookstoreuser.DeleteUser(id)
}

func FindUsersByStatus(status string) (bookstoreuser.Users, *errors.RestError) {
	return bookstoreuser.FindUserByStatus(status)
}
