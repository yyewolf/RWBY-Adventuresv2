package main

import (
	"os"
	"os/signal"
	"rwby-adventures/market_back/cache"
	"rwby-adventures/market_back/websocket"
	"syscall"
)

func main() {
	cache.Init()
	websocket.StartWebsocket()

	// websocket.CreateArena(&arenapc.CreateArenaReq{
	// 	Id: "main",
	// })

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
