package store

import (
	"fmt"
	"rwby-adventures/config"
	"time"

	"github.com/gorilla/sessions"
	"github.com/pmylund/go-cache"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

var Redirections = cache.New(5*time.Minute, 10*time.Minute)

func init() {
	maxAge := 86400 * 30 // 30 days cookie

	store := sessions.NewCookieStore(config.CookieKey)
	store.MaxAge(maxAge)
	store.Options.Domain = config.CookieHost
	store.Options.Path = "/"
	gothic.Store = store

	goth.UseProviders(
		discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("%sapi/login/callback", config.AuthHost), discord.ScopeIdentify, discord.ScopeGuilds),
	)
}
