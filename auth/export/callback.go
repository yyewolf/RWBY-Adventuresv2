package export

import (
	"errors"

	"github.com/yyewolf/goth/providers/discord"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
)

func Callback(c *gin.Context) (u *goth.User, err error) {
	secret := c.Request.URL.Query().Get("secret")
	refresh := c.Request.URL.Query().Get("refresh")
	if secret == "" || refresh == "" {
		return nil, errors.New("invalid secret or refresh")
	}

	provider, err := goth.GetProvider("discord")
	if err != nil {
		return nil, err
	}
	g, err := provider.BeginAuth("")
	if err != nil {
		return nil, err
	}
	s := g.(*discord.Session)
	s.AccessToken = secret
	s.RefreshToken = refresh

	user, err := provider.FetchUser(s)
	if err != nil {
		return nil, err
	}

	err = gothic.StoreInSession(provider.Name(), s.Marshal(), c.Request, c.Writer)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
