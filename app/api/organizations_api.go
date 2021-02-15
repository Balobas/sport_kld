package api

import (
	"encoding/json"
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
	b, err := json.Marshal(org)
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
		Organizations []organization_model.Organization `json:"organization"`
		Errors []string `json:"errors"`
	}{}

	var errs []error

	resultParams.Organizations, errs = organization_controller.GetOrganizationsByUIDs(uids)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	b, err := json.Marshal(resultParams)
	if err != nil {
		_ = utils.WriteToResponseWriter(w, []byte("cant marshal json"))
		return
	}
	_ = utils.WriteToResponseWriter(w, b)
}

