/*
Данный пакет содержит методы обработчики запросов

 */
package api

import (
	"../controllers/place_controller"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPlaceByUID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("uid") == "" {
		if _, err := w.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := r.URL.Query()["uid"][0]

	place, err := place_controller.GetPlaceByUID(uid)
	if err != nil {
		res := []byte(err.Error())
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	b, err := json.Marshal(place)
	if err != nil {
		res := []byte(err.Error())
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}

func GetPlacesByUIDs(w http.ResponseWriter, r *http.Request) {}

func GetPlacesByFields(w http.ResponseWriter, r *http.Request) {

}

func GetOrganizationPlaces(w http.ResponseWriter, r *http.Request) {

}

func GetPlacesByTag(w http.ResponseWriter, r *http.Request) {

}


