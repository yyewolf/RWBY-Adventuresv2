package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
)

func GetUser(c *gin.Context) {
	e := gin.H{
		"error":  "User not logged in",
		"logged": false,
	}
	provider, err := goth.GetProvider("discord")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, e)
		return
	}

	value, err := gothic.GetFromSession(provider.Name(), c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, e)
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, e)
		return
	}

	user, err := provider.FetchUser(sess)
	if err != nil {
		params := c.Request.URL.Query()
		if params.Encode() == "" && c.Request.Method == "POST" {
			c.Request.ParseForm()
			params = c.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, e)
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), c.Request, c.Writer)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, e)
			return
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, e)
			return
		}
		user = gu
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "User is logged in",
		"user":    user,
		"logged":  true,
	})
}
