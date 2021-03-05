package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

	authUid, userUid, err := VerifyAuth(headerParts[1])
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	ctx.Set("executorUid", userUid)
	ctx.Set("authUid", authUid)
}
