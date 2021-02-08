package main

import (
	"../app/api"
	"net/http"
)

func main() {
	http.HandleFunc("/place", api.GetPlaceByUID)
	http.HandleFunc("/places", api.GetPlacesByUIDs)
	http.HandleFunc("/places_by_fields", api.GetPlacesByFields)


	http.ListenAndServe(":8095", nil)
}

