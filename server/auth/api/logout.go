package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/server/auth"
)

func Logout(ctx *gin.Context) {
	authUid, ok := ctx.Get("authUid")
	if !ok {
		_, _ = ctx.Writer.Write([]byte("cant get token"))
	}
	err := auth.DeleteAuth(authUid.(string))
	if err != nil {
		_, _ = ctx.Writer.Write([]byte("cant delete auth" + err.Error()))
	}
	ctx.AbortWithStatus(http.StatusOK)
}
