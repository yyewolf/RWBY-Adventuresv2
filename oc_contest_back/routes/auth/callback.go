package auth

import (
	"net/http"
	"rwby-adventures/auth/export"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	_, err := export.Callback(c)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080")
}
