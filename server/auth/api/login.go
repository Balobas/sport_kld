package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/user_controller"
	"sport_kld/app/models/user_model"
	"sport_kld/app/utils"
	"sport_kld/server/auth"
	"sport_kld/server/auth/token"
)

func Login(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}
	decoder := json.NewDecoder(ctx.Request.Body)
	params := struct {
		Login    string    `json:"login"`
		Password string    `json:"password"`
	}{}
	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	user := user_model.User{}
	user.Login = params.Login
	user.Password = params.Password

	// check user in db
	if err := user_controller.VerifyUser(user); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	user, err := user_controller.GetUserByLogin(params.Login)
	if err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	// Пока допускается только один активный клиент для пользователя
	if err := auth.ExistAuth(user.UID.String()); err == nil {
		_, _ = ctx.Writer.Write([]byte("user already login"))
		return
	}

	td, err := token.CreateToken(user.UID.String())
	if err != nil {
		_, _ = ctx.Writer.Write([]byte(err.Error()))
		return
	}

	if err := auth.CreateAuth(user.UID.String(), td); err != nil {
		_, _ = ctx.Writer.Write([]byte(err.Error()))
		return
	}

	result := struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{td.AccessToken, td.RefreshToken}
	ctx.JSON(http.StatusOK, result)
}
