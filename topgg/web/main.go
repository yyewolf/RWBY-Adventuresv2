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
	fmt.Printf("[TOPGG] Starting web server with bearer '%s'.\n", config.TopGG)
	go srv.ListenAndServe()
}
