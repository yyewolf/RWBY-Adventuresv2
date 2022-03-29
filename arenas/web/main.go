package web

import (
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/main/static"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

var templates *template.Template

func StartWeb() {
	templates, _ = template.ParseFS(static.WebFS, "*.html")

	goth.UseProviders(
		discord.New("375700234120200194", "P6KOz6Uvl8PWhY-hfx5IXo_posPDBu7D", fmt.Sprintf("http://%s%s/auth/discord/callback", config.ArenaHost, config.ArenaPort), discord.ScopeIdentify),
	)

	startArenaService()
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
		return goth.User{}, err
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, req.URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	err = gothic.StoreInSession(providerName, sess.Marshal(), req, res)

	if err != nil {
		return goth.User{}, err
	}

	gu, err := provider.FetchUser(sess)
	return gu, err
}
