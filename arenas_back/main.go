package main

import (
	"os"
	"os/signal"
	"rwby-adventures/arenas_back/microservice"
	"rwby-adventures/arenas_back/web"
	"rwby-adventures/arenas_back/websocket"
	"syscall"
)

func main() {
	web.StartWeb()
	websocket.StartWebsocket()
	microservice.CreateMicroservice()

	// websocket.CreateArena(&arenapc.CreateArenaReq{
	// 	Id: "main",
	// })

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
