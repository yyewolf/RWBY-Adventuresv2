package web

import (
	"context"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/dungeons/static"
	"rwby-adventures/dungeons/websocket"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

var provider = discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("http://%s%s/auth/discord/callback", config.DungeonHost, config.DungeonPort), discord.ScopeIdentify)
var redirections = make(map[string]string)

func startDungeonService() {

	key := uuid.NewV5(uuid.NewV4(), "cookies").Bytes() // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                               // 30 days

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Domain = config.DungeonHost
	store.Options.Path = "/"

	gothic.Store = store

	mux := pat.New()
	port := config.DungeonPort
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
	mux.HandleFunc("/d/{id}", http.HandlerFunc(DungeonIndex))
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

func DungeonIndex(w http.ResponseWriter, r *http.Request) {
	DungeonID := ""
	// try to get it from the url param "provider"
	if p := r.URL.Query().Get(":id"); p != "" {
		DungeonID = p
	}

	d, found := websocket.DungeonCache.Get(DungeonID)
	if !found {
		fmt.Fprint(w, "kestufai ?")
		return
	}
	arena := d.(*websocket.DungeonStruct)

	u, err := UserLogged(w, r)
	if err != nil {
		goth.UseProviders(provider)
		state := gothic.SetState(r)
		redirections[state] = fmt.Sprintf("/d/%s", DungeonID)
		r.URL.RawQuery += fmt.Sprintf("&state=%s", state)
		state = gothic.SetState(r)
		r = r.WithContext(context.WithValue(r.Context(), "provider", "discord"))
		gothic.BeginAuthHandler(w, r)
		return
	}

	h := sha256.Sum256([]byte(arena.ID + u.UserID))
	token := fmt.Sprintf("%x", h)
	data := &websocket.DungeonUserData{
		Dungeon: arena,
		User: &websocket.WebUser{
			Name: u.Name,
			ID:   u.UserID,
		},
		Token: token,
		Host:  config.DungeonHost,
		Port:  config.DungeonWebsocket,
	}
	websocket.Tokens.Add(token, data, 0)
	templates.ExecuteTemplate(w, "dungeon.html", data)
}
