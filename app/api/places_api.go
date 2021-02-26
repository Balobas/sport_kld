/*
Данный пакет содержит методы обработчики запросов

*/
package api

import (
	"fmt"
	"net/http"
	"sport_kld/app/controllers/place_controller"
	"sport_kld/app/models/place_model"
	"sport_kld/app/utils"
)

func GetPlaceByUID(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
		return
	}

	if r.URL.Query().Get("uid") == "" {
		if utils.WriteToResponseWriter(w, []byte("not found uid key in get query")) != nil {
			return
		}
		return
	}

	uid := r.URL.Query()["uid"][0]

	place, err := place_controller.GetPlaceByUID(uid)
	if err != nil {
		_ = utils.WriteToResponseWriter(w, []byte(err.Error()))
		return
	}

	utils.WriteResult(w, place)
}

func GetPlacesByUIDs(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
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
		Errors []string            `json:"errors"`
	}{}

	var errs []error

	resultParams.Places, errs = place_controller.GetPlacesByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func GetPlacesByFields(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		if utils.WriteToResponseWriter(w, []byte("wrong http method. access denied")) != nil {
			return
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
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByFields(r.URL.Query()["fields"], searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
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
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByOrganizationUID(orgUid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func GetPlacesByTag(w http.ResponseWriter, r *http.Request) {
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
		Places []place_model.Place `json:"places"`
		Errors []string            `json:"errors"`
	}{}

	resultParams.Places, errs = place_controller.GetPlacesByTags(searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}
