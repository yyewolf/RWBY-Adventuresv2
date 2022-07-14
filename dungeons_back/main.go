package main

import (
	"os"
	"os/signal"
	dgrpc "rwby-adventures/dungeons_back/grpc"
	"rwby-adventures/dungeons_back/web"
	"rwby-adventures/dungeons_back/websocket"
	"syscall"
)

func main() {
	go dgrpc.Serve()
	web.StartWeb()
	websocket.StartWebsocket()

	// websocket.CreateArena(&arenapc.CreateArenaReq{
	// 	Id: "main",
	// })

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
