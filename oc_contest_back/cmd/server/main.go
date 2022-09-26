package main

import (
	"os"
	"os/signal"
	"rwby-adventures/main/websocket"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/market_back/microservice"
	"rwby-adventures/market_back/web"
	"syscall"
)

func main() {
	cache.Init()
	websocket.StartWebsocket()
	microservice.CreateMicroservice()
	web.StartMarket()

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
