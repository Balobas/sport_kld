package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
	"net/http"
	"sport_kld/server"
)

func main() {
	//port := os.Getenv("PORT")
	//if port == "" {
	//	log.Fatal("$PORT must be set")
	//}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(ginsession.New())

	router.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	router.GET("/test_session_set", func(context *gin.Context) {
		store := ginsession.FromContext(context)

		store.Set("kk",  store.SessionID())
		if err := store.Save(); err != nil {
			fmt.Println("err")
			context.AbortWithError(500, err)
			return
		}
		context.Redirect(302, "/test_session_get")
	})
	router.GET("/test_session_get", func(context *gin.Context) {
		store := ginsession.FromContext(context)
		if store == nil {
			fmt.Println("nil")
		}

		val, ok := store.Get("kk")
		if !ok {
			context.AbortWithStatus(404)
			return
		}

		context.String(http.StatusOK, "%s", val)
	})


	server.InitRoutes(router)

	router.Run(":" + "8080")
}
