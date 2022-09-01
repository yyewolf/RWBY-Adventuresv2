package web

import (
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/dungeons_back/static"

	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

var templates *template.Template

func StartWeb() {
	templates, _ = template.ParseFS(static.WebFS, "*.html")

	goth.UseProviders(
		discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("%sauth/discord/callback", config.DungeonHost), discord.ScopeIdentify),
	)

	startDungeonService()
	fmt.Println("[WEB] Started.")
}

func UserLogged(res http.ResponseWriter, req *http.Request) (goth.User, error) {
	providerName := "discord"

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return goth.User{}, err
	}

	value, err := gothic.GetFromSession(providerName, req)
	if err != nil {
		fmt.Println("GetFromSession didn't work")
		return goth.User{}, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		fmt.Println("UnmarshalSession didn't work")
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}
	fmt.Println("FetchUser didn't work")

	// get new token and retry fetch
	_, err = sess.Authorize(provider, req.URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	err = gothic.StoreInSession(providerName, sess.Marshal(), req, res)

	if err != nil {
		fmt.Println("StoreInSession didn't work", err)
		return goth.User{}, err
	}

	gu, err := provider.FetchUser(sess)
	return gu, err
}
