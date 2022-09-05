package main

import (
	"fmt"
	"net/http"
	"rwby-adventures/config"
	"strings"

	"github.com/gorilla/pat"
)

var evts = []string{"Halloween", "Xmas", "Summer", "Chinese_New_Year", "Valentines_Day"}

func hostImageService() {
	mux := pat.New()
	srv := http.Server{
		Addr:    config.IMGPort,
		Handler: mux,
	}
	mux.Get("/", func(res http.ResponseWriter, req *http.Request) {
		imageHandler(res, req)
	})

	go srv.ListenAndServe()
}

func contain(str string, substr []string) bool {
	for i := range substr {
		if strings.Contains(str, substr[i]) {
			return true
		}
	}
	return false
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	Path := r.URL.Path[1:]
	Path = "data/" + Path
	dat, err := webbox.ReadFile(Path)
	if err != nil && contain(Path, evts) {
		dat, err = webbox.ReadFile(Path)
		if err != nil {
			for i := range evts {
				Path = strings.ReplaceAll(Path, evts[i], "Default")
			}
			dat, _ = webbox.ReadFile(Path)
		}
	}
	fmt.Fprint(w, string(dat))
}
