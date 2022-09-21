package logout

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth/gothic"
)

func logout(c *gin.Context) {
	redir := c.Request.URL.Query().Get("redir")
	if redir == "" {
		redir = "/api/user"
	}

	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, redir)
}
