package api

import (
	"fmt"
	"net/http"
	"sport_kld/app/controllers/event_controller"
	"sport_kld/app/models/event_model"
	"sport_kld/app/utils"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {

}

func JoinEvent(w http.ResponseWriter, r *http.Request) {

}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {

}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {

}

func GetEventByUid(w http.ResponseWriter, r *http.Request) {
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

	event, err := event_controller.GetEventByUid(uid)
	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(w, event)
}

func GetEventsByPlace(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("place_uid") == "" {
		if _, err := w.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := r.URL.Query()["place_uid"][0]

	resultParams := struct {
		Events        []event_model.Event               `json:"events"`
		Errors        []string                          `json:"errors"`
	}{}

	var errs []error

	resultParams.Events, errs = event_controller.GetEventsByPlace(uid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}

func ChangeEventPrivateStatus(w http.ResponseWriter, r *http.Request) {

}

func ChangeUserEventRole(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserFromEvent(w http.ResponseWriter, r *http.Request) {

}

func GetEventUserRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("event_uid") == "" && r.URL.Query().Get("user_uid") == "" {
		if _, err := w.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	eventUid := r.URL.Query()["event_uid"][0]
	userUid := r.URL.Query()["user_uid"][0]

	role, err := event_controller.GetEventUserRole(eventUid, userUid)
	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(w, role)
}

func PutEventInfoPost(w http.ResponseWriter, r *http.Request) {

}

func GetEventInfoPost(w http.ResponseWriter, r *http.Request) {
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

	post, err := event_controller.GetEventInfoPost(uid)
	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(w, post)
}

func GetEventInfoPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := w.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if r.URL.Query().Get("event_uid") == "" {
		if _, err := w.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := r.URL.Query()["event_uid"][0]

	resultParams := struct {
		EventInfoPosts        []event_model.EventInfoPost               `json:"posts"`
		Errors        		  []string                          		`json:"errors"`
	}{}

	var errs []error

	resultParams.EventInfoPosts, errs = event_controller.GetEventInfoPosts(uid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(w, resultParams)
}