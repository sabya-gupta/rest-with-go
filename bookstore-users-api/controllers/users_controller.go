package controllers

import (
	"fmt"
	"net/http"

	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/errors"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/models"
	"github.com/sabya-gupta/rest-with-go/bookstore-users-api/services"

	"github.com/gin-gonic/gin"
)

var counter int

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet Implemented")
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
func SearchUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not yet Implemented")
}
