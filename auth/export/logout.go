package export

import (
	"fmt"
	"net/http"
	"rwby-adventures/config"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context, redir string) {
	URL := fmt.Sprintf("%sapi/logout?redir=%s", config.AuthHost, redir)
	c.Redirect(http.StatusTemporaryRedirect, URL)
}
