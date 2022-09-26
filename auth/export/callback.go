package export

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

func Callback(c *gin.Context) (u *goth.User, err error) {
	code := c.Request.URL.Query().Get("code")
	if code == "" {
		return nil, errors.New("invalid code")
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

	user, err := Code(code)
	if err != nil {
		return nil, err
	}

	s.AccessToken = user.AccessToken
	s.RefreshToken = user.RefreshToken

	err = gothic.StoreInSession(provider.Name(), s.Marshal(), c.Request, c.Writer)

	if err != nil {
		return nil, err
	}

	return user, nil
}
