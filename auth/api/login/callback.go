package login

import (
	"fmt"
	"net/http"
	"rwby-adventures/auth/store"

	"github.com/gin-gonic/gin"
	"github.com/pmylund/go-cache"
	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
)

func callback(c *gin.Context) {
	provider, err := goth.GetProvider("discord")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	value, err := gothic.GetFromSession(provider.Name(), c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	u, err := provider.FetchUser(sess)
	if err != nil {
		params := c.Request.URL.Query()
		if params.Encode() == "" && c.Request.Method == "POST" {
			c.Request.ParseForm()
			params = c.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), c.Request, c.Writer)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		u = gu
	}

	// Stores the data in the code
	code := uuid.NewV4().String()
	codes.Set(code, u, cache.DefaultExpiration)
	state := gothic.SetState(c.Request)
	redir, found := store.Redirections.Get(state)
	if !found {
		c.Redirect(http.StatusTemporaryRedirect, "/api/user")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?code=%s", redir.(string), code))
	}
}
