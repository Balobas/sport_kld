package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sport_kld/app/api"
)

func main() {
	//http.HandleFunc("/place", api.GetPlaceByUID)
	//http.HandleFunc("/places", api.GetPlacesByUIDs)
	//http.HandleFunc("/places_by_fields", api.GetPlacesByFields)
	//http.HandleFunc("/organization_places", api.GetOrganizationPlaces)
	//http.HandleFunc("/places_by_tag", api.GetPlacesByTag)
	//
	//http.HandleFunc("/organization", api.GetOrganizationByUID)
	//http.HandleFunc("/organizations", api.GetOrganizationsByUIDs)
	//http.HandleFunc("/organizations_by_fields", api.GetOrganizationsByFields)
	//http.HandleFunc("/place_organization", api.GetPlaceOrganization)
	//http.HandleFunc("/organizations_by_tag", api.GetOrganizationsByTag)
	//
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {writer.Write([]byte("//////"))})
	//http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {writer.Write([]byte("favicon ico"))})

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, "")
	})

	router.GET("/place", gin.WrapF(api.GetPlaceByUID))
	router.GET("/places", gin.WrapF(api.GetPlacesByUIDs))
	router.GET("/places_by_fields", gin.WrapF(api.GetPlacesByFields))
	router.GET("/organization_places", gin.WrapF(api.GetOrganizationPlaces))
	router.GET("/places_by_tag", gin.WrapF(api.GetPlacesByTag))

	router.GET("/organization", gin.WrapF(api.GetOrganizationByUID))
	router.GET("/organizations", gin.WrapF(api.GetOrganizationsByUIDs))
	router.GET("/organizations_by_fields", gin.WrapF(api.GetOrganizationsByFields))
	router.GET("/place_organization", gin.WrapF(api.GetPlaceOrganization))
	router.GET("/organizations_by_tag", gin.WrapF(api.GetOrganizationsByTag))

	router.Run(":" + port)

	//http.ListenAndServe(":8080", nil)
}
