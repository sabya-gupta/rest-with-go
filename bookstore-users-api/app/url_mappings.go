package app

import "github.com/sabya-gupta/rest-with-go/bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/user", controllers.CreateUser)
	router.PUT("/user/:user_id", controllers.UpdateUser)
	router.GET("/user/:user_id", controllers.GetUser)
	router.GET("/internal/users/search", controllers.FindUsersByStatus)
	router.DELETE("/user/:user_id", controllers.DeleteUser)
}
