package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/settings"
	"strings"
)

func AuthorizationMiddleware(ctx *gin.Context) {
	key := ctx.Request.Method + ctx.Request.URL.Path
	needAuth, ok := OnlyWithAuth[key]
	fmt.Println("request from", ctx.Request.Host, ctx.Request.RemoteAddr)

	if !needAuth {
		fmt.Println("not need auth", ok)
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

	// TODO: в случае, если у токена истек срок годности, продлить его при условии что действует refresh token

	userUid, err := ParseToken(headerParts[1], settings.SIGNING_KEY)
	if err != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatus(status)
		return
	}

	ctx.Set("executorUid", userUid)
}
