package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/settings"
	"strings"
)

func AuthorizationMiddleware(ctx *gin.Context) {

	checkPair := pair{ctx.Request.Method, ctx.Request.URL.Path}
	needAuth := OnlyWithAuth[checkPair]
	if !needAuth {
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userUid, err := ParseToken(headerParts[1], settings.SIGNING_KEY)
	if err != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatus(status)
		return
	}

	ctx.Set("executorUid", userUid)
}
