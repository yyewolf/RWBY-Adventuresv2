package web

import (
	"context"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/dungeons_back/static"
	"rwby-adventures/dungeons_back/websocket"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/yyewolf/goth"
	"github.com/yyewolf/goth/gothic"
	"github.com/yyewolf/goth/providers/discord"
)

var provider = discord.New(config.AppID, config.DiscordSecret, fmt.Sprintf("%sauth/discord/callback", config.DungeonHost), discord.ScopeIdentify)
var redirections = make(map[string]string)

func startDungeonService() {

	maxAge := 86400 * 30 // 30 days

	store := sessions.NewCookieStore(config.CookieKey)
	store.MaxAge(maxAge)
	store.Options.Domain = config.DungeonDomain
	store.Options.Secure = true
	store.Options.SameSite = http.SameSiteNoneMode
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
		UserLogged(res, req)
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
			providerName := "discord"

			provider, err := goth.GetProvider(providerName)
			if err != nil {
				fmt.Println(err)
				return
			}
			sess, err := provider.BeginAuth(gothic.SetState(req))
			if err != nil {
				fmt.Println(err)
				return
			}

			url, err := sess.GetAuthURL()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(sess)
			err = gothic.StoreInSession(providerName, sess.Marshal(), req, res)

			if err != nil {
				fmt.Println(err)
				return
			}

			http.Redirect(res, req, url, http.StatusTemporaryRedirect)
		}
	})

	mux.HandleFunc("/d/{id}", http.HandlerFunc(DungeonIndex))
	mux.PathPrefix("/").Handler(http.StripPrefix("/", DirectoryListing(http.FileServer(http.FS(static.Assets)))))
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
	dungeon := d.(*websocket.DungeonStruct)

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
	if u.UserID != dungeon.UserID {
		return
	}

	h := sha256.Sum256([]byte(dungeon.ID + u.UserID))
	token := fmt.Sprintf("%x", h)
	data := &websocket.DungeonUserData{
		Dungeon: dungeon,
		User: &websocket.WebUser{
			Name: u.Name,
			ID:   u.UserID,
		},
		Token: token,
	}
	websocket.Tokens.Add(token, data, 0)
	templates.ExecuteTemplate(w, "index.html", data)
}
