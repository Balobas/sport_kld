package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/server/auth"
	"sport_kld/server/auth/token"
)

func Refresh(ctx *gin.Context) {
	params := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	decoder := json.NewDecoder(ctx.Request.Body)
	if err := decoder.Decode(&params); err != nil {
		if _, err := ctx.Writer.Write([]byte(err.Error())); err != nil {
			fmt.Println("cant write bytes")
		}
		return
	}

	authUid, userUid, err := auth.VerifyRefresh(params.RefreshToken)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := auth.DeleteAuth(authUid); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tk, err := token.CreateToken(userUid)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := auth.CreateAuth(userUid, tk); err != nil {
		_ = ctx.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	result := struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{tk.AccessToken, tk.RefreshToken}
	ctx.JSON(http.StatusOK, result)
}
