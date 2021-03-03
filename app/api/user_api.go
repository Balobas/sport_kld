package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/user_controller"
	"sport_kld/app/models/user_model"
	"sport_kld/app/utils"
)

func CreateUser(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var user user_model.User
	if err := decoder.Decode(&user); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	uid, token, err := user_controller.CreateUser(user)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	returnParams := struct {
		Uid string `json:"uid"`
		Token string `json:"token"`
	}{uid.String(), token}

	utils.WriteResult(ctx.Writer, returnParams)
}

func GetUser(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("uid") == "" {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found uid key in get query"))
		return
	}

	uid := ctx.Request.URL.Query()["uid"][0]
	user, err := user_controller.GetUserByUid(uid)
	if err != nil {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte(err.Error()))
		return
	}

	utils.WriteResult(ctx.Writer, user)
}

func GetUserByLogin(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodGet, ctx.Request.Method) != nil {
		return
	}

	if ctx.Request.URL.Query().Get("login") == "" {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte("not found uid key in get query"))
		return
	}

	login := ctx.Request.URL.Query()["login"][0]
	user, err := user_controller.GetUserByLogin(login)
	if err != nil {
		_ = utils.WriteToResponseWriter(ctx.Writer, []byte(err.Error()))
		return
	}

	utils.WriteResult(ctx.Writer, user)
}

func UpdateUser(ctx *gin.Context) {
	if utils.HandleHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	var user user_model.User
	if err := decoder.Decode(&user); err != nil {
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
	err := user_controller.UpdateUser(user, executorUid.(string))
	utils.WriteResult(ctx.Writer, err)
}
