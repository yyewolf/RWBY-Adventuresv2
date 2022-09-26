package export

import (
	"fmt"
	"net/http"
	"rwby-adventures/config"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth/gothic"
)

func Logout(c *gin.Context, redir string) {
	gothic.Logout(c.Writer, c.Request)
	URL := fmt.Sprintf("%sapi/logout?redir=%s", config.AuthHost, redir)
	c.Redirect(http.StatusTemporaryRedirect, URL)
}
