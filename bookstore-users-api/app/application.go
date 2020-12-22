package app

import "github.com/gin-gonic/gin"

var (
	router = gin.New()
)

func StartApplication() {
	mapUrls()
	router.Run(":3000")
}
