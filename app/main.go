package main

import (
	"../app/api"
	"net/http"
)

func main() {
	http.HandleFunc("/place", api.GetPlaceByUID)
	http.HandleFunc("/places", api.GetPlacesByUIDs)

	http.ListenAndServe(":8095", nil)
}

