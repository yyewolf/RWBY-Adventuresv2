package web

import (
	"html/template"
	"net/http"
	"rwby-adventures/config"
	"rwby-adventures/main/static"
	"strings"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	uuid "github.com/satori/go.uuid"
)

func startTradeService() {

	key := uuid.NewV5(uuid.NewV4(), "cookies").Bytes() // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                               // 30 days

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Domain = "localhost"
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

	tradeProvider := discord.New("375700234120200194", "P6KOz6Uvl8PWhY-hfx5IXo_posPDBu7D", "http://localhost:50/auth/discord/callback", discord.ScopeIdentify)

	mux.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(tradeProvider)
		gothic.CompleteUserAuth(res, req)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		res.Write([]byte(""))
	})

	mux.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(tradeProvider)
		gothic.Logout(res, req)
		http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
		res.Write([]byte(""))
	})

	mux.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		goth.UseProviders(tradeProvider)
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(`<p>UserID: {{.UserID}}</p>`)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	mux.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", DirectoryListing(http.FileServer(http.FS(static.Assets)))))
	mux.HandleFunc("/", http.HandlerFunc(TradeIndex))
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

type WebUser struct {
	Name string
}

type TemplateData struct {
	User  WebUser
	Token string
}

func TradeIndex(w http.ResponseWriter, r *http.Request) {
	u, err := UserLogged(w, r)
	if err != nil {
		http.Redirect(w, r, "/auth/discord", http.StatusTemporaryRedirect)
		return
	}
	token := uuid.NewV5(uuid.NewV4(), "trade").String()
	data := TemplateData{
		User: WebUser{
			Name: u.Name,
		},
		Token: token,
	}
	templates.ExecuteTemplate(w, "trade.html", data)
}
