package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-isatty"
	"net/http"
	rt "sport_kld/server/router"
)

func main() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}

	router := gin.New()
	router.Use(gin.Logger())
	//router.Use(auth.AuthorizationMiddleware)

	router.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})
	router.GET("/favicon.ico", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	rt.InitRoutes(router)

	router.Run(":" + "8080")
}
