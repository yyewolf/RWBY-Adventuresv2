package web

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"rwby-adventures/auth/export"
	"rwby-adventures/config"
	"rwby-adventures/market_back/websocket"

	"github.com/gin-gonic/gin"
	"github.com/pmylund/go-cache"
	uuid "github.com/satori/go.uuid"
)

func startMarketService(g *gin.RouterGroup) {
	g.GET("/logout/", func(c *gin.Context) {
		export.Logout(c, config.MarketFront)
	})

	g.GET("/login/:token", MarketLogin)
	g.GET("/login/callback/:token", MarketCallback)

	g.GET("/token", MarketGetToken)
}

func MarketCallback(c *gin.Context) {
	token := c.Param("token")
	t, found := websocket.Tokens.Get(token)
	if !found {
		c.Redirect(http.StatusTemporaryRedirect, config.MarketFront)
		return
	}
	tk := t.(*websocket.Token)

	u, err := export.Callback(c)
	if err != nil {
		tk.Empty = true
		c.Redirect(http.StatusTemporaryRedirect, config.MarketFront)
		return
	}

	tk.Empty = false
	tk.UserID = u.UserID
	tk.Secret = u.AccessToken
	c.Redirect(http.StatusTemporaryRedirect, config.MarketFront)
}

func MarketLogin(c *gin.Context) {
	token := c.Param("token")
	_, found := websocket.Tokens.Get(token)
	if !found {
		c.Redirect(http.StatusTemporaryRedirect, config.MarketFront)
		return
	}
	export.Login(c, fmt.Sprintf("%slogin/callback/%s", config.MarketHost, token))
}

func MarketGetToken(c *gin.Context) {
	token := fmt.Sprintf("%x", sha256.Sum256(uuid.NewV4().Bytes()))
	t := &websocket.Token{
		Empty: true,
		Token: token,
	}

	u, err := export.GetUser(c)
	if err == nil {
		t.Empty = false
		t.UserID = u.UserID
		t.Secret = u.AccessTokenSecret
	}

	websocket.Tokens.Set(token, t, cache.DefaultExpiration)

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
