package auth

import (
	"net/http"
	"rwby-adventures/auth/export"
	"rwby-adventures/config"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	_, err := export.Callback(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, config.OCFront)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, config.OCFront)
}
