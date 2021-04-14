package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/event_controller"
	"sport_kld/app/models/event_model"
	"sport_kld/app/utils"
)

func CreateEvent(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var event event_model.Event

	if err := decoder.Decode(&event); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid, err := event_controller.CreateEvent(event)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(ctx.Writer, uid)
}

func JoinEvent(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		UserUid  string `json:"userUid"`
		EventUid string `json:"eventUid"`
		Password string `json:"password"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	err := event_controller.JoinEvent(params.UserUid, params.EventUid, params.Password)
	utils.WriteResult(ctx.Writer, err)
}

func UpdateEvent(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		Event       event_model.Event `json:"event"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}

	err := event_controller.UpdateEvent(params.Event, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}

func DeleteEvent(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodDelete, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		EventUid string `json:"eventUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}
	err := event_controller.DeleteEvent(params.EventUid, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}

func GetEventByUid(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := ctx.Request.URL.Query()["uid"][0]

	event, err := event_controller.GetEventByUid(uid)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(ctx.Writer, event)
}

func GetEventsByPlace(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("place_uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := ctx.Request.URL.Query()["place_uid"][0]

	resultParams := struct {
		Events []event_model.Event `json:"events"`
		Errors []string            `json:"errors"`
	}{}

	var errs []error

	resultParams.Events, errs = event_controller.GetEventsByPlace(uid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func ChangeEventPrivateStatus(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		EventUid string `json:"eventUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}

	err := event_controller.ChangeEventPrivateStatus(params.EventUid, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}

func ChangeUserEventRole(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		EventUser   event_model.EventUser `json:"eventUser"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}

	err := event_controller.ChangeUserEventRole(params.EventUser, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}

func DeleteUserFromEvent(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodDelete, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		UserUid     string `json:"userUid"`
		EventUid    string `json:"eventUid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}

	err := event_controller.DeleteEventUser(params.UserUid, params.EventUid, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}

func GetEventUserRole(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("event_uid") == "" && ctx.Request.URL.Query().Get("user_uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	eventUid := ctx.Request.URL.Query()["event_uid"][0]
	userUid := ctx.Request.URL.Query()["user_uid"][0]

	role, err := event_controller.GetEventUserRole(eventUid, userUid)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(ctx.Writer, role)
}

func PutEventInfoPost(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var post event_model.EventInfoPost
	if err := decoder.Decode(&post); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid, err := event_controller.PutEventInfoPost(post)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(ctx.Writer, uid)
}

func GetEventInfoPost(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := ctx.Request.URL.Query()["uid"][0]

	post, err := event_controller.GetEventInfoPost(uid)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	utils.WriteResult(ctx.Writer, post)
}

func GetEventInfoPosts(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		res := []byte("wrong http method. access denied")
		if _, err := ctx.Writer.Write(res); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	if ctx.Request.URL.Query().Get("event_uid") == "" {
		if _, err := ctx.Writer.Write([]byte("not found uid key in get query")); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid := ctx.Request.URL.Query()["event_uid"][0]
	resultParams := struct {
		EventInfoPosts []event_model.EventInfoPost `json:"posts"`
		Errors         []string                    `json:"errors"`
	}{}

	var errs []error
	resultParams.EventInfoPosts, errs = event_controller.GetEventInfoPosts(uid)

	for _, err := range errs {
		resultParams.Errors = append(resultParams.Errors, err.Error())
	}

	utils.WriteResult(ctx.Writer, resultParams)
}

func DeleteEventInfoPost(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodDelete, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var params struct {
		PostUid string `json:"uid"`
	}

	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}
	executorUid, ok := ctx.Get("executorUid")
	if !ok {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("cant get executor uid"))
		return
	}
	err := event_controller.DeleteEvent(params.PostUid, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}
