package bookstoreuser

import (
	"encoding/json"
	"fmt"
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

type PublicUser struct {
	Id          int64 `json:"Id"`
	DateCreated string
	Status      string
	// Password    string
}

type PrivateUser struct {
	Id          int64  `json:"Id"`
	FirstName   string `json:"FirstName"`
	LastName    string
	Email       string
	DateCreated string
	Status      string
}

type Users []*User

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

func (user *User) Marshall(isPublic bool) interface{} {
	fmt.Println("isPublic ? ", isPublic)
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	//another way if Json structures are same!
	userJson, _ := json.Marshal(user)
	var pvtUser PrivateUser

	json.Unmarshal(userJson, &pvtUser)

	return pvtUser
}

func (users Users) Marshall(isPublic bool) []interface{} {
	retUsers := make([]interface{}, len(users))
	for idx, user := range users {
		retUsers[idx] = user.Marshall(isPublic)
	}
	return retUsers
}
