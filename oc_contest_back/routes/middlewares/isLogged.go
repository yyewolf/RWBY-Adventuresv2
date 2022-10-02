package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, found := ctx.Get("Logged")
		if !found {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in."})
			return
		}
		loggedIn, ok := data.(bool)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in."})
			return
		}
		if !loggedIn {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in."})
			return
		}
	}
}
