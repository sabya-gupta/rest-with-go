package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/models"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/services"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

var counter int

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		retErr := errors.RestBadRequestError("invalid user id")
		c.JSON(retErr.Status, retErr)
		return
	}

	user, getUserErr := services.GetUser(userId)
	if getUserErr != nil {
		c.JSON(getUserErr.Status, getUserErr)
		return
	}

	c.JSON(http.StatusFound, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(bytes))
	// if err != nil {
	// 	fmt.Println("'-------------->>>'")
	// 	fmt.Println(err.Error(), err)
	// 	return
	// }
	//
	// e := json.Unmarshal(bytes, &user)
	// if e != nil {
	// 	fmt.Println("'-------------->>>'")
	// 	fmt.Println(e.Error(), e)
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		// retErr := errors.RestError{
		// 	Status:  http.StatusBadRequest,
		// 	Error:   "Bad Request",
		// 	Message: "User Json Not Proper",
		// }
		// retErr := errors.RestError{}
		// retErr.Error = "Bad Request"
		// retErr.Status = http.StatusBadRequest
		// retErr.Message = "User Json Not Proper!"
		retErr := errors.RestBadRequestError("User Json Not Proper!")
		c.JSON(retErr.Status, &retErr)
		return
	}
	fmt.Println(user)
	result, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		fmt.Println(saveErr)
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		retErr := errors.RestBadRequestError("invalid user id")
		c.JSON(retErr.Status, retErr)
		return
	}

	var newuser models.User
	if err := c.ShouldBindJSON(&newuser); err != nil {
		fmt.Println(err)
		retErr := errors.RestBadRequestError("User Json Not Proper!")
		c.JSON(retErr.Status, &retErr)
		return
	}

	newuser.Id = userId
	upDateErr := services.UpdateUser(&newuser)
	if upDateErr != nil {
		fmt.Println(upDateErr)
		c.JSON(upDateErr.Status, upDateErr)
		return
	}
	fmt.Println(newuser)
	c.JSON(http.StatusOK, newuser)
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		retErr := errors.RestBadRequestError("invalid user id")
		c.JSON(retErr.Status, retErr)
		return
	}

	upDateErr := services.DeleteUser(userId)
	if upDateErr != nil {
		fmt.Println(upDateErr)
		c.JSON(upDateErr.Status, upDateErr)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet Implemented")
}
