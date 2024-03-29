package login

import (
	"context"
	"fmt"
	"net/http"
	"rwby-adventures/auth/store"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmylund/go-cache"
	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
)

var codes = cache.New(7*24*time.Hour, 1*time.Hour)

func doLogin(c *gin.Context) {
	redir := c.Request.URL.Query().Get("redir")
	if redir == "" {
		redir = "/api/user"
	}
	state := gothic.SetState(c.Request)
	store.Redirections.Set(state, redir, cache.DefaultExpiration)
	c.Request.URL.RawQuery += fmt.Sprintf("&state=%s", state)
	state = gothic.SetState(c.Request)
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", "discord"))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func startLogin(c *gin.Context) {
	provider, err := goth.GetProvider("discord")
	if err != nil {
		doLogin(c)
		return
	}

	value, err := gothic.GetFromSession(provider.Name(), c.Request)
	if err != nil {
		doLogin(c)
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		doLogin(c)
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
			doLogin(c)
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), c.Request, c.Writer)

		if err != nil {
			doLogin(c)
			return
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			doLogin(c)
			return
		}
		u = gu
	}

	code := uuid.NewV4().String()
	codes.Set(code, u, cache.DefaultExpiration)
	redir := c.Request.URL.Query().Get("redir")
	if redir == "" {
		redir = "/api/user"
	}
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?code=%s", redir, code))
}
