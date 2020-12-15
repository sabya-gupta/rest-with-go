package services

import (
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/models"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"
)

func CreateUser(user *models.User) (*models.User, *errors.RestError) {
	valErr := user.Validate()
	if valErr != nil {
		return nil, valErr
	}
	err := models.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(iD int64) (*models.User, *errors.RestError) {
	return models.GetUserById(iD)
}

func UpdateUser(user *models.User) *errors.RestError {
	return models.UpdateUser(user)
}

func DeleteUser(id int64) *errors.RestError {
	return models.DeleteUser(id)
}

func FindUsersByStatus(status string) ([]*models.User, *errors.RestError) {
	return models.FindUserByStatus(status)
}
