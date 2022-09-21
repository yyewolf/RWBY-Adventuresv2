package export

import (
	"fmt"
	"rwby-adventures/config"

	"github.com/gorilla/sessions"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/providers/discord"
)

func Provider() {
	maxAge := 86400 * 30 // 30 days cookie

	store := sessions.NewCookieStore(config.CookieKey)
	store.MaxAge(maxAge)
	store.Options.Domain = config.CookieHost
	store.Options.Path = "/"

	goth.UseProviders(
		discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("%sapi/login/callback", config.AuthHost), discord.ScopeIdentify, discord.ScopeGuilds),
	)
}
