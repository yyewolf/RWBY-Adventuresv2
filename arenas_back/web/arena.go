package web

import (
	"context"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/arenas_back/cache"
	"rwby-adventures/arenas_back/static"
	"rwby-adventures/arenas_back/websocket"
	"rwby-adventures/config"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
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
	port := config.ArenaPort
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
	ArenaID := ""
	// try to get it from the url param "provider"
	if p := r.URL.Query().Get(":id"); p != "" {
		ArenaID = p
	}

	d, found := cache.Arenas.Get(ArenaID)
	if !found {
		fmt.Fprint(w, "kestufai ?")
		return
	}
	arena := d.(*cache.Arena)

	u, err := UserLogged(w, r)
	if err != nil {
		goth.UseProviders(provider)
		state := gothic.SetState(r)
		redirections[state] = fmt.Sprintf("/a/%s", ArenaID)
		r.URL.RawQuery += fmt.Sprintf("&state=%s", state)
		state = gothic.SetState(r)
		r = r.WithContext(context.WithValue(r.Context(), "provider", "discord"))
		gothic.BeginAuthHandler(w, r)
		return
	}

	h := sha256.Sum256([]byte(arena.ID + u.UserID))
	token := fmt.Sprintf("%x", h)
	data := &cache.User{
		Arena: arena,
		User: &cache.WebUser{
			Name: u.Name,
			ID:   u.UserID,
		},
		Token: token,
		Host:  config.ArenaHost,
		Port:  config.ArenaWebsocket,
	}
	websocket.Tokens.Add(token, data, 0)
	templates.ExecuteTemplate(w, "arena.html", data)
}
