package web

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/main/static"
	"rwby-adventures/main/websocket"
	"rwby-adventures/models"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	uuid "github.com/satori/go.uuid"
)

var provider = discord.New("375700234120200194", "P6KOz6Uvl8PWhY-hfx5IXo_posPDBu7D", fmt.Sprintf("http://%s%s/auth/discord/callback", config.ArenaHost, config.ArenaPort), discord.ScopeIdentify)
var redirections = make(map[string]string)

func startArenaService() {

	key := uuid.NewV5(uuid.NewV4(), "cookies").Bytes() // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                               // 30 days

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Domain = config.TradeHost
	store.Options.Path = "/"

	gothic.Store = store

	mux := pat.New()
	port := config.TradePort
	if config.TestMode {
		port = config.TradeTestPort
	}
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}

	mux.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(provider)
		gothic.CompleteUserAuth(res, req)
		state := gothic.SetState(req)
		redir, found := redirections[state]
		if !found {
			http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		} else {
			http.Redirect(res, req, redir, http.StatusTemporaryRedirect)
		}
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
	mux.HandleFunc("/a/{id}", http.HandlerFunc(ArenaIndex))
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

func ArenaIndex(w http.ResponseWriter, r *http.Request) {
	OtherID := ""
	// try to get it from the url param "provider"
	if p := r.URL.Query().Get(":id"); p != "" {
		OtherID = p
	}

	u, err := UserLogged(w, r)
	if err != nil {
		goth.UseProviders(provider)
		state := gothic.SetState(r)
		redirections[state] = fmt.Sprintf("/a/%s", OtherID)
		r.URL.RawQuery += fmt.Sprintf("&state=%s", state)
		state = gothic.SetState(r)
		r = r.WithContext(context.WithValue(r.Context(), "provider", "discord"))
		gothic.BeginAuthHandler(w, r)
		return
	}
	target := models.GetPlayer(OtherID)
	if target.IsNew {
		templates.ExecuteTemplate(w, "arenaMessage.html", struct {
			Message string
		}{
			Message: "You cannot trade with that person.",
		})
		return
	}

	token := uuid.NewV5(uuid.NewV4(), "trade").String()
	data := &websocket.TradeTemplateData{
		User: websocket.WebUser{
			Name: u.Name,
			ID:   u.UserID,
		},
		Token:   token,
		OtherID: OtherID,
		Host:    config.TradeHost,
	}
	websocket.Tokens.Add(token, data, 0)
	templates.ExecuteTemplate(w, "trade.html", data)
}
