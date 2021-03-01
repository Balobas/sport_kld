package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sport_kld/app/settings"
	"strings"
)

func CheckToken(ctx *gin.Context) {
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

	if _, err := ParseToken(headerParts[1], settings.SIGNING_KEY); err != nil {
		status := http.StatusBadRequest
		ctx.AbortWithStatus(status)
		return
	}
}
