package auth

import (
	"rwby-adventures/auth/export"
	"rwby-adventures/config"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	export.Logout(c, config.OCFront)
}
