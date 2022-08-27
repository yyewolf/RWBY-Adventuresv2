package web

import (
	"fmt"
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
	fmt.Println("[TOPGG] Starting web server.")
	go srv.ListenAndServe()
}
