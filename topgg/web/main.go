package web

import (
	"net/http"
	"rwby-adventures/config"
)

func StartWeb() {
	mux := http.NewServeMux()
	srv := http.Server{
		Addr:    config.TopGGPort,
		Handler: mux,
	}
	mux.HandleFunc("/", index)
	go srv.ListenAndServe()
}
