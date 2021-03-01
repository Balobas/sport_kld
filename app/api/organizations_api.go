package api

import (
	"fmt"
	"net/http"
	"sport_kld/app/controllers/organization_controller"
	"sport_kld/app/models/organization_model"
	"sport_kld/app/utils"
)

func GetOrganizationByUID(w http.ResponseWriter, r *http.Request) {
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

	org, err := organization_controller.GetOrganizationByUID(uid)
	if err != nil {
		res := []byte(err.Error())
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	utils.WriteResult(w, org)
}

func GetOrganizationsByUIDs(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
		return
	}

	uids := r.URL.Query()["uid"]
	if len(uids) == 0 {
		_ = utils.WriteToResponseWriter(w, []byte("not found uids in get query"))
		return
	}

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	var errs []error

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func GetOrganizationsByFields(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
		return
	}

	if r.URL.Query().Get("search") == "" {
		_ = utils.WriteToResponseWriter(w, []byte("not found 'search' key in get query"))
		return
	}

	searchString := r.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByFields(r.URL.Query()["fields"], searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func GetPlaceOrganization(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
		return
	}

	if r.URL.Query().Get("place_uid") == "" {
		_ = utils.WriteToResponseWriter(w, []byte("not found 'organization_uid' key in get query"))
		return
	}

	placeUid := r.URL.Query().Get("place_uid")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationByPlaceUID(placeUid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func GetOrganizationsByTag(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodGet, r.Method) != nil {
		return
	}

	if r.URL.Query().Get("search") == "" {
		_ = utils.WriteToResponseWriter(w, []byte("not found 'search' key in get query"))
		return
	}

	searchString := r.URL.Query().Get("search")

	var errs []error

	resultParams := struct {
		Organizations []organization_model.Organization `json:"organizations"`
		Errors        []string                          `json:"errors"`
	}{}

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByTags(searchString)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}
