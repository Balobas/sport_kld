package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sport_kld/app/controllers/event_controller"
	"sport_kld/app/models/event_model"
	"sport_kld/app/utils"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var event event_model.Event

	if err := decoder.Decode(&event); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid, err := event_controller.CreateEvent(event)
	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(w, uid)
}

func JoinEvent(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct{
		UserUid string `json:"userUid"`
		EventUid string `json:"eventUid"`
		Password string `json:"password"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.JoinEvent(params.UserUid, params.EventUid, params.Password)
	utils.WriteResult(w, err)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct {
		Event event_model.Event `json:"event"`
		ExecutorUid string `json:"executorUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.UpdateEvent(params.Event, params.ExecutorUid)
	utils.WriteResult(w, err)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodDelete, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct {
		EventUid string `json:"eventUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.DeleteEvent(params.EventUid)
	utils.WriteResult(w, err)
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
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct{
		EventUid string `json:"eventUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.ChangeEventPrivateStatus(params.EventUid)
	utils.WriteResult(w, err)
}

func ChangeUserEventRole(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct{
		EventUser event_model.EventUser `json:"eventUser"`
		ExecutorUid string `json:"executorUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.ChangeUserEventRole(params.EventUser, params.ExecutorUid)
	utils.WriteResult(w, err)
}

func DeleteUserFromEvent(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodDelete, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct {
		UserUid  string `json:"userUid"`
		EventUid string `json:"eventUid"`
		ExecutorUid string `json:"executorUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.DeleteEventUser(params.UserUid, params.EventUid, params.ExecutorUid)
	utils.WriteResult(w, err)
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
	if utils.HandleHTTPMethod(w, http.MethodPost, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var post event_model.EventInfoPost
	if err := decoder.Decode(&post); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid, err := event_controller.PutEventInfoPost(post)
	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(w, uid)
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

func DeleteEventInfoPost(w http.ResponseWriter, r *http.Request) {
	if utils.HandleHTTPMethod(w, http.MethodDelete, r.Method) != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params struct {
		PostUid string `json:"uid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.DeleteEvent(params.PostUid)
	utils.WriteResult(w, err)
}
