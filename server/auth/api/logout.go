package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/utils"
	"sport_kld/server/auth"
)

func Logout(ctx *gin.Context) {
	if utils.CheckHTTPMethod(ctx.Writer, http.MethodPost, ctx.Request.Method) != nil {
		return
	}
	authUid, ok := ctx.Get("authUid")
	if !ok {
		_, _ = ctx.Writer.Write([]byte("cant get token"))
	}

	if err := auth.DeleteAuth(authUid.(string)); err != nil {
		_, _ = ctx.Writer.Write([]byte("cant delete auth" + err.Error()))
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
