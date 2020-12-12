package services

import (
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/errors"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/models"
)

func CreateUser(user *models.User) (*models.User, *errors.RestError) {
	valErr := user.Validate()
	if valErr != nil {
		return nil, valErr
	}
	return user, nil
}
