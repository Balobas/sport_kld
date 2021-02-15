package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sport_kld/app/controllers/organization_controller"
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