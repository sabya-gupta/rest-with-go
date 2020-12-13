package models

import (
	"fmt"
	"strings"

	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/dateutils"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"
	"github.com/sabya-gupta/rest-with-go/database/mysql/bookstoredb"
)

var userDB = make(map[int64]*User)

const (
	userInsQry = "INSERT INTO USERS (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
)

func GetUserById(id int64) (*User, *errors.RestError) {
	pingErr := bookstoredb.DBClient.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
	user := userDB[id]
	if user == nil {
		return nil, errors.RestNotFoundError(fmt.Sprintf("user %d not found", id))
	}
	return user, nil
}

func SaveUser(user *User) *errors.RestError {
	pingErr := bookstoredb.DBClient.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	stmnt, err1 := bookstoredb.DBClient.Prepare(userInsQry)
	if err1 != nil {
		fmt.Println(err1)
		return errors.RestInternalServerError("Cannot Create Statement")
	}
	defer stmnt.Close()

	user.DateCreated = dateutils.GetNowAsString()
	result, err2 := stmnt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err2 != nil {
		fmt.Println("The error is : ", err2.Error())
		if strings.Contains(err2.Error(), "email_UNIQUE") {
			return errors.RestBadRequestError("Email already exists")
		}
		return errors.RestInternalServerError("Cannot Create User")
	}

	userId, err3 := result.LastInsertId()
	if err3 != nil {
		fmt.Println(err3)
		return errors.RestInternalServerError("Cannot get User")
	}

	user.Id = userId
	return nil
}
