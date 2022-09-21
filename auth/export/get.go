package export

import (
	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
)

func GetUser(c *gin.Context) (user *goth.User, e error) {
	provider, err := goth.GetProvider("discord")
	if err != nil {
		return nil, err
	}

	value, err := gothic.GetFromSession(provider.Name(), c.Request)
	if err != nil {
		return nil, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), c.Request, c.Writer)
		if err != nil {
			return nil, err
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			return nil, err
		}
		u = gu
	}

	return &u, nil
}
