package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sport_kld/app/api"
	"sport_kld/baseMiddleware"
	rt "sport_kld/server/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(func(context *gin.Context) {
		baseMiddleware.AuthorizationMiddleware(context, api.OnlyWithAuth)
	})

	router.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	rt.InitRoutes(router)

	router.Run(":" + port)
}
