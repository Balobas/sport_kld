package baseMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/server/auth"
	"strings"
)

func AuthorizationMiddleware(ctx *gin.Context, AccessRoutesMap map[string]bool) {
	key := ctx.Request.Method + ctx.Request.URL.Path
	needAuth, ok := AccessRoutesMap[key]
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

	authUid, userUid, err := auth.VerifyAuth(headerParts[1])
	if err != nil {
		_ = ctx.AbortWithError(http.StatusUnauthorized, err)
	}

	ctx.Set("executorUid", userUid)
	ctx.Set("authUid", authUid)
}
