package models

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/dateutils"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"
	"github.com/sabya-gupta/rest-with-go/database/mysql/bookstoredb"
)

// var userDB = make(map[int64]*User)

const (
	userInsQry    = "INSERT INTO USERS (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	emailUnique   = "email_UNIQUE"
	userGetByID   = "SELECT id, first_name, last_name, email, date_created FROM USERS WHERE id = ?;"
	userUpdtQry   = "UPDATE USERS SET first_name = ? , last_name = ?, email = ?  WHERE id = ?;"
	norows        = "no rows in result set"
	userDeleteQry = "DELETE FROM USERS WHERE id=?"
)

func GetUserById(id int64) (*User, *errors.RestError) {
	pingErr := bookstoredb.DBClient.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	stmnt, err1 := bookstoredb.DBClient.Prepare(userGetByID)
	if err1 != nil {
		fmt.Println(err1)
		return nil, errors.RestInternalServerError("Cannot Create Prepare Get Statement")
	}
	defer stmnt.Close()
	usrRow := stmnt.QueryRow(id) //if query then you need to close the userRow
	var retUser User
	err2 := usrRow.Scan(&retUser.Id, &retUser.FirstName, &retUser.LastName, &retUser.Email, &retUser.DateCreated)
	if err2 != nil {
		fmt.Println(err2.Error())
		sqlErr, ok := err2.(*mysql.MySQLError)
		fmt.Println(">>>", sqlErr, ok)

		if !ok {
			return nil, errors.RestInternalServerError(">Error Getting User")
		}

		fmt.Println(">>>", sqlErr.Number)
		// switch sqlErr.Number{
		//
		// }

		if strings.Contains(err2.Error(), norows) {
			return nil, errors.RestNotFoundError("No Records Found")
		}
		return nil, errors.RestInternalServerError("Error Getting User")
	}
	return &retUser, nil

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
		if strings.Contains(err2.Error(), emailUnique) {
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

func UpdateUser(user *User) *errors.RestError {
	stmnt, err1 := bookstoredb.DBClient.Prepare(userUpdtQry)
	if err1 != nil {
		fmt.Println(err1)
		return errors.RestInternalServerError("Cannot Create Update Statement")
	}
	defer stmnt.Close()

	_, err2 := stmnt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err2 != nil {
		fmt.Println("The error is : ", err2.Error())
		return errors.RestInternalServerError("Cannot Update User")
	}

	return nil
}

func DeleteUser(id int64) *errors.RestError {
	stmnt, err1 := bookstoredb.DBClient.Prepare(userDeleteQry)
	if err1 != nil {
		fmt.Println(err1)
		return errors.RestInternalServerError("Cannot Create Delete Statement")
	}
	defer stmnt.Close()

	_, err2 := stmnt.Exec(id)

	if err2 != nil {
		fmt.Println("The error is : ", err2.Error())
		return errors.RestInternalServerError("Cannot Delete User")
	}

	return nil
}
