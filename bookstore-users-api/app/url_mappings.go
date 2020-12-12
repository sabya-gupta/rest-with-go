package app

import "github.com/sabya-gupta/rest-with-go/bookstore-users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:user_id", controllers.GetUser)
	router.GET("/search/user", controllers.SearchUsers)
}