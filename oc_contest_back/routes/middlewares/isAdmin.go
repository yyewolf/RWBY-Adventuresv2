package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
)

var list = []string{
	"144472011924570113",
}

func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, found := ctx.Get("User")
		if !found {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not admin."})
			return
		}
		user, ok := data.(*goth.User)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not admin."})
			return
		}
		for _, id := range list {
			if id == user.UserID {
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not admin."})
	}
}
