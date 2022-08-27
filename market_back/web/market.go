package web

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/market_back/static"
	"rwby-adventures/market_back/websocket"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

var provider = discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("%sauth/discord/callback", config.MarketHost), discord.ScopeIdentify)
var stateToken = make(map[string]*websocket.Token)

func startMarketService() {

	key := uuid.NewV5(uuid.NewV4(), "cookies").Bytes() // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                               // 30 days

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Domain = config.MarketDomain
	store.Options.Path = "/"

	gothic.Store = store

	mux := pat.New()
	port := config.MarketPort
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}

	mux.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(provider)
		state := gothic.SetState(req)
		token := stateToken[state]
		usr, err := UserLogged(res, req)
		if err != nil {
			token.Empty = true
		} else {
			token.Empty = false
			token.UserID = usr.UserID
			token.Secret = usr.AccessTokenSecret
		}
		http.Redirect(res, req, config.MarketFront, http.StatusTemporaryRedirect)
		res.Write([]byte(""))
	})

	mux.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(provider)
		gothic.Logout(res, req)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		res.Write([]byte(""))
	})

	mux.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(provider)
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(`<p>UserID: {{.UserID}}</p>`)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", DirectoryListing(http.FileServer(http.FS(static.Assets)))))
	mux.HandleFunc("/", http.HandlerFunc(MarketIndex))
	mux.HandleFunc("/login/{token}", http.HandlerFunc(MarketLogin))
	go srv.ListenAndServe()
}

func DirectoryListing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MarketLogin(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get(":token")
	t, found := websocket.Tokens.Get(token)
	if !found {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	goth.UseProviders(provider)
	state := gothic.SetState(r)
	stateToken[state] = t.(*websocket.Token)
	r.URL.RawQuery += fmt.Sprintf("&state=%s", state)
	state = gothic.SetState(r)
	r = r.WithContext(context.WithValue(r.Context(), "provider", "discord"))
	gothic.BeginAuthHandler(w, r)
}

func MarketIndex(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "market.html", nil)
}
