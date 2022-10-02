package auth

import (
	"rwby-adventures/auth/export"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	export.Logout(c, "http://localhost:8080")
}
