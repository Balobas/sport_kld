package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/server"
	"sport_kld/server/auth"
)

func main() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(auth.AuthorizationMiddleware)

	router.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	server.InitRoutes(router)

	router.Run(":" + "8080")
}
