package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/controllers/user_controller"
	"sport_kld/app/models/user_model"
	"sport_kld/server/auth"
	"sport_kld/server/auth/token"
)

func Login(ctx *gin.Context) {

	decoder := json.NewDecoder(ctx.Request.Body)
	var user user_model.User
	if err := decoder.Decode(&user); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	// check user in db
	if err := user_controller.VerifyUser(user); err != nil {
		return
	}

	// Пока допускается только один активный клиент для пользователя
	if err := auth.AuthExist(user.UID.String()); err == nil {
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
		AccessToken string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{td.AccessToken, td.RefreshToken}
	ctx.JSON(http.StatusOK, result)
}
