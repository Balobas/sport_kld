package main

import (
	"net/http"
	"sport_kld/app/api"
)

func main() {
	http.HandleFunc("/place", api.GetPlaceByUID)
	http.HandleFunc("/places", api.GetPlacesByUIDs)
	http.HandleFunc("/places_by_fields", api.GetPlacesByFields)
	http.HandleFunc("/organization_places", api.GetOrganizationPlaces)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {})
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {})

	http.ListenAndServe(":8080", nil)
}
