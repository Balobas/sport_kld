/*
Данный пакет содержит методы обработчики запросов

 */
package api

import (
	"../controllers/place_controller"
	"../models/place_model"
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

func GetPlacesByUIDs(w http.ResponseWriter, r *http.Request) {
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

	uids := r.URL.Query()["uid"]
	if len(uids) == 0 {
		if _, err := w.Write([]byte("not found uids in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	resultParams := struct {
		Places []place_model.Place `json:"places"`
		Errors []string `json:"errors"`
	}{}

	var errs []error

	resultParams.Places, errs = place_controller.GetPlacesByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	b, err := json.Marshal(resultParams)
	if err != nil {
		if _, err := w.Write([]byte("cant marshal json")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}

func GetPlacesByFields(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("search") == "" {
		if _, err := w.Write([]byte("not found 'search' key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	searchString := r.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Places []place_model.Place  `json:"places"`
		Errors []string 			`json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByFields(r.URL.Query()["fields"], searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	b, err := json.Marshal(resultParams)
	if err != nil {
		if _, err := w.Write([]byte("cant marshal json")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}

func GetOrganizationPlaces(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("organization_uid") == "" {
		if _, err := w.Write([]byte("not found 'organization_uid' key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	orgUid := r.URL.Query().Get("organization_uid")

	var errs []error

	resultParams := struct {
		Places []place_model.Place  `json:"places"`
		Errors []string 			`json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByOrganizationUID(orgUid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	b, err := json.Marshal(resultParams)
	if err != nil {
		if _, err := w.Write([]byte("cant marshal json")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("cant write bytes")
	}
}

func GetPlacesByTag(w http.ResponseWriter, r *http.Request) {

}
