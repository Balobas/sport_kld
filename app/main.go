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
	http.HandleFunc("/places_by_tag", api.GetPlacesByTag)

	http.HandleFunc("/organization", api.GetOrganizationByUID)
	http.HandleFunc("/organizations", api.GetOrganizationsByUIDs)
	http.HandleFunc("/organizations_by_fields", api.GetOrganizationsByFields)
	http.HandleFunc("/place_organization", api.GetPlaceOrganization)
	http.HandleFunc("/organizations_by_tag", api.GetOrganizationsByTag)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {writer.Write([]byte("//////"))})
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {writer.Write([]byte("favicon ico"))})

	http.ListenAndServe(":8080", nil)
}
