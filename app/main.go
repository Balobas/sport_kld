package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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
		context.JSON(http.StatusOK, "{lol:kek, kek:lol}")
	})

	router.Run(":" + port)

	//http.ListenAndServe(":8080", nil)
}
